// @Author fcihpy
// @Date 2024/7/25 16:46:00
// @Desc
//

package constant

const (
	TokenUSDT  string = "USDT"
	TokenUSDC  string = "USDC"
	TokenETH   string = "ETH"
	TokenMatic string = "MATIC"
	TokenAvax  string = "AVAX"
	TokenBNB   string = "BNB"
)

const (
	MultiCall           string = "0xcA11bde05977b3631167028862bE2a173976CA11"
	MethodBalanceOf     string = "balanceOf"
	MethodGetEthBalance string = "getEthBalance"
	MethodAggregate     string = "aggregate"
	ZeroAddr                   = "0x0000000000000000000000000000000000000000"
	Zero42Addr                 = "0x4200000000000000000000000000000000000006"
	F64                        = "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	X1901                      = "\x19\x01"
	MoreZero                   = "00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001"
	CHAINS                     = "CHAINS"
	NODES                      = "NODES"
	VIRTUALWALLET              = "VirtualWallet"
	TICKERINFOAPI              = "https://data.gateapi.io/api2/1/ticker/"
)

type TaskType string

const (
	TaskTypePayDB    TaskType = "payDB"
	TaskTypeIntentDB TaskType = "intentDB"
)
