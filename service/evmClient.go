// @Author fcihpy
// @Date 2024/7/25 14:53:00
// @Desc
//

package service

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fcihpy001/chain-service/logger"
	"math/big"
	"sync"
	"time"
)

type EvmClient struct {
	sync.Mutex
	RawClients []*ethclient.Client
	Urls       []string
	ChainId    *big.Int
}

func GetEvmClient(rpc string, chainId int) (*ethclient.Client, error) {
	clients, err := Dial(context.Background(), chainId, []string{rpc})
	if err != nil {
		return nil, err
	}
	return clients.RawClients[0], nil
}

func Dial(ctx context.Context, chainId int, rpcs []string) (*EvmClient, error) {
	var wg sync.WaitGroup
	var client EvmClient
	logger.Debug("starting get evm client...")

	c, cancel := context.WithTimeout(ctx, time.Duration(time.Second)*3)
	defer cancel()

	start := time.Now()
	for _, rpc := range rpcs {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			_client, err := ethclient.DialContext(c, url)
			if err != nil {
				logger.ErrorF("Failed to dial %s-%s", url, err.Error())
				return
			}

			_chainId, err := _client.ChainID(ctx)
			if err != nil {
				logger.ErrorF("Failed to get chain id: %s", err.Error())
				return
			}

			if _chainId.Int64() == int64(chainId) {
				client.Lock()
				client.Urls = append(client.Urls, url)
				client.RawClients = append(client.RawClients, _client)
				client.Unlock()
			} else {
				logger.WarnF("Invalid RPC url %s", rpc)
			}
			logger.DebugF("dial url %s took %s", url, time.Since(start))

		}(rpc)
	}
	wg.Wait()

	if len(client.RawClients) == 0 {
		logger.Error("Failed to dial,no rpc available", "urls", client.Urls)
		return nil, errors.New("no rpc available")
	}
	client.ChainId = big.NewInt(int64(chainId))
	logger.DebugF("get evm client took (%.3f)s", time.Since(start).Seconds())

	return &client, nil
}
