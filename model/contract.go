// @Author fcihpy
// @Date 2024/7/25 16:44:00
// @Desc
//

package model

import "github.com/ethereum/go-ethereum/common"

type CalInfo struct {
	Target   common.Address
	CallData []byte
}
