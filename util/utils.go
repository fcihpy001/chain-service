// @Author fcihpy
// @Date 2024/7/25 17:57:00
// @Desc
//

package util

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/DappOSDao/node-service/model"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fcihpy001/chain-service/constant"
	"math/big"
)

var (
	intentDBNodes = [2]int{21, 22}
	payDBNodes    = []int{2, 14, 17, 18, 24, 25, 26, 27, 28, 29}
)

func ContainNodeId(taskType constant.TaskType, nodeId int) bool {

	if taskType == constant.TaskTypeIntentDB {
		for _, v := range intentDBNodes {
			if v == nodeId {
				return true
			}
		}
	} else if taskType == constant.TaskTypePayDB {
		for _, v := range payDBNodes {
			if v == nodeId {
				return true
			}
		}
	}
	return false
}

func GenRandomHash() (string, error) {
	hashLength := 32
	randomBytes := make([]byte, hashLength)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	hash := hex.EncodeToString(randomBytes)
	return fmt.Sprintf("0x%s", hash), nil
}

func GetDomainV2Hash(chainId int, verifyContractAddr string) (*common.Hash, error) {
	functionStr := "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
	eip712DomainHash := crypto.Keccak256Hash([]byte(functionStr))

	abiBytes32, _ := abi.NewType("bytes32", "", nil)
	abiUint256, _ := abi.NewType("uint256", "", nil)
	abiAddress, _ := abi.NewType("address", "", nil)
	args := abi.Arguments{
		{Type: abiBytes32},
		{Type: abiBytes32},
		{Type: abiBytes32},
		{Type: abiUint256},
		{Type: abiAddress},
	}

	chainIdBig := big.NewInt(0).SetInt64(int64(chainId))
	packData, err := args.Pack(
		eip712DomainHash,
		crypto.Keccak256Hash([]byte(constant.VIRTUALWALLET)),
		crypto.Keccak256Hash([]byte("2")),
		chainIdBig,
		common.HexToAddress(verifyContractAddr))
	if err != nil {
		return nil, errors.New("pack data err:" + err.Error())
	}
	domainV2Hash := crypto.Keccak256Hash(packData)
	return &domainV2Hash, nil
}

func GetSignHash(domainV2Hash, signRootHash *common.Hash) (*common.Hash, error) {
	abiBytes32, _ := abi.NewType("bytes32", "", nil)
	args := abi.Arguments{
		{Type: abiBytes32},
		{Type: abiBytes32},
	}

	packedData, err := args.Pack(domainV2Hash, signRootHash)
	if err != nil {
		return nil, err
	}
	packedData = encodePacked([]byte(constant.X1901), packedData)
	signHash := crypto.Keccak256Hash(packedData)

	return &signHash, nil
}

func encodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}

func GetChainIdAndExpTime(code string) (srcChainId, dstChainId, expTime int64) {
	x := int64(1<<32 - 1)
	if n, ok := new(big.Int).SetString(code, 10); ok {
		return new(big.Int).And(new(big.Int).Rsh(n, 64), big.NewInt(x)).Int64(),
			new(big.Int).And(new(big.Int).Rsh(n, 32), big.NewInt(x)).Int64(),
			new(big.Int).And(new(big.Int).Rsh(n, 96), big.NewInt(x)).Int64()
	}
	return 0, 0, 0
}

func GetExecuteParamHash(vw model.Vw) (*common.Hash, error) {
	approveFunc := "approve(" +
		"uint256 code," +
		"bytes32 data," +
		"address service," +
		"address gasToken," +
		"uint256 gasTokenPrice," +
		"uint256 priorityFee," +
		"uint256 gasLimit," +
		"bool isGateway)"
	approveHash := crypto.Keccak256Hash([]byte(approveFunc))
	abiBytes32, _ := abi.NewType("bytes32", "", nil)
	abiUint256, _ := abi.NewType("uint256", "", nil)
	abiAddress, _ := abi.NewType("address", "", nil)
	abiBool, _ := abi.NewType("bool", "", nil)
	args := abi.Arguments{
		{Type: abiBytes32},
		{Type: abiUint256},
		{Type: abiBytes32},
		{Type: abiAddress},
		{Type: abiAddress},
		{Type: abiUint256},
		{Type: abiUint256},
		{Type: abiUint256},
		{Type: abiBool},
	}

	keccakDataHash := crypto.Keccak256Hash(common.FromHex(vw.Data))

	codeBig, ok := big.NewInt(0).SetString(vw.Code, 10)
	if !ok {
		return nil, errors.New("code error")
	}

	gasTokenPriceBig, ok := big.NewInt(0).SetString(vw.GasTokenPrice, 10)
	if !ok {
		return nil, errors.New("gasTokenPrice error")
	}

	gasLimitBig, ok := big.NewInt(0).SetString(vw.GasLimit, 10)
	if !ok {
		return nil, errors.New("gasLimit error")
	}

	priorityFeeBig, ok := big.NewInt(0).SetString(vw.PriorityFee, 10)
	if !ok {
		return nil, errors.New("priorityFee error")
	}

	packedData, err := args.Pack(
		approveHash,
		codeBig,
		keccakDataHash,
		common.HexToAddress(vw.Service),
		common.HexToAddress(vw.GasToken),
		gasTokenPriceBig,
		priorityFeeBig,
		gasLimitBig,
		vw.IsGateWay == 1)
	if err != nil {
		return nil, errors.New("pack data err:" + err.Error())
	}

	paramHash := crypto.Keccak256Hash(packedData)
	return &paramHash, nil

}
