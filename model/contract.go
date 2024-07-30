// @Author fcihpy
// @Date 2024/7/25 16:44:00
// @Desc
//

package model

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/fcihpy001/chain-service/contract"
	"math/big"
)

type CalInfo struct {
	Target   common.Address
	CallData []byte
}

type ContractParam struct {
	CreatePays       []contract.IPayDBCreatePayOrderParam
	Vw               contract.IPayDBVwOrderDetail
	CallData         contract.IPayDBCallParam
	IntentCreatePays []contract.IIntentDBCreateIntentOrderParam
	IntentCallParam  contract.IIntentDBCallParam
	Owner            common.Address
	Wallet           common.Address
	Receiver         common.Address
	PayDB            common.Address
	Amount           *big.Int
	ChainId          int
}
