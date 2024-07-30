// @Author fcihpy
// @Date 2024/7/25 14:48:00
// @Desc
//

package main

import (
	"github.com/fcihpy001/chain-service/constant"
	"github.com/fcihpy001/chain-service/model"
	"github.com/fcihpy001/chain-service/swap"
)

func main() {
	functionTest()
}

func functionTest() {
	task1 := model.Task{
		Type:       constant.TaskTypePayDB,
		NodeId:     26,
		SrcChainId: 56,
		DstChainId: 10,
		Amounts:    []float64{0.2},
		TokenIns:   []string{"USDT"},
		TokenOuts:  []string{"USDT"},
		HasSrcVw:   false,
		HasDstVw:   false,
	}
	_, err := swap.PayDBSwap([]model.Task{task1}, false)
	if err != nil {
		return
	}
}
