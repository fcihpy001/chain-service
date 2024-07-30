// @Author fcihpy
// @Date 2024/7/29 18:52:00
// @Desc
//

package service

import (
	"errors"
	"github.com/DappOSDao/node-service/model"
	"github.com/DappOSDao/node-service/network"
	"github.com/fcihpy001/chain-service/logger"
	"time"
)

func checkExecutePaysState(orderHash string) (int, int) {
	orderInfo, err := network.GetOrderInfo(orderHash)
	if err != nil {
		return 0, 0
	}
	payState := 0
	vwState := 0
	for _, tx := range orderInfo.Txs {
		if len(tx.ExecutePays) > 0 {
			for _, pay := range tx.ExecutePays {
				if pay.State == 3 || pay.State == 5 {
					payState = pay.State
				}
				if len(pay.SrcHash) > 0 {
					logger.DebugF("get src chain hash: %s", pay.SrcHash)
				}
			}
			if tx.Vw != nil {
				vwState = tx.Vw.State
			}
		}
	}
	return payState, vwState
}

func RetriesCheckTxsSate(orderHash string) (isSuccess bool, err error) {
	retries := 1000
	for i := 0; i < retries; i++ {

		payState, _ := checkExecutePaysState(orderHash)
		if payState == 3 {
			return true, nil
		}
		if payState == 5 {
			return false, nil
		}
		time.Sleep(time.Second * 5)
	}
	return false, errors.New("check tx state error")
}

// todo: verify dst chain execute receipt

func VerifyTransaction(order model.Order) error {
	for _, tx := range order.Txs {
		if tx.Vw != nil {
			return nil
		}
	}
	return nil
}
