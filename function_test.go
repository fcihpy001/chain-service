// @Author fcihpy
// @Date 2024/7/25 15:32:00
// @Desc
//

package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fcihpy001/chain-service/constant"
	"github.com/fcihpy001/chain-service/contract"
	"github.com/fcihpy001/chain-service/logger"
	"github.com/fcihpy001/chain-service/model"
	"github.com/fcihpy001/chain-service/service"
	"github.com/fcihpy001/chain-service/swap"
	"github.com/fcihpy001/chain-service/token"
	"github.com/fcihpy001/chain-service/util"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"math"
	"math/big"
	"testing"
)

func TestClient(t *testing.T) {

	client, err := service.Dial(context.Background(), 10, []string{"https://opt-mainnet.nodereal.io/v1/44acc0b89a0c4643be14ed0b58307a4c"})
	if err != nil {
		panic(err)
	}
	balanceBig, err := client.RawClients[0].BalanceAt(context.Background(), common.HexToAddress("0x178ad9283bba524d2e10ad1102a079144bc569e0"), nil)
	if err != nil {
		return
	}
	decimal := big.NewFloat(0).SetFloat64(math.Pow10(18))
	balance := big.NewFloat(0).Quo(big.NewFloat(0).SetInt(balanceBig), decimal)
	logger.DebugF("balance: %v", balance)
}

func TestTransferAsset(t *testing.T) {
	err := token.TransferAsset(25, 42161, 10, constant.TokenUSDC)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func TestMonitor(t *testing.T) {
	service.MonitorVWEvent()
}

func TestGetPayDBCreateVlogs(t *testing.T) {
	hash := "0xc03d7709a52414124d534c340186647b36858ed36671410dee0a610522719206"
	receipt, err := service.TryGetVlogs(10, hash, 3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if receipt.Status == 0 {
		logger.Error("receipt status err", zap.Any("receipt", receipt))
	}
	eventId, err := service.GetPayDBCreatedEventId()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	payDBAddr := common.HexToAddress("0xc9316435aef3762f25d2d0a31ec95fe4e6798085")
	payDBContract, err := contract.NewPayDB(payDBAddr, nil)
	if err != nil {
		return
	}

	for _, vlog := range receipt.Logs {
		if vlog.Address == payDBAddr &&
			vlog.Topics[0] == eventId {

			orderCreatedEvent, err := payDBContract.ParseOrderCreated(*vlog)
			if err != nil {
				logger.Error(err.Error())
				return
			}
			logger.Info("-------------OrderCreatedEvent  start--------------")
			logger.InfoF("orderCreatedEvent paydb %s", receipt.Logs[0].Address)
			logger.InfoF("orderCreatedEvent Owner %s", orderCreatedEvent.Owner.String())
			logger.InfoF("orderCreatedEvent payOrderId %s", orderCreatedEvent.PayOrderId.String())
			logger.InfoF("orderCreatedEvent TokenIn %s", orderCreatedEvent.TokenIn.String())
			logger.InfoF("orderCreatedEvent AmountIn %s", orderCreatedEvent.AmountIn.String())
			logger.InfoF("orderCreatedEvent TokenOut %s", orderCreatedEvent.TokenOut.String())
			logger.InfoF("orderCreatedEvent AmountOut %s", orderCreatedEvent.AmountOut.String())
			logger.InfoF("orderCreatedEvent Receiver %s", orderCreatedEvent.Receiver.String())
			logger.InfoF("orderCreatedEvent Node %s", orderCreatedEvent.Node.String())
			logger.InfoF("orderCreatedEvent Code %s", orderCreatedEvent.Code.String())
			logger.InfoF("orderCreatedEvent Sender %s", orderCreatedEvent.Sender.String())
			logger.InfoF("orderCreatedEvent WfHash %s", common.BytesToHash(orderCreatedEvent.WfHash[:]).Hex())
			logger.InfoF("orderCreatedEvent %s", orderCreatedEvent)
			logger.Info("-------------OrderCreatedEvent  end--------------")
		}
	}
}

func TestGetPayDBExecuteVlogs(t *testing.T) {
	hash := "0x7608de18dcfa688dfa3f5e27dd0a455f17abee6748256f2d9becd21fbe2a152e"
	receipt, err := service.TryGetVlogs(43114, hash, 3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if receipt.Status != 1 {
		logger.Error("receipt status err", zap.Any("receipt", receipt))
	}
	eventId, err := service.GetPayDBExecutedEventId()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	vwEventId, err := service.GetVwManagerEventId()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	vwAddr := common.HexToAddress("0x0acaf075270510f4886416b5069f8c6be3364f7e")
	vwContract, err := contract.NewVWManager(vwAddr, nil)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	payDBAddr := common.HexToAddress("0xa13db51beb831545dcde204ec0dac0d867b3b467")
	payDBContract, err := contract.NewPayDB(payDBAddr, nil)
	if err != nil {
		return
	}

	for _, vlog := range receipt.Logs {
		if vlog.Address == payDBAddr &&
			vlog.Topics[0] == eventId {

			orderExecutedEvent, err := payDBContract.ParseOrderExecuted(*vlog)
			if err != nil {
				logger.Error(err.Error())
				return
			}
			logger.Info("-------------orderExecutedEvent start--------------")
			logger.InfoF("orderExecutedEvent paydb: %s", vlog.Address)
			logger.InfoF("orderExecutedEvent Owner: %s", orderExecutedEvent.Owner.String())
			logger.InfoF("orderExecutedEvent payOrderId: %s", orderExecutedEvent.PayOrderId.String())
			logger.InfoF("orderExecutedEvent TokenOut: %s", orderExecutedEvent.TokenOut.String())
			logger.InfoF("orderExecutedEvent AmountOut: %s", orderExecutedEvent.AmountOut.String())
			logger.InfoF("orderExecutedEvent Receiver: %s", orderExecutedEvent.Receiver.String())
			logger.InfoF("orderExecutedEvent Code: %s", orderExecutedEvent.Code.String())
			logger.InfoF("orderExecutedEvent WfHash: %s", common.BytesToHash(orderExecutedEvent.WfHash[:]).Hex())
			logger.Info("-------------orderExecutedEvent  end--------------")
			assert.EqualValues(t, orderExecutedEvent.PayOrderId.String(), "1252837268618252239918749603762610577376739917825")
		}

		if vlog.Address == vwAddr && vlog.Topics[0] == vwEventId {
			vwTxExecuted, err := vwContract.ParseTxExecuted(*vlog)
			if err != nil {
				logger.Error(err.Error())
				return
			}

			logger.Info("-------------vwTxExecuted  start--------------")
			logger.InfoF("txExecuted event Wallet %s", vwTxExecuted.Wallet.String())
			logger.InfoF("txExecuted event Owner %s", vwTxExecuted.Owner.String())
			logger.InfoF("txExecuted event Code %s", vwTxExecuted.Code.String())
			logger.InfoF("txExecuted event RootHash %s", common.BytesToHash(vwTxExecuted.RootHash[:]).Hex())
			logger.InfoF("txExecuted event Res %v", vwTxExecuted.ResCode)
			logger.InfoF("txExecuted:%s", vwTxExecuted)
			logger.Info("-------------vwTxExecuted  end--------------")
			assert.EqualValues(t, vwTxExecuted.ResCode, true)
		}
	}

}

func TestQueryAssetOnSameChain(t *testing.T) {
	infos, err := token.QueryAssetOnSameChain(service.GetAccount().Address, 56, []string{"USDT", "USDC", "ETH"})
	if err != nil {
		logger.Error(err.Error())
	}
	for _, info := range infos {
		fmt.Printf("chain:%d-wallet:%s-token:%s-balance:%#v\n", info.ChainId, info.Wallet, info.TokenName, info.Balance)
	}
}

func TestContainNodeId(t *testing.T) {

	r1 := util.ContainNodeId(constant.TaskTypeIntentDB, 21)
	r2 := util.ContainNodeId(constant.TaskTypePayDB, 28)
	t.Log(r1, r2)
}

func TestMismatchType(t *testing.T) {

	nodeTask := model.Task{
		Type:       constant.TaskTypeIntentDB,
		NodeId:     21,
		SrcChainId: 10,
		DstChainId: 42161,
		Amounts:    []float64{0.05},
		TokenIns:   []string{"USDT"},
		TokenOuts:  []string{"USDT"},
		HasSrcVw:   false,
		HasDstVw:   false,
	}
	logger.Info("task info", zap.Any("task", nodeTask))
	_, err := swap.PayDBSwap([]model.Task{nodeTask}, false)
	if err != nil {
		logger.Error(err.Error())
		return
	}
}

func TestBatchQueryAssetBalance(t *testing.T) {
	token.ParallelQueryAsset(nil, nil)
}
