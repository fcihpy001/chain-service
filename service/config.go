// @Author fcihpy
// @Date 2024/7/25 16:33:00
// @Desc
//

package service

var (
	chainUSDT   map[int]string
	chainRPC    map[int]string
	accounts    []string
	nodeAddrMap map[int]string
)

func init() {
	chainUSDT = map[int]string{
		1:     "0xdac17f958d2ee523a2206206994597c13d831ec7",
		10:    "0x94b008aa00579c1307b0ef2c499ad98a8ce58e58",
		56:    "0x55d398326f99059ff775485246999027b3197955",
		137:   "0xc2132d05d31c914a87c6611c10748aeb04b58e8f",
		169:   "0xf417f5a458ec102b90352f697d6e2ac3a3d2851f",
		324:   "0x493257fd37edb34451f62edf8d2a0c418852ba4c",
		42161: "0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9",
		43114: "0x9702230a8ea53601f5cd2dc00fdbc13d4df4a8c7",
	}

	chainRPC = map[int]string{
		1:     "https://eth-mainnet.nodereal.io/v1/44acc0b89a0c4643be14ed0b58307a4c",
		10:    "https://opt-mainnet.nodereal.io/v1/44acc0b89a0c4643be14ed0b58307a4c",
		56:    "https://bsc-mainnet.nodereal.io/v1/44acc0b89a0c4643be14ed0b58307a4c",
		137:   "https://polygon-mainnet.nodereal.io/v1/44acc0b89a0c4643be14ed0b58307a4c",
		169:   "https://manta-pacific.drpc.org",
		324:   "https://zksync-mainnet.infura.io/v3/bd4216db4b0e471088d366a9eb68fa02",
		42161: "https://rpc.ankr.com/arbitrum",
		43114: "https://avalanche-mainnet.infura.io/v3/bd4216db4b0e471088d366a9eb68fa02",
	}

	nodeAddrMap = make(map[int]string)
	nodeAddrMap[24] = "0x178ad9283bba524d2e10ad1102a079144bc569e0"
	nodeAddrMap[25] = "0xaee4f643c3f75757664f7e5aadfb4a1816ac6f10"
	nodeAddrMap[26] = "0x60ab2f89f3aa9d4255d696771c9f017f190f3eca"
	nodeAddrMap[27] = "0x1a1423edfb518959856450339a1aaa3527ad74c7"
	nodeAddrMap[28] = "0xa7f5a210c92a82d2a05e0072d8dccd33acd0d092"
	nodeAddrMap[29] = "0x81ddd9e8c83ef5b2cb643b1076238e288a6967fb"

	accounts = []string{
		"0x178ad9283bba524d2e10ad1102a079144bc569e0",
		"0xaee4f643c3f75757664f7e5aadfb4a1816ac6f10",
		"0x60ab2f89f3aa9d4255d696771c9f017f190f3eca",
		"0x1a1423edfb518959856450339a1aaa3527ad74c7",
		"0xa7f5a210c92a82d2a05e0072d8dccd33acd0d092",
		"0x81ddd9e8c83ef5b2cb643b1076238e288a6967fb",
		"0x6B0C341E68C75dD1e60733b62B130d192f5A136b",
	}
}

func getUSDT(chainId int) string {
	return chainUSDT[chainId]
}

func getRPC(chainId int) string {
	return chainRPC[chainId]
}

func GetAccounts() []string {
	return accounts
}

func GetGasWater(chainId int) WaterInfo {
	switch chainId {
	case 1, 10, 169, 42161:
		return WaterInfo{
			GasName: "ETH",
			Min:     0.02,
		}
	case 56:
		return WaterInfo{
			GasName: "BNB",
			Min:     0.1,
		}
	case 137:
		return WaterInfo{
			GasName: "MATIC",
			Min:     20,
		}
	case 43114:
		return WaterInfo{
			GasName: "AVAX",
			Min:     0.1,
		}
	default:
		return WaterInfo{}
	}

}

type WaterInfo struct {
	GasName string
	Min     float64
}

func GetNodeAddr(nodeId int) string {
	return nodeAddrMap[nodeId]
}
