// @Author fcihpy
// @Date 2024/7/25 16:14:00
// @Desc
//

package service

import (
	"bytes"
	"context"
	"fmt"
	swapservice "github.com/DappOSDao/NodeSwapX/service"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fcihpy001/chain-service/contract"
	"github.com/fcihpy001/chain-service/logger"
	"github.com/fcihpy001/chain-service/model"
	"github.com/fcihpy001/chain-service/util"
	"github.com/jedib0t/go-pretty/v6/table"
	"math"
	"math/big"
	"sync"
)

var (
	mutex sync.Mutex
)

func ParallelQueryAsset() {
	tokens := []string{util.TokenUSDT, util.TokenUSDC, util.TokenETH}
	chainIds := []int{1, 10, 56, 137, 169, 42161, 43114}
	rowHeaders := table.Row{"wallet", "chainId", "token", "balance", "tokenLevel", "gasName", "gasLevel"}

	wg := sync.WaitGroup{}
	for _, tokenName := range tokens {
		for _, chainId := range chainIds {
			wg.Add(1)

			go func(tokenName string, chainId int) {
				defer wg.Done()

				infos, err := BatchQueryAssetBalance(GetAccounts(), chainId, tokenName)
				if err != nil {
					return
				}

				mutex.Lock()
				defer mutex.Unlock()
				t := table.NewWriter()
				t.SetStyle(table.StyleLight)
				tTemp := table.Table{}
				tTemp.Render()
				t.AppendHeader(rowHeaders)

				fmt.Printf("*************%d-%s-**********\n", chainId, tokenName)
				for _, info := range infos {
					t.AppendRow([]interface{}{info.Wallet, info.ChainId, info.TokenName, info.Balance, "20u", info.GasName, info.Min})
				}
				fmt.Println(t.Render())
			}(tokenName, chainId)
		}
	}
	wg.Wait()
}

func BatchQueryAssetBalance(wallets []string, chainId int, tokenName string) ([]model.AssetInfo, error) {
	tokenInfo := swapservice.GetTokenInfo(2, chainId, tokenName)
	tokenAddr := common.HexToAddress(tokenInfo.TokenAddress)
	tokenDecimal := tokenInfo.TokenDecimal

	rpc := getRPC(chainId)
	client, err := GetEvmClient(rpc, chainId)
	if err != nil {
		return nil, err
	}

	multiCallAddr := common.HexToAddress(util.MultiCall)
	multiCallABI, err := abi.JSON(bytes.NewReader([]byte(contract.GetMultiCallABI())))
	if err != nil {
		return nil, err
	}

	tokenABI, err := abi.JSON(bytes.NewReader([]byte(contract.GetErc20ABI())))
	if err != nil {
		return nil, err
	}

	methodName := util.MethodBalanceOf
	if tokenName == util.TokenETH {
		tokenAddr = multiCallAddr
		tokenABI = multiCallABI
		methodName = util.MethodGetEthBalance
	}

	var calls []model.CalInfo
	for _, wallet := range wallets {
		callData, err := tokenABI.Pack(methodName, common.HexToAddress(wallet))
		if err != nil {
			logger.Panic(err.Error())
		}
		calls = append(calls, model.CalInfo{
			Target:   tokenAddr,
			CallData: callData,
		})
	}

	packedCalls, err := multiCallABI.Pack(util.MethodAggregate, calls)
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{
		To:   &multiCallAddr,
		Data: packedCalls,
	}
	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, err
	}

	var contractData struct {
		BlockNumber *big.Int
		ReturnData  [][]byte
	}
	err = multiCallABI.UnpackIntoInterface(&contractData, util.MethodAggregate, result)
	if err != nil {
		return nil, err
	}

	balances := make([]model.AssetInfo, 0)
	for i, data := range contractData.ReturnData {
		num := big.NewInt(0).SetBytes(data)
		divisor := big.NewFloat(0).SetFloat64(math.Pow10(tokenDecimal))
		balance := big.NewFloat(0).Quo(big.NewFloat(0).SetInt(num), divisor)

		if chainId == 56 && tokenName == util.TokenETH {
			tokenName = util.TokenBNB
		} else if chainId == 43114 && tokenName == util.TokenETH {
			tokenName = util.TokenAvax
		}

		balances = append(balances, model.AssetInfo{
			ChainId:   chainId,
			Wallet:    wallets[i],
			Balance:   balance,
			Decimal:   tokenDecimal,
			GasName:   GetGasWater(chainId).GasName,
			TokenName: tokenName,
			Min:       GetGasWater(chainId).Min,
		})
	}
	return balances, err
}
