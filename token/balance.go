// @Author fcihpy
// @Date 2024/7/25 16:14:00
// @Desc
//

package token

import (
	"bytes"
	"context"
	"fmt"
	"github.com/DappOSDao/node-service/network"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fcihpy001/chain-service/constant"
	"github.com/fcihpy001/chain-service/contract"
	"github.com/fcihpy001/chain-service/logger"
	"github.com/fcihpy001/chain-service/model"
	"github.com/fcihpy001/chain-service/service"
	"github.com/jedib0t/go-pretty/v6/table"
	"math"
	"math/big"
	"sync"
)

var (
	mutex sync.Mutex
)

func ParallelQueryAsset(chainIds []int, tokens []string) {
	if len(tokens) == 0 {
		tokens = []string{constant.TokenUSDT, constant.TokenUSDC, constant.TokenETH}
	}
	if len(chainIds) == 0 {
		chainIds = []int{1, 10, 56, 137, 169, 42161, 43114}
	}
	rowHeaders := table.Row{"wallet", "chainId", "token", "balance", "tokenLevel", "gasName", "gasLevel"}

	wg := sync.WaitGroup{}
	for _, tokenName := range tokens {
		for _, chainId := range chainIds {
			wg.Add(1)

			go func(tokenName string, chainId int) {
				defer wg.Done()

				infos, err := BatchQueryAssetBalance(service.GetAccounts(), chainId, tokenName)
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
	tokenInfo := network.GetTokenInfo(2, chainId, tokenName)
	tokenAddr := common.HexToAddress(tokenInfo.TokenAddress)
	tokenDecimal := tokenInfo.TokenDecimal

	rpc := service.GetRPC(chainId)
	client, err := service.GetEvmClient(rpc, chainId)
	if err != nil {
		return nil, err
	}

	multiCallAddr := common.HexToAddress(constant.MultiCall)
	multiCallABI, err := abi.JSON(bytes.NewReader([]byte(contract.GetMultiCallABI())))
	if err != nil {
		return nil, err
	}

	tokenABI, err := abi.JSON(bytes.NewReader([]byte(contract.GetERC20ABI())))
	if err != nil {
		return nil, err
	}

	methodName := constant.MethodBalanceOf
	if tokenName == constant.TokenETH {
		tokenAddr = multiCallAddr
		tokenABI = multiCallABI
		methodName = constant.MethodGetEthBalance
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

	packedCalls, err := multiCallABI.Pack(constant.MethodAggregate, calls)
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
	err = multiCallABI.UnpackIntoInterface(&contractData, constant.MethodAggregate, result)
	if err != nil {
		return nil, err
	}

	balances := make([]model.AssetInfo, 0)
	for i, data := range contractData.ReturnData {
		num := big.NewInt(0).SetBytes(data)
		divisor := big.NewFloat(0).SetFloat64(math.Pow10(tokenDecimal))
		balance := big.NewFloat(0).Quo(big.NewFloat(0).SetInt(num), divisor)

		if chainId == 56 && tokenName == constant.TokenETH {
			tokenName = constant.TokenBNB
		} else if chainId == 43114 && tokenName == constant.TokenETH {
			tokenName = constant.TokenAvax
		}

		balances = append(balances, model.AssetInfo{
			ChainId:   chainId,
			Wallet:    wallets[i],
			Balance:   balance,
			Decimal:   tokenDecimal,
			GasName:   service.GetGasWater(chainId).GasName,
			TokenName: tokenName,
			Min:       service.GetGasWater(chainId).Min,
		})
	}
	return balances, err
}

func QueryAssetOnSameChain(wallet string, chainId int, tokens []string) ([]model.AssetInfo, error) {
	wg := sync.WaitGroup{}
	var infos []model.AssetInfo

	client, err := service.GetEvmClient(service.GetRPC(chainId), chainId)
	if err != nil {
		return nil, err
	}

	tokenABI, err := abi.JSON(bytes.NewReader([]byte(contract.GetERC20ABI())))
	if err != nil {
		return nil, err
	}

	callData, err := tokenABI.Pack("balanceOf", common.HexToAddress(wallet))
	if err != nil {
		logger.Panic(err.Error())
	}

	for _, token := range tokens {
		wg.Add(1)
		go func(tokenName string, chainId int) {
			defer wg.Done()
			tokenInfo := network.GetTokenInfo(2, chainId, tokenName)

			to := common.HexToAddress(tokenInfo.TokenAddress)
			msg := ethereum.CallMsg{
				To:   &to,
				Data: callData,
			}
			balanceBig := big.NewInt(0)

			if tokenName == constant.TokenETH {
				balanceBig, err = client.BalanceAt(context.Background(), common.HexToAddress(wallet), nil)
			} else {
				balanceByte, err := client.CallContract(context.Background(), msg, nil)
				if err != nil {
					return
				}
				balanceBig = big.NewInt(0).SetBytes(balanceByte)
			}

			divisor := big.NewFloat(0).SetFloat64(math.Pow10(tokenInfo.TokenDecimal))
			balance := big.NewFloat(0).Quo(big.NewFloat(0).SetInt(balanceBig), divisor)

			if chainId == 56 && tokenName == constant.TokenETH {
				tokenName = constant.TokenBNB
			} else if chainId == 43114 && tokenName == constant.TokenETH {
				tokenName = constant.TokenAvax
			}

			infos = append(infos, model.AssetInfo{
				ChainId:   chainId,
				Wallet:    wallet,
				Balance:   balance,
				Decimal:   tokenInfo.TokenDecimal,
				TokenName: tokenName,
				GasName:   service.GetGasWater(chainId).GasName,
				Min:       service.GetGasWater(chainId).Min,
			})

		}(token, chainId)

	}
	wg.Wait()

	return infos, nil
}
