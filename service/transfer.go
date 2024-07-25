// @Author fcihpy
// @Date 2024/7/25 14:55:00
// @Desc
//

package service

import (
	"context"
	"github.com/DappOSDao/NodeSwapX/service"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fcihpy001/chain-service/contract"
	"github.com/fcihpy001/chain-service/logger"
	"github.com/fcihpy001/chain-service/util"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
)

func TransferAsset(nodeId int, chainId int, amount float64, tokenName string) error {

	client, err := GetEvmClient(getRPC(chainId), chainId)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	fromAddr := common.HexToAddress(service.GetAccount().Address)
	receiverAddr := common.HexToAddress(GetNodeAddr(nodeId))
	tokenInfo := service.GetTokenInfo(nodeId, chainId, tokenName)

	callData, err := contract.GetTransferCallData(receiverAddr, amount, tokenInfo.TokenDecimal)
	if err != nil {
		return err
	}

	amountBIg := big.NewInt(0)
	if tokenName != util.TokenETH {
		receiverAddr = common.HexToAddress(service.GetTokenInfo(nodeId, chainId, tokenName).TokenAddress)
	} else {
		callData = nil
		tokenDecimal := math.Pow10(tokenInfo.TokenDecimal)
		amountDecimal := decimal.NewFromFloat(amount).Mul(decimal.NewFromFloat(tokenDecimal))
		amountBIg = amountDecimal.BigInt()
	}

	callMsg := ethereum.CallMsg{
		From:     fromAddr,
		To:       &receiverAddr,
		Gas:      0,
		GasPrice: nil,
		Value:    amountBIg,
		Data:     callData,
	}
	estimateGas, err := client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		return err
	}

	nonce, err := client.NonceAt(context.Background(), fromAddr, nil)
	if err != nil {
		return err
	}

	chainIdBig := big.NewInt(0).SetInt64(int64(chainId))
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      estimateGas,
		To:       &receiverAddr,
		Value:    amountBIg,
		Data:     callData,
	})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainIdBig), service.GetAccount().PrivateKey)
	if err != nil {
		return err
	}

	if err = client.SendTransaction(context.Background(), signedTx); err != nil {
		return err
	}

	logger.InfoF("%s/tx/%s", service.GetChainInfo(chainId).BlockExplorerUrl, signedTx.Hash().String())

	return nil
}
