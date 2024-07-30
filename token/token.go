// @Author fcihpy
// @Date 2024/7/29 18:57:00
// @Desc
//

package token

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/DappOSDao/node-service/network"
	"github.com/fcihpy001/chain-service/constant"
	"github.com/fcihpy001/chain-service/logger"
	"github.com/fcihpy001/chain-service/model"
	"github.com/fcihpy001/chain-service/service"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func isNativeToken(token string) bool {
	return token == constant.ZeroAddr || token == constant.Zero42Addr
}

func CalculateBridgeFee(nodeId, chainId int) decimal.Decimal {

	k, c := network.GetVaultKC(nodeId, chainId)
	if k == "" || c == "" {

		//logger.Error("cannot calculate vw bridge fee", zap.Any("chainId", chainId))
		return decimal.NewFromInt(0)
	}

	client, err := service.GetEvmClient(service.GetRPC(chainId), chainId)
	if err != nil {
		logger.Error("get client error", zap.Any("chainId", chainId))
		return decimal.NewFromInt(0)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Error("e2040: suggest gas price error", zap.Any("chainId", chainId))
		return decimal.NewFromInt(0)
	}

	gasPriceDecimal := decimal.NewFromBigInt(gasPrice, 0)
	kDecimal, err := decimal.NewFromString(k)
	cDecimal, err := decimal.NewFromString(c)
	if err != nil {
		logger.Error("e2041: get k,c decimal error", zap.Any("chainId", chainId))
		return decimal.NewFromInt(0)
	}
	chainInfo := network.GetChainInfo(chainId)
	price := GetTokenPrice(chainInfo.NativeCurrency.Name)

	tokenPriceDecimal := decimal.NewFromFloat(price)
	baseFee := gasPriceDecimal.
		Mul(kDecimal).
		Mul(tokenPriceDecimal).
		Mul(decimal.NewFromBigInt(big.NewInt(1), -int32(chainInfo.NativeCurrency.Decimals)))
	bridgeFee := baseFee.Add(cDecimal.Div(decimal.NewFromInt(100000000)))
	return bridgeFee
}

func GetTokenPrice(tokenName string) float64 {

	tokenName = strings.ToLower(tokenName)
	if tokenName == "weth.e" {
		tokenName = "eth"
	}
	if tokenName == "usdt" || tokenName == "usdc" || tokenName == "busd" {
		return 1.00
	}
	pairToken2Usdt := tokenName + "_usdt"

	price, err := requestPrice(pairToken2Usdt)
	if err != nil {
		return 0.00
	}
	return price
}

func requestPrice(pairToken string) (float64, error) {
	var gateIoPrice float64 = 0.00
	var priceInfo model.TokenPriceInfo
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr, Timeout: time.Second * 10}
	var resp *http.Response
	resp, err := httpClient.Get(fmt.Sprintf("%s%s", constant.TICKERINFOAPI, pairToken))
	if err != nil {
		return 0.00, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorF("e2042: get gateio price error:%v", err.Error())
		return 0.00, err
	}

	logger.DebugF("priceInfo: %v", string(body))
	if err := json.Unmarshal(body, &priceInfo); err != nil {
		logger.ErrorF("e2043: parse price error:%v", err.Error())
		return 0.00, err
	}
	gateIoPrice, _ = strconv.ParseFloat(priceInfo.Last, 64)
	return gateIoPrice, nil
}

func calculateGasTokenPrice(nodeId, chainId int) string {
	// gas
	gas := "35000000"

	// gas token
	gasToken := network.GetTokenInfo(nodeId, chainId, "USDT")
	gasTokenDecimal := decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(gasToken.TokenDecimal)))

	// native token
	chainInfo := network.GetChainInfo(chainId)
	nativeToken := network.GetTokenInfo(nodeId, chainId, chainInfo.NativeCurrency.Name)
	nativeTokenDecimal := decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(nativeToken.TokenDecimal)))
	nativeTokenPrice := GetTokenPrice(nativeToken.TokenName)

	// client
	client, err := service.GetEvmClient(service.GetRPC(chainId), chainId)
	if err != nil {
		logger.Error("get client error", zap.Any("chainId", chainId))
		return gas
	}

	// gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Error("suggest gas price error", zap.Any("chainId", chainId))
		return gas
	}

	// gas token price
	gasTokenPriceDecimal := decimal.NewFromInt(10).Pow(decimal.NewFromInt(8))
	gasTokenPrice := gasTokenDecimal.
		Mul(decimal.NewFromFloat(nativeTokenPrice)).
		Mul(decimal.NewFromBigInt(gasPrice, 0)).
		Div(nativeTokenDecimal).
		Mul(gasTokenPriceDecimal)
	return gasTokenPrice.BigInt().String()

}
