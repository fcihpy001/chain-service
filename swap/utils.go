// @Author fcihpy
// @Date 2024/7/29 18:40:00
// @Desc
//

package swap

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/DappOSDao/node-service/model"
	"github.com/DappOSDao/node-service/network"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fcihpy001/chain-service/constant"
	"github.com/fcihpy001/chain-service/contract"
	"github.com/fcihpy001/chain-service/logger"
	serviceModel "github.com/fcihpy001/chain-service/model"
	"github.com/fcihpy001/chain-service/token"
	"github.com/fcihpy001/chain-service/util"

	"github.com/fcihpy001/chain-service/service"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"golang.org/x/crypto/sha3"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var rate float64 = 1.1

func getTransferCallData() ([]byte, error) {

	contractAbi, err := abi.JSON(strings.NewReader(contract.GetERC20ABI()))
	if err != nil {
		logger.ErrorF("AbiJson err:%v", err)
		return nil, err
	}

	callData, err := contractAbi.Pack(
		"transfer",
		common.HexToAddress(service.GetAccount().Address),
		big.NewInt(1),
	)
	if err != nil {
		logger.ErrorF("Pack err:%v", err)
		return nil, err
	}
	return callData, nil
}

func GetCreateSrcOrderCallData(contractParam serviceModel.ContractParam) ([]byte, error) {

	contractAbi, err := abi.JSON(strings.NewReader(contract.GetPayDbABI()))
	if err != nil {
		logger.ErrorF("AbiJson err:%v", err)
		return nil, err
	}

	callData, err := contractAbi.Pack(
		"createSrcOrderETH",
		contractParam.Owner,
		contractParam.Wallet,
		contractParam.Receiver,
		contractParam.CreatePays,
		contractParam.Vw,
		contractParam.CallData,
	)
	if err != nil {
		logger.ErrorF("Pack err:%v", err)
		return nil, err
	}
	return callData, nil
}

func GetIntentDBCreateSrcOrderCallData(contractParam serviceModel.ContractParam) ([]byte, error) {

	contractAbi, err := abi.JSON(strings.NewReader(contract.GetIntentDBABI()))
	if err != nil {
		logger.ErrorF("ABIJson err:%v", err)
		return nil, err
	}

	callData, err := contractAbi.Pack(
		"createSrcIntentOrderETH",
		contractParam.Owner,
		contractParam.Receiver,
		contractParam.IntentCreatePays,
		contractParam.IntentCallParam,
	)
	if err != nil {
		logger.ErrorF("Pack err:%v", err)
		return nil, err
	}
	return callData, nil
}

func genOrderHash(
	sender, vwSign, code string,
	createPays []*model.HybridPayment,
	executePays []*model.HybridPayment,
) (string, error) {

	addressBytes, err := hex.DecodeString(strings.TrimPrefix(sender, "0x"))
	if err != nil {
		return "", err
	}
	if len(addressBytes) != 20 {
		return "", errors.New("invalid address")
	}
	rawData := append(addressBytes, []byte(vwSign)...)
	rawData = append(rawData, []byte(code)...)

	if len(createPays) > 0 {
		for _, pay := range createPays {
			rawData = append(rawData, []byte(pay.AmountIn)...)
			rawData = append(rawData, []byte(pay.AmountOut)...)
			rawData = append(rawData, []byte(pay.Code)...)
			rawData = append(rawData, []byte(strconv.Itoa(pay.DstChain))...)
			rawData = append(rawData, []byte(pay.DstHash)...)
			rawData = append(rawData, []byte(pay.Receiver)...)
			rawData = append(rawData, []byte(pay.Sender)...)

			rawData = append(rawData, []byte(strconv.Itoa(pay.SrcChain))...)
			rawData = append(rawData, []byte(pay.SrcHash)...)
			rawData = append(rawData, []byte(pay.TokenIn)...)
			rawData = append(rawData, []byte(pay.TokenOut)...)
			rawData = append(rawData, []byte(pay.PayOrderId)...)
			rawData = append(rawData, []byte(pay.VaultAddress)...)

		}
	} else if len(executePays) > 0 {
		for _, pay := range executePays {
			rawData = append(rawData, []byte(pay.AmountIn)...)
			rawData = append(rawData, []byte(pay.AmountOut)...)
			rawData = append(rawData, []byte(pay.Code)...)
			rawData = append(rawData, []byte(strconv.Itoa(pay.DstChain))...)
			rawData = append(rawData, []byte(pay.DstHash)...)
			rawData = append(rawData, []byte(pay.Receiver)...)
			rawData = append(rawData, []byte(pay.Sender)...)
			rawData = append(rawData, []byte(strconv.Itoa(pay.SrcChain))...)
			rawData = append(rawData, []byte(pay.SrcHash)...)
			rawData = append(rawData, []byte(pay.TokenIn)...)
			rawData = append(rawData, []byte(pay.TokenOut)...)
			rawData = append(rawData, []byte(pay.PayOrderId)...)
			rawData = append(rawData, []byte(pay.VaultAddress)...)
		}
	}

	hashBytes := make([]byte, 32)
	keccak256 := sha3.NewLegacyKeccak256()
	keccak256.Write(rawData)
	keccak256.Sum(hashBytes[:0])
	orderHash := fmt.Sprintf("0x%s", hex.EncodeToString(hashBytes))

	return orderHash, nil
}

func genWorkFlowHash(
	wallet, code, service string,
	data [32]byte,
	isLast bool,
) string {

	// define type
	addressType, _ := abi.NewType("address", "", []abi.ArgumentMarshaling{})
	bytesType, _ := abi.NewType("bytes32", "", nil)
	uin256Type, _ := abi.NewType("uint256", "", []abi.ArgumentMarshaling{})

	// define arg
	args := abi.Arguments{
		{Type: addressType},
		{Type: uin256Type},
		{Type: addressType},
		{Type: bytesType},
	}

	// cover code to big.int
	bigInt := new(big.Int)
	bigInt.SetString(code, 10)

	// define service addr
	serviceAddr := common.HexToAddress("0x")
	if isLast {
		serviceAddr = common.HexToAddress(service)
	}

	// pack data
	packed, err := args.Pack(
		common.HexToAddress(wallet),
		bigInt,
		serviceAddr,
		data,
	)
	if err != nil {
		fmt.Println(err)
	}

	// byte to hash
	workflowHash := crypto.Keccak256Hash(packed)
	return workflowHash.Hex()
}

func taskMsg(task serviceModel.Task) string {
	taskDesc := "<payDBTask>"
	if task.Type == constant.TaskTypeIntentDB {
		taskDesc = "<intentDBTask>"
	}
	var transferMsg string
	for i, taskIn := range task.TokenIns {
		transferMsg += fmt.Sprintf("\n(%d/%d): "+
			"srChainId(%d)->dstChainId(%d),"+
			"amountIn:(%s)->amountOut(%s),amount(%f)",
			i+1,
			len(task.TokenIns),
			task.SrcChainId,
			task.DstChainId,
			taskIn,
			task.TokenOuts[i],
			task.Amounts[i])
	}
	msg := fmt.Sprintf("node_%d execute %s: %s",
		task.NodeId,
		taskDesc,
		transferMsg)
	return msg
}

func calculateAmountInOut(
	task serviceModel.Task,
	amount float64,
	tokenInName string,
	tokenOutName string,
) (*big.Int, *big.Int, error) {

	// token info
	tokenIn := network.GetTokenInfo(task.NodeId, task.SrcChainId, tokenInName)
	tokenOut := network.GetTokenInfo(task.NodeId, task.DstChainId, tokenOutName)

	// amountIn
	amountBig := decimal.NewFromFloat(amount)
	amountBig = amountBig.RoundFloor(min(int32(tokenIn.TokenDecimal), int32(tokenOut.TokenDecimal)))
	amountInDecimal := math.Pow(10, float64(tokenIn.TokenDecimal))
	amountIn := amountBig.Mul(decimal.NewFromFloat(amountInDecimal))

	// intent logic
	if task.Type == constant.TaskTypeIntentDB {
		return amountIn.BigInt(), big.NewInt(0), nil
	}

	// same token
	if tokenInName == tokenOutName {
		tokenOutDecimal := math.Pow10(tokenOut.TokenDecimal)
		//amountOutDecimal := decimal.NewFromFloat(amount)
		amountOut := amountBig.Mul(decimal.NewFromFloat(tokenOutDecimal))
		logger.Debug("same token amount info: ",
			zap.Any("amountIn", amountIn.BigInt()),
			zap.Any("amountOut", amountOut.BigInt()),
			zap.Any("tokenInName", tokenInName),
			zap.Any("tokenOutName", tokenOutName),
		)
		return amountIn.BigInt(), amountOut.BigInt(), nil
		//return amountIn.BigInt(), big.NewInt(0).Sub(amountOut.BigInt(), big.NewInt(300000)), nil
	}

	// price
	tokenInPrice := token.GetTokenPrice(tokenIn.TokenName)
	tokenOutPrice := token.GetTokenPrice(tokenOut.TokenName)
	if tokenInPrice == 0 || tokenOutPrice == 0 {
		return nil, nil, errors.New("token in/out insufficient")
	}

	amountOutDecimal := math.Pow10(tokenOut.TokenDecimal)
	amountInUSD := decimal.NewFromFloat(amount * tokenInPrice)

	// bridgeFee
	bridgeFee := token.CalculateBridgeFee(task.NodeId, task.DstChainId)
	if bridgeFee.String() == "0" {

	}

	amountOut := amountInUSD.Sub(bridgeFee)
	amountOut = amountOut.Div(decimal.NewFromFloat(tokenOutPrice)).Mul(decimal.NewFromFloat(amountOutDecimal))

	// ie token
	if strings.ToLower(tokenIn.TokenName) == "ie" {
		amountOut = amountIn.Mul(decimal.NewFromFloat(rate))
	} else if strings.ToLower(tokenOut.TokenName) == "ie" {
		amountOut = amountIn.Div(decimal.NewFromFloat(rate))
	}
	logger.Debug("amount Info:", zap.Any("amountIn", amountIn), zap.Any("amountOut", amountOut))
	// return amountIn/out
	logger.Debug("calculateAmountInOut:",
		zap.Any("amountIn", amountIn.BigInt()),
		zap.Any("amountOut", amountOut.BigInt()),
		zap.Any("tokenInName", tokenInName),
		zap.Any("tokenOutName", tokenOutName),
	)
	return amountIn.BigInt(), amountOut.BigInt(), nil
}

func getPayOrderId(
	srcChainId,
	dstChainId int,
	taskType constant.TaskType,
) *big.Int {

	// chainId
	srcChainIdInt64 := int64(srcChainId)
	dstChainIdInt64 := int64(dstChainId)

	// nonce
	nonceStr := crypto.Keccak256Hash([]byte(time.Now().String())).Hex()
	nonce, _ := big.NewInt(0).SetString(nonceStr[2:], 16)
	number128 := big.NewInt(0).Sub(
		big.NewInt(0).Lsh(big.NewInt(1), 128),
		big.NewInt(1),
	)
	nonce = big.NewInt(0).And(nonce, number128)

	// expiration
	timeUint64 := time.Now().Add(time.Hour * 24).Unix()
	expireTime := big.NewInt(0).SetInt64(timeUint64)
	expireTime = big.NewInt(0).Lsh(expireTime, 96)

	// srcChainId
	srcId := big.NewInt(0).SetInt64(srcChainIdInt64)
	srcId = big.NewInt(0).Lsh(srcId, 64)

	// DstChainId
	dstId := big.NewInt(0).SetInt64(dstChainIdInt64)
	dstId = big.NewInt(0).Lsh(dstId, 32)

	// payOrderId
	payOrderId := big.NewInt(0).Lsh(nonce, 128)
	payOrderId = big.NewInt(0).Add(payOrderId, expireTime)
	payOrderId = big.NewInt(0).Add(payOrderId, srcId)
	payOrderId = big.NewInt(0).Add(payOrderId, dstId)
	payOrderId = big.NewInt(0).Add(payOrderId, big.NewInt(0))
	if taskType == constant.TaskTypeIntentDB {
		payOrderId = big.NewInt(0).Add(payOrderId, big.NewInt(0xffffffff))
	}
	return payOrderId
}

func approve(
	chainId int,
	tokenAddr string,
	spenderAddr string,
	ethClient *ethclient.Client) error {
	//return nil
	// init contract
	account := service.GetAccount()
	tokenContract, err := contract.NewToken(common.HexToAddress(tokenAddr), ethClient)
	if err != nil {
		log.Printf("NewToken err: %v", err)
		return err
	}
	// calculate balance
	balance, err := tokenContract.Allowance(
		&bind.CallOpts{},
		common.HexToAddress(account.Address),
		common.HexToAddress(spenderAddr))
	logger.Debug("approve info: ", zap.Any("allowance", balance))
	if err != nil {
		log.Printf("Allowance err: %v\n", err)
		return err
	}
	//enoughBalance, _ := big.NewInt(0).SetString(constant.F60, 16)
	if balance.Cmp(big.NewInt(10)) > 0 {
		return nil
	}

	// set approve balance
	uin256Max, _ := big.NewInt(0).SetString(constant.F64, 16)

	// get private key

	ops, err := bind.NewKeyedTransactorWithChainID(account.PrivateKey, big.NewInt(int64(chainId)))
	if err != nil {
		log.Printf("NewKeyedTransactorWithChainID err: %v", err)
		return err
	}

	// approve
	tx, err := tokenContract.Approve(ops, common.HexToAddress(spenderAddr), uin256Max)
	if err != nil {
		logger.ErrorF("Approve err: %v", err)
		return fmt.Errorf("approve err: %s", err.Error())
	}

	// wait for receive receipt of tx
	_, err = bind.WaitMined(context.Background(), ethClient, tx)
	if err != nil {
		log.Printf("tx on chain  err: %v\n", err)
		return err
	}
	return nil
}

func payDBHybridServiceData(
	createParams []contract.IPayDBCreatePayOrderParam,
	callParam contract.IPayDBCallParam,
	owner, wallet common.Address,
) (string, error) {

	// tuple
	tuple1, err := abi.NewType(
		"tuple",
		"",
		[]abi.ArgumentMarshaling{
			{Name: "owner", Type: "address"},
			{Name: "wallet", Type: "address"},
			{Name: "receiver", Type: "address"},
			{Name: "action", Type: "uint256"},
		})

	tuple2, err := abi.NewType(
		"tuple[]",
		"",
		[]abi.ArgumentMarshaling{
			{Name: "amountIn", Type: "uint128"},
			{Name: "amountOut", Type: "uint128"},
			{Name: "payOrderId", Type: "uint256"},
			{Name: "bridgeFee", Type: "uint128"},
			{Name: "tokenIn", Type: "address"},
			{Name: "tokenOut", Type: "address"},
			{Name: "node", Type: "address"},
		})

	tuple3, err := abi.NewType(
		"tuple",
		"",
		[]abi.ArgumentMarshaling{
			{Name: "code", Type: "uint256"},
			{Name: "service", Type: "address"},
			{Name: "dataHashZero", Type: "bytes32"},
		})

	tuple4, err := abi.NewType(
		"tuple",
		"",
		[]abi.ArgumentMarshaling{
			{Name: "node", Type: "address"},
			{Name: "nodeCallData", Type: "bytes"},
		})

	// tupleData
	tupleData1 := struct {
		Owner    common.Address
		Wallet   common.Address
		Receiver common.Address
		Action   *big.Int
	}{
		Owner:    owner,
		Wallet:   wallet,
		Receiver: wallet,
		Action:   big.NewInt(1),
	}
	tupleData3 := struct {
		Code         *big.Int
		Service      common.Address
		DataHashZero [32]byte
	}{
		Code:         big.NewInt(0),
		Service:      common.HexToAddress("0x"),
		DataHashZero: crypto.Keccak256Hash(common.FromHex("")),
	}

	// package
	transferArgs := abi.Arguments{
		{Type: tuple1},
		{Type: tuple2},
		{Type: tuple3},
		{Type: tuple4},
	}
	transferPackage, err := transferArgs.Pack(
		tupleData1,
		createParams,
		tupleData3,
		callParam,
	)
	if err != nil {
		logger.ErrorF("e2007: payDB hybrid transferArgs.Pack err: %v", err)
		return "", err
	}
	return common.Bytes2Hex(transferPackage), nil
}

func intentHybridServiceData(
	createParams []contract.IIntentDBCreateIntentOrderParam,
	callParam contract.IIntentDBCallParam,
	owner, wallet common.Address,
) (string, error) {

	// tuple
	tuple1, err := abi.NewType(
		"tuple",
		"",
		[]abi.ArgumentMarshaling{
			{Name: "owner", Type: "address"},
			{Name: "receiver", Type: "address"},
			{Name: "action", Type: "uint256"},
		})

	tuple2, err := abi.NewType(
		"tuple[]",
		"",
		[]abi.ArgumentMarshaling{
			{Name: "srcIntentToken", Type: "address"},
			{Name: "underlyingAssets", Type: "address"},
			{Name: "underlyingAssetsAmount", Type: "uint256"},
			{Name: "payOrderId", Type: "uint256"},
			{Name: "bridgeFee", Type: "uint256"},
			{Name: "node", Type: "address"},
		})

	tuple3, err := abi.NewType(
		"tuple",
		"",
		[]abi.ArgumentMarshaling{
			{Name: "node", Type: "address"},
			{Name: "nodeCallData", Type: "bytes"},
		})

	// tupleData
	tupleData1 := struct {
		Owner    common.Address
		Receiver common.Address
		Action   *big.Int
	}{
		Owner:    owner,
		Receiver: wallet,
		Action:   big.NewInt(1),
	}

	// package
	transferArgs := abi.Arguments{
		{Type: tuple1},
		{Type: tuple2},
		{Type: tuple3},
	}
	transferPackage, err := transferArgs.Pack(
		tupleData1,
		createParams,
		callParam,
	)
	if err != nil {
		logger.ErrorF("e2006: transferArgs.Pack err: %v", err)
		return "", err
	}
	return common.Bytes2Hex(transferPackage), nil
}

func transferServiceData(token string, to string, amount *big.Int) (string, error) {
	tupleStruct, _ := abi.NewType("tuple", "struct Overloader.F", []abi.ArgumentMarshaling{
		{Name: "Token", Type: "address"},
		{Name: "To", Type: "address"},
		{Name: "Amount", Type: "uint256"},
	})

	packStr := packData(constant.ZeroAddr, 0, 1)

	tupleData := struct {
		Token  common.Address
		To     common.Address
		Amount *big.Int
	}{
		Token:  common.HexToAddress(token),
		To:     common.HexToAddress(to),
		Amount: amount,
	}

	transferArgs := abi.Arguments{{Type: tupleStruct}}

	transferPack, err := transferArgs.Pack(tupleData)
	if err != nil {
		logger.ErrorF("e2005: transferArgs.Pack err: %v", err)
	}

	return fmt.Sprintf("%s%s%s", packStr, constant.MoreZero, common.Bytes2Hex(transferPack)), nil
}

func transferServiceSign(
	vw model.Vw,
	chainId int,
) (string, error) {

	// paramHash
	executeParamHash, err := util.GetExecuteParamHash(vw)
	if err != nil {
		return "", err
	}

	// domainHash
	domainHash, err := util.GetDomainV2Hash(
		chainId,
		network.GetChainInfo(chainId).Contract.Manager)

	// signHash
	signHash, err := util.GetSignHash(domainHash, executeParamHash)
	if err != nil {
		return "", err
	}

	// sign
	signature, err := crypto.Sign(signHash[:], service.GetAccount().PrivateKey)
	if err != nil {
		return "", err
	}
	signature[64] += 27
	return hexutil.Encode(signature), nil
}

func packData(gasToken string, remainForGas int, action uint8) string {
	packed := big.NewInt(0)
	packed.SetBytes(common.HexToAddress(gasToken).Bytes())

	// remainForGas
	packed.Lsh(packed, 94)
	packed.Or(packed, big.NewInt(int64(remainForGas)))

	// action
	packed.Lsh(packed, 2)
	packed.Or(packed, big.NewInt(int64(action)))

	return common.BytesToHash(packed.Bytes()).Hex()
}

func getChainPriorityFee(nodeId, chainId int) string {
	vaultInfos := network.GetNodeInfo(nodeId).NodeVaultInfos
	for _, vaultInfo := range vaultInfos {
		for _, nodeFee := range vaultInfo.NodeFees {
			if nodeFee.ChainId == chainId && nodeFee.TokenInfo.TokenName == "USDT" {
				return nodeFee.PriorityFee
			}
		}
	}
	return "1000000"
}

func multiTransferServiceData(executePays []*model.HybridPayment) (string, error) {
	tupleType, _ := abi.NewType("tuple[]", "", []abi.ArgumentMarshaling{
		{Name: "Token", Type: "address"},
		{Name: "To", Type: "address"},
		{Name: "Amount", Type: "uint256"},
	})

	var transferDatas []TransferData
	for _, payment := range executePays {
		amountBig, _ := new(big.Int).SetString(payment.AmountOut, 10)
		data := TransferData{
			Token:  common.HexToAddress(payment.TokenOut),
			To:     common.HexToAddress(payment.Owner),
			Amount: amountBig,
		}
		transferDatas = append(transferDatas, data)
	}
	transferArgs := abi.Arguments{
		{Type: tupleType},
	}
	transferPack, err := transferArgs.Pack(transferDatas)
	if err != nil {
		logger.ErrorF("e2008: multi transferArgs.Pack err: %v", err)
		return "", nil
	}
	return fmt.Sprintf("%s%s", constant.MoreZero, common.Bytes2Hex(transferPack)), nil
}

type TransferData struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
}

func checkNodeStatus(tasks []serviceModel.Task) error {

	for _, task := range tasks {
		if len(task.TokenIns) != len(task.Amounts) || len(task.Amounts) != len(task.TokenOuts) {
			return errors.New("task info slice length mismatch")
		}
		//if !service.IsSupportNodeId(task.NodeId) {
		//	return errors.New("node id is invalid")
		//}
		//
		//if !service.IsSupportChainId(task.SrcChainId) {
		//	return errors.New("task chainId is not in support list of node")
		//}
		if (task.DstChainId == 0 || task.DstChainId == 1) && task.Type != constant.TaskTypeIntentDB {
			return errors.New("task type mismatch")
		}
		if task.Type == constant.TaskTypeIntentDB && (task.NodeId != 21 && task.NodeId != 22) {
			return errors.New("intentDB task only support in node21,node22")
		}
		for i, token := range task.TokenIns {
			tokenIn := network.GetTokenInfo(task.NodeId, task.SrcChainId, token)
			if len(strings.TrimSpace(tokenIn.TokenName)) < 1 {
				return errors.New("token of tokenIns is not support in node")
			}
			tokenOut := network.GetTokenInfo(task.NodeId, task.DstChainId, task.TokenOuts[i])
			if len(strings.TrimSpace(tokenOut.TokenName)) < 1 && task.Type != constant.TaskTypeIntentDB {
				return errors.New("token of tokenOuts is not support in node")
			}
		}
	}
	return nil
}

func BridgeAsset(task serviceModel.Task) (string, error) {
	orderHash, err := PayDBSwap([]serviceModel.Task{task}, false)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	isSuccess, err := service.RetriesCheckTxsSate(orderHash)
	if err != nil || !isSuccess {
		logger.ErrorF("previous order process failure---%d", isSuccess)
		return "", err
	}

	return orderHash, nil
}

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
