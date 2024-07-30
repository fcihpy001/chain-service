// @Author fcihpy
// @Date 2024/7/29 17:40:00
// @Desc
//

package model

import (
	"crypto/ecdsa"
	"github.com/fcihpy001/chain-service/constant"
)

type Account struct {
	Address    string
	Vw         string
	MantaVw    string
	ZkSyncVw   string
	PrivateKey *ecdsa.PrivateKey
}

type Task struct {
	Type       constant.TaskType `json:"type"`
	NodeId     int               `json:"node_id"`
	SrcChainId int               `json:"src_chain_id"`
	DstChainId int               `json:"dst_chain_id"`
	Amounts    []float64         `json:"amounts"`
	TokenIns   []string          `json:"token_in_s"`
	TokenOuts  []string          `json:"token_out_s"`
	HasSrcVw   bool              `json:"has_src_vw"`
	HasDstVw   bool              `json:"has_dst_vw"`
}
