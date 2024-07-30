// @Author fcihpy
// @Date 2024/7/19 15:10:00
// @Desc
//

package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"math"
	"strings"
)

func GetMultiCallABI() string {
	return `[
  {
    "inputs": [
      {
        "components": [
          {
            "internalType": "address",
            "name": "target",
            "type": "address"
          },
          {
            "internalType": "bytes",
            "name": "callData",
            "type": "bytes"
          }
        ],
        "internalType": "struct Multicall3.Call[]",
        "name": "calls",
        "type": "tuple[]"
      }
    ],
    "name": "aggregate",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "blockNumber",
        "type": "uint256"
      },
      {
        "internalType": "bytes[]",
        "name": "returnData",
        "type": "bytes[]"
      }
    ],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "components": [
          {
            "internalType": "address",
            "name": "target",
            "type": "address"
          },
          {
            "internalType": "bool",
            "name": "allowFailure",
            "type": "bool"
          },
          {
            "internalType": "bytes",
            "name": "callData",
            "type": "bytes"
          }
        ],
        "internalType": "struct Multicall3.Call3[]",
        "name": "calls",
        "type": "tuple[]"
      }
    ],
    "name": "aggregate3",
    "outputs": [
      {
        "components": [
          {
            "internalType": "bool",
            "name": "success",
            "type": "bool"
          },
          {
            "internalType": "bytes",
            "name": "returnData",
            "type": "bytes"
          }
        ],
        "internalType": "struct Multicall3.Result[]",
        "name": "returnData",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "components": [
          {
            "internalType": "address",
            "name": "target",
            "type": "address"
          },
          {
            "internalType": "bool",
            "name": "allowFailure",
            "type": "bool"
          },
          {
            "internalType": "uint256",
            "name": "value",
            "type": "uint256"
          },
          {
            "internalType": "bytes",
            "name": "callData",
            "type": "bytes"
          }
        ],
        "internalType": "struct Multicall3.Call3Value[]",
        "name": "calls",
        "type": "tuple[]"
      }
    ],
    "name": "aggregate3Value",
    "outputs": [
      {
        "components": [
          {
            "internalType": "bool",
            "name": "success",
            "type": "bool"
          },
          {
            "internalType": "bytes",
            "name": "returnData",
            "type": "bytes"
          }
        ],
        "internalType": "struct Multicall3.Result[]",
        "name": "returnData",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "components": [
          {
            "internalType": "address",
            "name": "target",
            "type": "address"
          },
          {
            "internalType": "bytes",
            "name": "callData",
            "type": "bytes"
          }
        ],
        "internalType": "struct Multicall3.Call[]",
        "name": "calls",
        "type": "tuple[]"
      }
    ],
    "name": "blockAndAggregate",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "blockNumber",
        "type": "uint256"
      },
      {
        "internalType": "bytes32",
        "name": "blockHash",
        "type": "bytes32"
      },
      {
        "components": [
          {
            "internalType": "bool",
            "name": "success",
            "type": "bool"
          },
          {
            "internalType": "bytes",
            "name": "returnData",
            "type": "bytes"
          }
        ],
        "internalType": "struct Multicall3.Result[]",
        "name": "returnData",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getBasefee",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "basefee",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "blockNumber",
        "type": "uint256"
      }
    ],
    "name": "getBlockHash",
    "outputs": [
      {
        "internalType": "bytes32",
        "name": "blockHash",
        "type": "bytes32"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getBlockNumber",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "blockNumber",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getChainId",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "chainid",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getCurrentBlockCoinbase",
    "outputs": [
      {
        "internalType": "address",
        "name": "coinbase",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getCurrentBlockDifficulty",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "difficulty",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getCurrentBlockGasLimit",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "gaslimit",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getCurrentBlockTimestamp",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "timestamp",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "addr",
        "type": "address"
      }
    ],
    "name": "getEthBalance",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "balance",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getLastBlockHash",
    "outputs": [
      {
        "internalType": "bytes32",
        "name": "blockHash",
        "type": "bytes32"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bool",
        "name": "requireSuccess",
        "type": "bool"
      },
      {
        "components": [
          {
            "internalType": "address",
            "name": "target",
            "type": "address"
          },
          {
            "internalType": "bytes",
            "name": "callData",
            "type": "bytes"
          }
        ],
        "internalType": "struct Multicall3.Call[]",
        "name": "calls",
        "type": "tuple[]"
      }
    ],
    "name": "tryAggregate",
    "outputs": [
      {
        "components": [
          {
            "internalType": "bool",
            "name": "success",
            "type": "bool"
          },
          {
            "internalType": "bytes",
            "name": "returnData",
            "type": "bytes"
          }
        ],
        "internalType": "struct Multicall3.Result[]",
        "name": "returnData",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bool",
        "name": "requireSuccess",
        "type": "bool"
      },
      {
        "components": [
          {
            "internalType": "address",
            "name": "target",
            "type": "address"
          },
          {
            "internalType": "bytes",
            "name": "callData",
            "type": "bytes"
          }
        ],
        "internalType": "struct Multicall3.Call[]",
        "name": "calls",
        "type": "tuple[]"
      }
    ],
    "name": "tryBlockAndAggregate",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "blockNumber",
        "type": "uint256"
      },
      {
        "internalType": "bytes32",
        "name": "blockHash",
        "type": "bytes32"
      },
      {
        "components": [
          {
            "internalType": "bool",
            "name": "success",
            "type": "bool"
          },
          {
            "internalType": "bytes",
            "name": "returnData",
            "type": "bytes"
          }
        ],
        "internalType": "struct Multicall3.Result[]",
        "name": "returnData",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "payable",
    "type": "function"
  }
]`
}

func GetERC20ABI() string {
	return `[
			{
				"anonymous": false,
				"inputs": [
					{
						"indexed": true,
						"internalType": "address",
						"name": "owner",
						"type": "address"
					},
					{
						"indexed": true,
						"internalType": "address",
						"name": "spender",
						"type": "address"
					},
					{
						"indexed": false,
						"internalType": "uint256",
						"name": "value",
						"type": "uint256"
					}
				],
				"name": "Approval",
				"type": "event"
			},
			{
				"anonymous": false,
				"inputs": [
					{
						"indexed": true,
						"internalType": "address",
						"name": "from",
						"type": "address"
					},
					{
						"indexed": true,
						"internalType": "address",
						"name": "to",
						"type": "address"
					},
					{
						"indexed": false,
						"internalType": "uint256",
						"name": "value",
						"type": "uint256"
					}
				],
				"name": "Transfer",
				"type": "event"
			},
			{
				"constant": true,
				"inputs": [
					{
						"internalType": "address",
						"name": "owner",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "spender",
						"type": "address"
					}
				],
				"name": "allowance",
				"outputs": [
					{
						"internalType": "uint256",
						"name": "",
						"type": "uint256"
					}
				],
				"payable": false,
				"stateMutability": "view",
				"type": "function"
			},
			{
				"constant": false,
				"inputs": [
					{
						"internalType": "address",
						"name": "spender",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "amount",
						"type": "uint256"
					}
				],
				"name": "approve",
				"outputs": [
					{
						"internalType": "bool",
						"name": "",
						"type": "bool"
					}
				],
				"payable": false,
				"stateMutability": "nonpayable",
				"type": "function"
			},
			{
				"constant": true,
				"inputs": [
					{
						"internalType": "address",
						"name": "account",
						"type": "address"
					}
				],
				"name": "balanceOf",
				"outputs": [
					{
						"internalType": "uint256",
						"name": "",
						"type": "uint256"
					}
				],
				"payable": false,
				"stateMutability": "view",
				"type": "function"
			},
			{
				"constant": true,
				"inputs": [],
				"name": "totalSupply",
				"outputs": [
					{
						"internalType": "uint256",
						"name": "",
						"type": "uint256"
					}
				],
				"payable": false,
				"stateMutability": "view",
				"type": "function"
			},
			{
				"constant": false,
				"inputs": [
					{
						"internalType": "address",
						"name": "recipient",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "amount",
						"type": "uint256"
					}
				],
				"name": "transfer",
				"outputs": [
					{
						"internalType": "bool",
						"name": "",
						"type": "bool"
					}
				],
				"payable": false,
				"stateMutability": "nonpayable",
				"type": "function"
			},
			{
				"constant": false,
				"inputs": [
					{
						"internalType": "address",
						"name": "sender",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "recipient",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "amount",
						"type": "uint256"
					}
				],
				"name": "transferFrom",
				"outputs": [
					{
						"internalType": "bool",
						"name": "",
						"type": "bool"
					}
				],
				"payable": false,
				"stateMutability": "nonpayable",
				"type": "function"
			}
		]`
}

func GetPayDbABI() string {
	return `[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "executor",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "receiver",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "amountOut",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "tokenOut",
				"type": "address"
			}
		],
		"name": "IsolateOrderExecuted",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "node",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "payOrderId",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "cancelOperator",
				"type": "address"
			}
		],
		"name": "OrderCancelled",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "node",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "payOrderId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "sender",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "receiver",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "amountIn",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "tokenIn",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "amountOut",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "tokenOut",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "code",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "bytes32",
				"name": "wfHash",
				"type": "bytes32"
			}
		],
		"name": "OrderCreated",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "executor",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "payOrderId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "receiver",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "amountOut",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "tokenOut",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "code",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "bytes32",
				"name": "wfHash",
				"type": "bytes32"
			}
		],
		"name": "OrderExecuted",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "previousOwner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "OwnershipTransferred",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "reason",
				"type": "string"
			}
		],
		"name": "VWFailedReason",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "vwmanager",
				"type": "address"
			}
		],
		"name": "VWManagerSet",
		"type": "event"
	},
	{
		"stateMutability": "payable",
		"type": "fallback"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "sender",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "receiver",
				"type": "address"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "amountIn",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "amountOut",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "payOrderId",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "bridgeFee",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "tokenIn",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "tokenOut",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "node",
						"type": "address"
					}
				],
				"internalType": "struct IPayDB.CreatePayOrderParam[]",
				"name": "cparams",
				"type": "tuple[]"
			},
			{
				"internalType": "bytes32[]",
				"name": "workFlowHashs",
				"type": "bytes32[]"
			}
		],
		"name": "cancelOrderETH",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "orderOwner",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "receiver",
				"type": "address"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "amountIn",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "amountOut",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "payOrderId",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "bridgeFee",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "tokenIn",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "tokenOut",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "node",
						"type": "address"
					}
				],
				"internalType": "struct IPayDB.CreatePayOrderParam[]",
				"name": "cparams",
				"type": "tuple[]"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "code",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "service",
						"type": "address"
					},
					{
						"internalType": "bytes32",
						"name": "dataHash",
						"type": "bytes32"
					}
				],
				"internalType": "struct IPayDB.VwOrderDetail",
				"name": "vwDetail",
				"type": "tuple"
			},
			{
				"components": [
					{
						"internalType": "address",
						"name": "node",
						"type": "address"
					},
					{
						"internalType": "bytes",
						"name": "nodeCallData",
						"type": "bytes"
					}
				],
				"internalType": "struct IPayDB.CallParam",
				"name": "callParam",
				"type": "tuple"
			}
		],
		"name": "createSrcOrder",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "orderOwner",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "receiver",
				"type": "address"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "amountIn",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "amountOut",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "payOrderId",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "bridgeFee",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "tokenIn",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "tokenOut",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "node",
						"type": "address"
					}
				],
				"internalType": "struct IPayDB.CreatePayOrderParam[]",
				"name": "cparams",
				"type": "tuple[]"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "code",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "service",
						"type": "address"
					},
					{
						"internalType": "bytes32",
						"name": "dataHash",
						"type": "bytes32"
					}
				],
				"internalType": "struct IPayDB.VwOrderDetail",
				"name": "vwDetail",
				"type": "tuple"
			},
			{
				"components": [
					{
						"internalType": "address",
						"name": "node",
						"type": "address"
					},
					{
						"internalType": "bytes",
						"name": "nodeCallData",
						"type": "bytes"
					}
				],
				"internalType": "struct IPayDB.CallParam",
				"name": "callParam",
				"type": "tuple"
			}
		],
		"name": "createSrcOrderETH",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"name": "dstOrder",
		"outputs": [
			{
				"internalType": "bytes32",
				"name": "",
				"type": "bytes32"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "orderOwner",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "receiver",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "amountOut",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "tokenOut",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "payOrderId",
						"type": "uint256"
					}
				],
				"internalType": "struct IPayDB.ExePayOrderParam[]",
				"name": "eparams",
				"type": "tuple[]"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "code",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gasTokenPrice",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "priorityFee",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gasLimit",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "manager",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "service",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "gasToken",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "feeReceiver",
						"type": "address"
					},
					{
						"internalType": "bool",
						"name": "isGateway",
						"type": "bool"
					},
					{
						"internalType": "bytes",
						"name": "data",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "serviceSignature",
						"type": "bytes"
					},
					{
						"internalType": "bytes32[]",
						"name": "proof",
						"type": "bytes32[]"
					}
				],
				"internalType": "struct IVWManager.VWExecuteParam",
				"name": "vwExeParam",
				"type": "tuple"
			}
		],
		"name": "executeDstOrderETH",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "receiver",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "amountOut",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "tokenOut",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "payOrderId",
						"type": "uint256"
					}
				],
				"internalType": "struct IPayDB.ExePayOrderParam[]",
				"name": "eparams",
				"type": "tuple[]"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "code",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gasTokenPrice",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "priorityFee",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gasLimit",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "manager",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "service",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "gasToken",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "feeReceiver",
						"type": "address"
					},
					{
						"internalType": "bool",
						"name": "isGateway",
						"type": "bool"
					},
					{
						"internalType": "bytes",
						"name": "data",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "serviceSignature",
						"type": "bytes"
					},
					{
						"internalType": "bytes32[]",
						"name": "proof",
						"type": "bytes32[]"
					}
				],
				"internalType": "struct IVWManager.VWExecuteParam",
				"name": "vwExeParam",
				"type": "tuple"
			}
		],
		"name": "executeIsolateOrder",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "receiver",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "amountOut",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "tokenOut",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "payOrderId",
						"type": "uint256"
					}
				],
				"internalType": "struct IPayDB.ExePayOrderParam[]",
				"name": "eparams",
				"type": "tuple[]"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "code",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gasTokenPrice",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "priorityFee",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gasLimit",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "manager",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "service",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "gasToken",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "feeReceiver",
						"type": "address"
					},
					{
						"internalType": "bool",
						"name": "isGateway",
						"type": "bool"
					},
					{
						"internalType": "bytes",
						"name": "data",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "serviceSignature",
						"type": "bytes"
					},
					{
						"internalType": "bytes32[]",
						"name": "proof",
						"type": "bytes32[]"
					}
				],
				"internalType": "struct IVWManager.VWExecuteParam",
				"name": "vwExeParam",
				"type": "tuple"
			}
		],
		"name": "executeIsolateOrderETH",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "renounceOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"name": "srcOrder",
		"outputs": [
			{
				"internalType": "address",
				"name": "node",
				"type": "address"
			},
			{
				"internalType": "uint96",
				"name": "status",
				"type": "uint96"
			},
			{
				"internalType": "bytes32",
				"name": "orderDataHash",
				"type": "bytes32"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "transferOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "orderOwner",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "receiver",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "amountOut",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "tokenOut",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "payOrderId",
						"type": "uint256"
					}
				],
				"internalType": "struct IPayDB.ExePayOrderParam[]",
				"name": "eparams",
				"type": "tuple[]"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "code",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gasTokenPrice",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "priorityFee",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gasLimit",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "manager",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "service",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "gasToken",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "feeReceiver",
						"type": "address"
					},
					{
						"internalType": "bool",
						"name": "isGateway",
						"type": "bool"
					},
					{
						"internalType": "bytes",
						"name": "data",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "serviceSignature",
						"type": "bytes"
					},
					{
						"internalType": "bytes32[]",
						"name": "proof",
						"type": "bytes32[]"
					}
				],
				"internalType": "struct IVWManager.VWExecuteParam",
				"name": "vwExeParam",
				"type": "tuple"
			}
		],
		"name": "tryExecuteDstOrderETH",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"stateMutability": "payable",
		"type": "receive"
	}
]`
}

func GetIntentDBABI() string {
	return `[
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "node",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "payOrderId",
        "type": "uint256"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "cancelOperator",
        "type": "address"
      }
    ],
    "name": "IntentOrderCancelled",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "owner",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "node",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "payOrderId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "contract IIntentToken",
        "name": "srcIntentToken",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "sender",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "receiver",
        "type": "address"
      },
      {
        "components": [
          {
            "internalType": "address",
            "name": "underlyingAssets",
            "type": "address"
          },
          {
            "internalType": "uint256",
            "name": "underlyingAssetsAmount",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "tokenDecimals",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "bridgeFee",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "priceInBase",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "priceDecimals",
            "type": "uint256"
          }
        ],
        "indexed": false,
        "internalType": "struct IIntentAssetDB.UnderlyingAssetsInfo",
        "name": "underlyingAssetsInfo",
        "type": "tuple"
      }
    ],
    "name": "IntentOrderCreated",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "executor",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "payOrderId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "owner",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "receiver",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "vwManager",
        "type": "address"
      },
      {
        "components": [
          {
            "internalType": "contract IIntentToken",
            "name": "intentToken",
            "type": "address"
          },
          {
            "internalType": "uint256",
            "name": "intentTokenAmount",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "intentTokenDecimals",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "baseAssetsRate",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "BASE_ASSET_RATE_BASE",
            "type": "uint256"
          }
        ],
        "indexed": false,
        "internalType": "struct IIntentAssetDB.IntentTokenInfo",
        "name": "intentTokenInfo",
        "type": "tuple"
      },
      {
        "components": [
          {
            "internalType": "address",
            "name": "underlyingAssets",
            "type": "address"
          },
          {
            "internalType": "uint256",
            "name": "underlyingAssetsAmount",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "tokenDecimals",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "bridgeFee",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "priceInBase",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "priceDecimals",
            "type": "uint256"
          }
        ],
        "indexed": false,
        "internalType": "struct IIntentAssetDB.UnderlyingAssetsInfo",
        "name": "underlyingAssetsInfo",
        "type": "tuple"
      }
    ],
    "name": "IntentOrderExecuted",
    "type": "event"
  },
  {
    "inputs": [
      {
        "components": [
          {
            "internalType": "address",
            "name": "underlyingAssets",
            "type": "address"
          },
          {
            "internalType": "uint256",
            "name": "underlyingAssetsAmount",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "tokenDecimals",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "bridgeFee",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "priceInBase",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "priceDecimals",
            "type": "uint256"
          }
        ],
        "internalType": "struct IIntentAssetDB.UnderlyingAssetsInfo",
        "name": "underlyingAssetsInfo",
        "type": "tuple"
      },
      {
        "internalType": "uint256",
        "name": "intentTokenDecimals",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "intentTokenBaseRate",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "intentTokenBaseRateBase",
        "type": "uint256"
      }
    ],
    "name": "calculateIntentTokenAmount",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "pure",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "sender",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "receiver",
        "type": "address"
      },
      {
        "components": [
          {
            "internalType": "contract IIntentToken",
            "name": "srcIntentToken",
            "type": "address"
          },
          {
            "internalType": "address",
            "name": "underlyingAssets",
            "type": "address"
          },
          {
            "internalType": "uint256",
            "name": "underlyingAssetsAmount",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "payOrderId",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "bridgeFee",
            "type": "uint256"
          },
          {
            "internalType": "address",
            "name": "node",
            "type": "address"
          }
        ],
        "internalType": "struct IIntentAssetDB.CreateIntentOrderParam[]",
        "name": "cparams",
        "type": "tuple[]"
      }
    ],
    "name": "cancelOrderETH",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "orderOwner",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "receiver",
        "type": "address"
      },
      {
        "components": [
          {
            "internalType": "contract IIntentToken",
            "name": "srcIntentToken",
            "type": "address"
          },
          {
            "internalType": "address",
            "name": "underlyingAssets",
            "type": "address"
          },
          {
            "internalType": "uint256",
            "name": "underlyingAssetsAmount",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "payOrderId",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "bridgeFee",
            "type": "uint256"
          },
          {
            "internalType": "address",
            "name": "node",
            "type": "address"
          }
        ],
        "internalType": "struct IIntentAssetDB.CreateIntentOrderParam[]",
        "name": "cparams",
        "type": "tuple[]"
      },
      {
        "components": [
          {
            "internalType": "address",
            "name": "node",
            "type": "address"
          },
          {
            "internalType": "bytes",
            "name": "nodeCallData",
            "type": "bytes"
          }
        ],
        "internalType": "struct IIntentAssetDB.CallParam",
        "name": "callParam",
        "type": "tuple"
      }
    ],
    "name": "createSrcIntentOrderETH",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "name": "dstIntentOrder",
    "outputs": [
      {
        "internalType": "bytes32",
        "name": "",
        "type": "bytes32"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "orderOwner",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "receiver",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "vwManager",
        "type": "address"
      },
      {
        "components": [
          {
            "internalType": "address",
            "name": "underlyingAssets",
            "type": "address"
          },
          {
            "internalType": "uint256",
            "name": "underlyingAssetsAmount",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "tokenDecimals",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "bridgeFee",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "priceInBase",
            "type": "uint256"
          },
          {
            "internalType": "uint256",
            "name": "priceDecimals",
            "type": "uint256"
          }
        ],
        "internalType": "struct IIntentAssetDB.UnderlyingAssetsInfo[]",
        "name": "underlyingAssetsInfos",
        "type": "tuple[]"
      },
      {
        "components": [
          {
            "internalType": "contract IIntentToken",
            "name": "dstIntentToken",
            "type": "address"
          },
          {
            "internalType": "uint256",
            "name": "payOrderId",
            "type": "uint256"
          }
        ],
        "internalType": "struct IIntentAssetDB.ExeIntentOrderParam[]",
        "name": "eparams",
        "type": "tuple[]"
      }
    ],
    "name": "executeDstIntentOrder",
    "outputs": [
      {
        "internalType": "uint256[]",
        "name": "intentTokenAmounts",
        "type": "uint256[]"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "name": "srcIntentOrder",
    "outputs": [
      {
        "internalType": "address",
        "name": "node",
        "type": "address"
      },
      {
        "internalType": "uint64",
        "name": "status",
        "type": "uint64"
      },
      {
        "internalType": "uint32",
        "name": "createTime",
        "type": "uint32"
      },
      {
        "internalType": "bytes32",
        "name": "orderDataHash",
        "type": "bytes32"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
]`
}

func GetVwABI() string {
	return `[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "srcChainID",
				"type": "uint256"
			}
		],
		"name": "DomainSeparatorCanceled",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "srcChainID",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "verifyContract",
				"type": "address"
			}
		],
		"name": "DomainSeparatorConfiged",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "address",
				"name": "deployer",
				"type": "address"
			}
		],
		"name": "Initialized",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "previousOwner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "OwnershipTransferred",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "reseter",
				"type": "address"
			}
		],
		"name": "ResetterChanged",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "bool",
				"name": "protocolFeeOpened",
				"type": "bool"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "feeVault",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "feeProportion",
				"type": "uint256"
			}
		],
		"name": "SetFee",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "code",
				"type": "uint256"
			}
		],
		"name": "TxCanceled",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "code",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "bytes32",
				"name": "rootHash",
				"type": "bytes32"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "resCode",
				"type": "uint256"
			}
		],
		"name": "TxExecuted",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "owner",
				"type": "address"
			}
		],
		"name": "VWCreated",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "previousOwner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "VWOwnerChanged",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "code",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "resetter",
				"type": "address"
			},
			{
				"internalType": "bool",
				"name": "approved",
				"type": "bool"
			},
			{
				"components": [
					{
						"internalType": "address",
						"name": "gasToken",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "gasTokenPrice",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "priorityFee",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gasLimit",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "feeReceiver",
						"type": "address"
					}
				],
				"internalType": "struct VWManagerService.FeeParam",
				"name": "fParam",
				"type": "tuple"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "approveResetter",
		"outputs": [
			{
				"internalType": "bytes32",
				"name": "dataHash",
				"type": "bytes32"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"name": "approvedResetter",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "code",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "codeToCancel",
				"type": "uint256"
			},
			{
				"components": [
					{
						"internalType": "address",
						"name": "gasToken",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "gasTokenPrice",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "priorityFee",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gasLimit",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "feeReceiver",
						"type": "address"
					}
				],
				"internalType": "struct VWManagerService.FeeParam",
				"name": "fParam",
				"type": "tuple"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "cancelTx",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "code",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			},
			{
				"components": [
					{
						"internalType": "address",
						"name": "gasToken",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "gasTokenPrice",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "priorityFee",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gasLimit",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "feeReceiver",
						"type": "address"
					}
				],
				"internalType": "struct VWManagerService.FeeParam",
				"name": "fParam",
				"type": "tuple"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "changeOwner",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "bool",
				"name": "_protocolFeeOpened",
				"type": "bool"
			},
			{
				"internalType": "address",
				"name": "_feeVault",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "_feeProportion",
				"type": "uint256"
			}
		],
		"name": "configFee",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_chainId",
				"type": "uint256"
			},
			{
				"internalType": "bool",
				"name": "_support",
				"type": "bool"
			},
			{
				"internalType": "address",
				"name": "_verifyingContract",
				"type": "address"
			}
		],
		"name": "configSrcChain",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "owner",
				"type": "address"
			}
		],
		"name": "createWallet",
		"outputs": [
			{
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "deployer",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"name": "domainSeparator",
		"outputs": [
			{
				"internalType": "bytes32",
				"name": "",
				"type": "bytes32"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "code",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gasTokenPrice",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "priorityFee",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gasLimit",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "manager",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "service",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "gasToken",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "feeReceiver",
						"type": "address"
					},
					{
						"internalType": "bool",
						"name": "isGateway",
						"type": "bool"
					},
					{
						"internalType": "bytes",
						"name": "data",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "serviceSignature",
						"type": "bytes"
					},
					{
						"internalType": "bytes32[]",
						"name": "proof",
						"type": "bytes32[]"
					}
				],
				"internalType": "struct IVWManager.VWExecuteParam",
				"name": "vweParam",
				"type": "tuple"
			}
		],
		"name": "execute",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "resCode",
				"type": "uint256"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "feeProportion",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "feeVault",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_deployer",
				"type": "address"
			}
		],
		"name": "initialize",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"name": "ownerWallet",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "protocolFeeOpened",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "renounceOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "code",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			},
			{
				"components": [
					{
						"internalType": "address",
						"name": "gasToken",
						"type": "address"
					},
					{
						"internalType": "uint256",
						"name": "gasTokenPrice",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "priorityFee",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gasLimit",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "feeReceiver",
						"type": "address"
					}
				],
				"internalType": "struct VWManagerService.FeeParam",
				"name": "fParam",
				"type": "tuple"
			},
			{
				"internalType": "bytes",
				"name": "data",
				"type": "bytes"
			}
		],
		"name": "resetOwner",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"name": "result",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "transferOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"internalType": "bytes32",
				"name": "digest",
				"type": "bytes32"
			},
			{
				"internalType": "bytes",
				"name": "signature",
				"type": "bytes"
			}
		],
		"name": "verifyEIP1271Signature",
		"outputs": [
			{
				"internalType": "bool",
				"name": "isValid",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "resCode",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "wallet",
				"type": "address"
			},
			{
				"components": [
					{
						"internalType": "uint256",
						"name": "code",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gasTokenPrice",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "priorityFee",
						"type": "uint256"
					},
					{
						"internalType": "uint256",
						"name": "gasLimit",
						"type": "uint256"
					},
					{
						"internalType": "address",
						"name": "manager",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "service",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "gasToken",
						"type": "address"
					},
					{
						"internalType": "address",
						"name": "feeReceiver",
						"type": "address"
					},
					{
						"internalType": "bool",
						"name": "isGateway",
						"type": "bool"
					},
					{
						"internalType": "bytes",
						"name": "data",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "serviceSignature",
						"type": "bytes"
					},
					{
						"internalType": "bytes32[]",
						"name": "proof",
						"type": "bytes32[]"
					}
				],
				"internalType": "struct IVWManager.VWExecuteParam",
				"name": "vweParam",
				"type": "tuple"
			}
		],
		"name": "verifyProof",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "volatileService",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"name": "walletOwner",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]`
}

//func PackCallData(
//	parsedABI abi.ABI,
//	contractAddr common.Address,
//	methodName string,
//	args ...interface{},
//) model.CalInfo {
//
//	callData, err := parsedABI.Pack(methodName, args...)
//	if err != nil {
//		return model.CalInfo{}
//	}
//	return model.CalInfo{
//		Target:   contractAddr,
//		CallData: callData,
//	}
//}

func GetTransferCallData(receiver common.Address, amount float64, tokenDecimal int) ([]byte, error) {
	contractABI, err := abi.JSON(strings.NewReader(GetERC20ABI()))
	if err != nil {
		return nil, err
	}
	_tokenDecimal := math.Pow10(tokenDecimal)
	amountDecimal := decimal.NewFromFloat(amount).Mul(decimal.NewFromFloat(_tokenDecimal))
	amountBIg := amountDecimal.BigInt()
	callData, err := contractABI.Pack("transfer", receiver, amountBIg)
	if err != nil {
		return nil, err
	}
	return callData, nil

}
