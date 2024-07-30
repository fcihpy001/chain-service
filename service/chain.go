// @Author fcihpy
// @Date 2024/7/30 11:43:00
// @Desc
//

package service

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fcihpy001/chain-service/contract"
	"github.com/fcihpy001/chain-service/logger"
	"github.com/fcihpy001/chain-service/model"
	"go.uber.org/zap"
	"math/big"
	"strings"
)

func SendTransaction(contractParam model.ContractParam) (*types.Transaction, error) {
	client, err := GetEvmClient(GetRPC(contractParam.ChainId), contractParam.ChainId)
	if err != nil {
		logger.Error("GetMultiEthClients err:", zap.Error(err))
		return nil, err
	}
	// gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	// todo increase gas
	if err != nil {
		logger.Error("client.SuggestGasPrice err:", zap.Error(err))
		return nil, err
	}
	gasPrice.Mul(gasPrice, big.NewInt(110))
	gasPrice.Div(gasPrice, big.NewInt(100))

	sender := GetAccount().Address

	callData, err := GetCreateSrcOrderCallData(contractParam)
	if err != nil {
		logger.Error("getTransferCallData err:", zap.Error(err))
	}
	callMsg := ethereum.CallMsg{
		From:     common.HexToAddress(sender),
		To:       &contractParam.PayDB,
		Gas:      0,
		GasPrice: nil,
		Value:    contractParam.Amount,
		Data:     callData,
	}

	estimateGas, err := client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		logger.Error("client.EstimateGas err:", zap.Error(err))
		return nil, err
	}
	estimateGas += 1000000

	nonce, err := client.NonceAt(context.Background(), common.HexToAddress(sender), nil)
	if err != nil {
		logger.ErrorF("client.PendingNonceAt err:%v", zap.Error(err))
	}

	chainIdBig := big.NewInt(int64(contractParam.ChainId))
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      estimateGas,
		To:       &contractParam.PayDB,
		Value:    contractParam.Amount,
		Data:     callData,
	})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainIdBig), GetAccount().PrivateKey)
	if err != nil {
		logger.ErrorF("SignTx err:%v", zap.Error(err))
		return nil, err
	}

	if err = client.SendTransaction(context.Background(), signedTx); err != nil {
		logger.ErrorF("client.SendTransaction err:", zap.Error(err))
		return nil, err
	}
	logger.InfoF("Successfully sent transaction! %v", signedTx.Hash().String())
	return signedTx, nil

}

func SendIntentTransaction(contractParam model.ContractParam) (*types.Transaction, error) {
	client, err := GetEvmClient(GetRPC(contractParam.ChainId), contractParam.ChainId)
	if err != nil {
		logger.Error("GetMultiEthClients err:", zap.Error(err))
		return nil, err
	}
	// gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	// todo increase gas
	if err != nil {
		logger.Error("client.SuggestGasPrice err:", zap.Error(err))
		return nil, err
	}
	gasPrice.Mul(gasPrice, big.NewInt(110))
	gasPrice.Div(gasPrice, big.NewInt(100))

	sender := GetAccount().Address

	callData, err := GetIntentDBCreateSrcOrderCallData(contractParam)
	if err != nil {
		logger.Error("getTransferCallData err:", zap.Error(err))
	}
	callMsg := ethereum.CallMsg{
		From:     common.HexToAddress(sender),
		To:       &contractParam.PayDB,
		Gas:      0,
		GasPrice: nil,
		Value:    contractParam.Amount,
		Data:     callData,
	}

	estimateGas, err := client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		logger.Error("client.EstimateGas err:", zap.Error(err))
		return nil, err
	}
	estimateGas += 1000000

	nonce, err := client.NonceAt(context.Background(), common.HexToAddress(sender), nil)
	if err != nil {
		logger.ErrorF("client.PendingNonceAt err:%v", zap.Error(err))
	}

	chainIdBig := big.NewInt(int64(contractParam.ChainId))
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      estimateGas,
		To:       &contractParam.PayDB,
		Value:    contractParam.Amount,
		Data:     callData,
	})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainIdBig), GetAccount().PrivateKey)
	if err != nil {
		logger.ErrorF("SignTx err:%v", zap.Error(err))
		return nil, err
	}

	if err = client.SendTransaction(context.Background(), signedTx); err != nil {
		logger.ErrorF("client.SendTransaction err:", zap.Error(err))
		return nil, err
	}
	logger.InfoF("Successfully sent transaction! %v", signedTx.Hash().String())
	return signedTx, nil

}

func GetCreateSrcOrderCallData(contractParam model.ContractParam) ([]byte, error) {

	contractAbi, err := abi.JSON(strings.NewReader(contract.GetPayDbABI()))
	if err != nil {
		logger.ErrorF("AbiJson err:%v", err)
		return nil, err
	}

	callData, err := contractAbi.Pack(
		"createSrcOrderETH",
		contractParam.Owner,
		contractParam.Wallet,
		contractParam.Receiver,
		contractParam.CreatePays,
		contractParam.Vw,
		contractParam.CallData,
	)
	if err != nil {
		logger.ErrorF("Pack err:%v", err)
		return nil, err
	}
	return callData, nil
}

func GetIntentDBCreateSrcOrderCallData(contractParam model.ContractParam) ([]byte, error) {

	contractAbi, err := abi.JSON(strings.NewReader(contract.GetIntentDBABI()))
	if err != nil {
		logger.ErrorF("ABIJson err:%v", err)
		return nil, err
	}

	callData, err := contractAbi.Pack(
		"createSrcIntentOrderETH",
		contractParam.Owner,
		contractParam.Receiver,
		contractParam.IntentCreatePays,
		contractParam.IntentCallParam,
	)
	if err != nil {
		logger.ErrorF("Pack err:%v", err)
		return nil, err
	}
	return callData, nil
}
