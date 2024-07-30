// @Author fcihpy
// @Date 2024/7/25 14:54:00
// @Desc
//

package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fcihpy001/chain-service/contract"
	"github.com/fcihpy001/chain-service/logger"
	"go.uber.org/zap"
	"log"
	"sync"
	"time"
)

func TryGetVlogs(chainId int, txHash string, maxQueryTimes int) (*types.Receipt, error) {

	start := time.Now()
	var (
		err     error
		receipt *types.Receipt
	)

	if maxQueryTimes < 1 {
		maxQueryTimes = 10
	}

	defer func() {
		logger.Info("handleReceipt elapsed seconds", zap.Int("chainId", chainId), zap.Any("timeElapsed", time.Since(start)))
	}()

	for i := 0; i < maxQueryTimes; i++ {

		receipt, err = getVlogs(chainId, txHash)
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}
		return receipt, nil
	}
	return receipt, nil
}

func getVlogs(chainId int, txHash string) (*types.Receipt, error) {
	start := time.Now()
	evmClient, err := Dial(context.Background(), chainId, []string{GetRPC(chainId)})
	if err != nil {
		log.Fatalf("GetEthClientByChainId err: %v", err)
	}

	logsChannel := make(chan *types.Receipt)
	quitChannel := make(chan struct{})
	var wg sync.WaitGroup
	var once sync.Once

	for _, client := range evmClient.RawClients {
		wg.Add(1)
		go func(hash string, client *ethclient.Client) {
			defer wg.Done()

			receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(hash))
			if err != nil {
				return
			}
			if len(receipt.Logs) > 0 || receipt.Status == 0 {
				once.Do(func() {
					logsChannel <- receipt
					close(quitChannel)
				})
			}
		}(txHash, client)
	}

	go func() {
		wg.Wait()
		close(logsChannel)
	}()

	var receipt *types.Receipt
	select {
	case receipt = <-logsChannel:
		logger.Info("get receipt transaction finished, elapsed time:", zap.Any("time", time.Since(start)))
		return receipt, nil

	case <-time.After(3 * time.Second):
		return nil, errors.New("get receipt timeout")

	case <-quitChannel:
		logger.Info("get receipt channel quit")
	}
	return receipt, nil
}

func MonitorPayDBEvent() {
	client, err := ethclient.Dial("wss://open-platform.nodereal.io/ws/44acc0b89a0c4643be14ed0b58307a4c/arbitrum-nitro/")
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress("0xa13db51beb831545dcde204ec0dac0d867b3b467")},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println("--------New PayDB Event", time.Now())

			fmt.Println("address", vLog.Address.Hex())
			fmt.Println("event: ", vLog.Topics[0].Hex())
			fmt.Println("from: ", vLog.Topics[1].Hex())
			fmt.Println("Value: ", vLog.Data)
			fmt.Println("TxHash: ", vLog.TxHash.Hex())
			fmt.Println()

		}
	}
}

func MonitorVWEvent() {
	client, err := ethclient.Dial("wss://open-platform.nodereal.io/ws/44acc0b89a0c4643be14ed0b58307a4c/arbitrum-nitro/")
	if err != nil {
		log.Fatal(err)
	}

	vwTxExecuteEventId, err := GetVwManagerEventId()
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress("0x4b88f06b6b1b58c675243932d35f674972194fd7")},
		Topics:    [][]common.Hash{{vwTxExecuteEventId}},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println("--------New Transfer Event", time.Now())

			fmt.Println("address", vLog.Address.Hex())
			fmt.Println("event: ", vLog.Topics[0].Hex())
			fmt.Println("from: ", vLog.Topics[1].Hex())
			fmt.Println("Value: ", vLog.Data)
			fmt.Println("TxHash: ", vLog.TxHash.Hex())
			fmt.Println()

		}
	}
}

func GetPayDBExecutedEventId() (common.Hash, error) {

	contractABI, err := abi.JSON(bytes.NewReader([]byte(contract.GetPayDbABI())))
	if err != nil {
		logger.ErrorF("abi json error: %s", err.Error())
		return common.Hash{}, errors.New("abi error")
	}
	return contractABI.Events["OrderExecuted"].ID, nil
}

func GetPayDBCreatedEventId() (common.Hash, error) {

	contractABI, err := abi.JSON(bytes.NewReader([]byte(contract.GetPayDbABI())))
	if err != nil {
		logger.ErrorF("abi json error: %s", err.Error())
		return common.Hash{}, errors.New("abi error")
	}
	return contractABI.Events["OrderCreated"].ID, nil
}

func GetVwManagerEventId() (common.Hash, error) {

	contractABI, err := abi.JSON(bytes.NewReader([]byte(contract.GetVwABI())))
	if err != nil {
		logger.ErrorF("abi json error: %s", err.Error())
		return common.Hash{}, errors.New("abi error")
	}
	return contractABI.Events["TxExecuted"].ID, nil
}
