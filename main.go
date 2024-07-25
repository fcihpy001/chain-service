// @Author fcihpy
// @Date 2024/7/25 14:48:00
// @Desc
//

package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fcihpy001/chain-service/service"
)

func main() {
	fmt.Println("hello world")

	client, err := service.Dial(context.Background(), 10, []string{"https://opt-mainnet.nodereal.io/v1/44acc0b89a0c4643be14ed0b58307a4c"})
	if err != nil {
		panic(err)
	}
	balance, err := client.RawClients[0].BalanceAt(context.Background(), common.HexToAddress("0x178ad9283bba524d2e10ad1102a079144bc569e0"), nil)
	if err != nil {
		return
	}
	fmt.Println(balance)
}
