// @Author fcihpy
// @Date 2024/7/25 15:32:00
// @Desc
//

package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fcihpy001/chain-service/logger"
	"github.com/fcihpy001/chain-service/service"
	"github.com/fcihpy001/chain-service/util"
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

func TestBatchQueryGasBalance(t *testing.T) {
	service.ParallelQueryAsset()
}

func TestTransferAsset(t *testing.T) {
	err := service.TransferAsset(28, 10, 0.00003, util.TokenETH)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
