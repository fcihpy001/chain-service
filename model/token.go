// @Author fcihpy
// @Date 2024/7/25 16:30:00
// @Desc
//

package model

import "math/big"

type AssetInfo struct {
	ChainId   int
	Wallet    string
	Decimal   int
	Balance   *big.Float
	Min       float64
	GasName   string
	TokenName string
}

type TokenPriceInfo struct {
	Elapsed       string `json:"elapsed"`
	Last          string `json:"last"`
	QuoteVolume   string `json:"quoteVolume"`
	BaseVolume    string `json:"baseVolume"`
	HighestBid    string `json:"highestBid"`
	LowestAsk     string `json:"lowestAsk"`
	High24hr      string `json:"high24hr"`
	Low24hr       string `json:"low24hr"`
	PercentChange string `json:"percentChange"`
	Result        string `json:"result"`
}

type NodeInfos struct {
	Total    int        `json:"total"`
	Page     int        `json:"page"`
	Size     int        `json:"size"`
	DataList []NodeInfo `json:"dataList"`
}

type NodeInfo struct {
	Id               int              `json:"id"`
	Name             string           `json:"name"`
	AvailableAmount  string           `json:"availableAmount"`
	NodeOwnerAddress string           `json:"nodeOwnerAddress"`
	NodeVaultInfos   []*NodeVaultInfo `json:"nodeVaultInfos"`
}

type NodeVaultInfo struct {
	ChainId               int           `json:"chainId"`
	NodeId                int           `json:"nodeId"`
	VaultAddress          string        `json:"vaultAddress"`
	CallParamNodeAddress  string        `json:"callParamNodeAddress"`
	CallParamNodeCallData string        `json:"callParamNodeCallData"`
	NodeFees              []*NodeFees   `json:"nodeFees"`
	VaultTokens           []*VaultToken `json:"vaultTokens"`
}

type VaultToken struct {
	Id          int       `json:"id" `
	NodeVaultId int       `json:"nodeVaultId" `
	TokenId     int       `json:"tokenId"`
	SupportIn   string    `json:"supportIn"`
	SupportOut  string    `json:"supportOut"`
	C           string    `json:"c"`
	K           string    `json:"k"`
	TokenInfo   TokenInfo `json:"tokenInfo"`
}

type TokenInfo struct {
	TokenId      int    `json:"id"`
	ChainId      int    `json:"chainId"`
	TokenName    string `json:"tokenName"`
	TokenSymbol  string `json:"tokenSymbol"`
	TokenAddress string `json:"tokenAddress"`
	TokenDecimal int    `json:"tokenDecimal"`
}

type NodeFeesInfos struct {
	NodeFees  `xorm:"extends"`
	TokenInfo TokenWhitelists `json:"tokenInfo"  xorm:"extends"`
}

type NodeFees struct {
	ChainId     int       `json:"chainId"`
	Id          int       `json:"id"`
	NodeId      int       `json:"nodeId"`
	TokenId     int       `json:"tokenId"`
	PriorityFee string    `json:"priorityFee"`
	TokenInfo   TokenInfo `json:"tokenInfo"`
}
type TokenWhitelists struct {
	Id           int    `json:"id"  xorm:"not null pk autoincr INT"`
	ChainId      int    `json:"chainId"  xorm:"not null INT"`
	TokenAddress string `json:"tokenAddress"  xorm:"not null CHAR(42)"`
	TokenDecimal int    `json:"tokenDecimal"  xorm:"not null INT"`
	TokenName    string `json:"tokenName"  xorm:"not null VARCHAR(20)"`
	TokenSymbol  string `json:"tokenSymbol"  xorm:"not null VARCHAR(20)"`
	TokenClassId int    `json:"tokenClassId"`
}

type NodeKCSingleChain struct {
	ChainId     int
	C           string `json:"c" `
	K           string `json:"k" `
	PriorityFee string `json:"priorityFee" `
}
