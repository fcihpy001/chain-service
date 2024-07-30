// @Author fcihpy
// @Date 2024/7/29 18:02:00
// @Desc
//

package swap

import (
	"context"
	"errors"
	"fmt"
	"github.com/DappOSDao/node-service/model"
	"github.com/DappOSDao/node-service/network"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fcihpy001/chain-service/constant"
	"github.com/fcihpy001/chain-service/contract"
	"github.com/fcihpy001/chain-service/logger"
	serviceModel "github.com/fcihpy001/chain-service/model"
	"github.com/fcihpy001/chain-service/service"
	"github.com/fcihpy001/chain-service/util"
	"go.uber.org/zap"
	"math/big"
	"sync"
)

func PayDBSwap(tasks []serviceModel.Task, force bool) (string, error) {
	if force {
		//if err := checkNodeStatus(tasks); err != nil {
		//	return "", err
		//}
	}

	var wg sync.WaitGroup
	orders := make([]model.Order, 0)
	var errorList []error
	for _, task := range tasks {
		if task.Type != constant.TaskTypePayDB {
			return "", errors.New("e2047: task type mismatch")
		}
		if !util.ContainNodeId(task.Type, task.NodeId) {
			return "", errors.New("e2048: nodeId is not support")
		}
		if task.DstChainId == 0 {
			return "", errors.New("e2049: dst chainId can't be zero")
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			tempOrder, err := getPayDBOrder(task)
			if err != nil {
				errorList = append(errorList, err)
				return
			}
			fmt.Printf("---------------->>>>>>>>node_%d:payDBTask:%d-%dexecute successfully,orderHash:%s\n",
				task.NodeId,
				task.SrcChainId,
				task.DstChainId,
				tempOrder.OrderHash,
			)
			orders = append(orders, tempOrder)
		}()
	}
	wg.Wait()
	if len(errorList) > 0 {
		return "", errorList[0]
	}

	var order model.Order
	chainTxMap := make(map[int]*model.Tx)
	for i, o := range orders {
		order.OrderHash = o.OrderHash
		order.Sender = o.Sender

		// srcChainTx
		tempTx := o.Txs[0]
		srcChainTx := model.Tx{Id: i}

		// createPay
		srcChainTx.CreatePays = append(srcChainTx.CreatePays, tempTx.CreatePays...)
		srcChainTx.Vw = tempTx.Vw
		order.Txs = append(order.Txs, &srcChainTx)

		// executePay
		index := len(tasks) + i
		dstChainId := o.Txs[1].ExecutePays[0].DstChain
		dstChainTx, ok := chainTxMap[dstChainId]
		if !ok {
			dstChainTx = &model.Tx{Id: index}
			chainTxMap[dstChainId] = dstChainTx
		}
		dstChainTx.ExecutePays = append(dstChainTx.ExecutePays, o.Txs[1].ExecutePays...)
		dstChainTx.Vw = o.Txs[1].Vw
		dstChainTx.DependenceIds = append(dstChainTx.DependenceIds, i)
	}

	for _, tx := range chainTxMap {
		// sign service data
		if tx.Vw != nil {
			transferData, _ := multiTransferServiceData(tx.ExecutePays)
			vw := tx.Vw
			vw.Data = transferData
			sign, err := transferServiceSign(*vw, tx.ExecutePays[0].DstChain)
			vw.Signature = sign
			if err != nil {
				return "", err
			}
			tx.Vw = vw
		}
		order.Txs = append(order.Txs, tx)
	}

	// support order to supperNode
	err := network.PostIntentOrder(order)
	if err != nil {
		return "", fmt.Errorf("e2011: sendOrderToSuperNode err: %s", err.Error())
	}
	logger.Info("PayDBTask(8/9): 💯 Congratulate 🎊 Send to supperNode successfully")

	// 9. verify transaction receipt
	err = service.VerifyTransaction(order)
	if err != nil {
		return "", fmt.Errorf("e2012: verifyTransaction err: %v", err)
	}
	logger.InfoF("PayDBTask(9/9): ⛳️execute successfully: %s%s", network.GetQueryOrderHashApi(), order.OrderHash)
	return order.OrderHash, nil
}

func getPayDBOrder(task serviceModel.Task) (model.Order, error) {
	logger.InfoF("PayDBTask(0/9): ⭐  %s", taskMsg(task))

	// 2. package params
	contractParam, err := getPayDBContractParam(task)
	if err != nil {
		return model.Order{}, err
	}
	logger.Info("PayDBTask(2/9): 🧮 contract param package successfully...")

	// 3. send transaction
	txHash := ""
	err = errors.New("")
	if !task.HasSrcVw {
		txHash, err = callPayDBContract(task, contractParam)
		if err != nil {
			return model.Order{}, err
		}
	}

	// 7. create order
	order, err := createPayDBOrder(task, contractParam, txHash)
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func getPayDBContractParam(task serviceModel.Task) (serviceModel.ContractParam, error) {

	// verify param length
	if len(task.TokenIns) != len(task.Amounts) {
		return serviceModel.ContractParam{}, errors.New("e2014: slice length mismatch")
	}

	// token info
	srcChainId := task.SrcChainId
	dstChainId := task.DstChainId

	// node addr
	nodeAddr := network.GetNodeVault(task.NodeId, task.SrcChainId)

	// createPayOrderParam
	var createPayParams []contract.IPayDBCreatePayOrderParam
	for i, amount := range task.Amounts {

		// 1. generate payOrderId
		payOrderId := getPayOrderId(task.SrcChainId, task.DstChainId, task.Type)
		logger.InfoF("PayDBTask(1/9): 💹 pays(%d/%d) Generate payOrderId(%d) successfully...", i+1, len(task.Amounts), payOrderId)

		tokenIn := network.GetTokenInfo(task.NodeId, srcChainId, task.TokenIns[i])
		tokenOut := network.GetTokenInfo(task.NodeId, dstChainId, task.TokenOuts[i])
		if task.Type == constant.TaskTypeIntentDB {
			tokenOutName := fmt.Sprintf("Intent %s", tokenIn.TokenName)
			tokenOutName = tokenOutName[:len(tokenOutName)-1]
			tokenOut = network.GetTokenInfo(task.NodeId, srcChainId, tokenOutName)
		}

		// amountIn/amountOut
		amountIn, amountOut, err := calculateAmountInOut(task, amount, tokenIn.TokenName, tokenOut.TokenName)
		if err != nil {
			return serviceModel.ContractParam{}, fmt.Errorf("e2015: calculateAmountInOut err: %v", err)
		}
		logger.Debug("amount info:", zap.Any("amountIn", amountIn), zap.Any("amountOut", amountOut))

		// cretePayParam
		createPayParam := contract.IPayDBCreatePayOrderParam{
			AmountIn:   amountIn,
			AmountOut:  amountOut,
			PayOrderId: payOrderId,
			BridgeFee:  big.NewInt(0),
			TokenIn:    common.HexToAddress(tokenIn.TokenAddress),
			TokenOut:   common.HexToAddress(tokenOut.TokenAddress),
			Node:       common.HexToAddress(nodeAddr),
		}
		createPayParams = append(createPayParams, createPayParam)
	}

	// vwParam
	vwParam := contract.IPayDBVwOrderDetail{
		Code:     big.NewInt(0),
		Service:  common.HexToAddress("0x"),
		DataHash: crypto.Keccak256Hash(common.FromHex("")),
	}

	// callDataParam
	callParam := contract.IPayDBCallParam{
		Node:         common.HexToAddress(nodeAddr),
		NodeCallData: common.FromHex(network.GetNodeCallData(task.NodeId, task.SrcChainId)),
	}

	return serviceModel.ContractParam{
		CreatePays: createPayParams,
		Vw:         vwParam,
		CallData:   callParam,
	}, nil
}

func createPayDBOrder(
	task serviceModel.Task,
	contractParam serviceModel.ContractParam,
	txHash string,
) (model.Order, error) {

	// contract param
	vwParam := contractParam.Vw

	// account info
	account := service.GetAccount()
	sender := account.Address
	receiver := account.Address
	wallet := account.Address
	owner := account.Address
	if task.HasDstVw {
		receiver = service.GetVW(task.DstChainId)
	}
	if task.HasSrcVw {
		sender = service.GetVW(task.SrcChainId)
	}
	if task.HasDstVw && task.HasSrcVw {
		wallet = service.GetVW(task.DstChainId)
	}
	vwManager := network.GetChainInfo(task.DstChainId).Contract.Manager
	workflowHash := genWorkFlowHash(
		wallet,
		vwParam.Code.String(),
		vwParam.Service.String(),
		vwParam.DataHash,
		true)

	srcChainTx := &model.Tx{Id: 0}
	dstChainTx := &model.Tx{Id: len(contractParam.CreatePays) - 1, DependenceIds: []int{0}}

	for _, createParam := range contractParam.CreatePays {
		// createPay and executePay construct
		payment := &model.HybridPayment{
			VaultAddress: createParam.Node.String(),
			PayOrderId:   createParam.PayOrderId.String(),
			AmountIn:     createParam.AmountIn.String(),
			AmountOut:    createParam.AmountOut.String(),
			TokenIn:      createParam.TokenIn.String(),
			TokenOut:     createParam.TokenOut.String(),
			Code:         vwParam.Code.String(),
			WorkFlowHash: workflowHash,
			Sender:       sender,
			Receiver:     receiver,
			Owner:        owner,
			Manager:      vwManager,
			SrcHash:      txHash,
			SrcChain:     task.SrcChainId,
			DstChain:     task.DstChainId,
		}

		// construct createPayTx and executePayTx
		srcChainTx.CreatePays = append(srcChainTx.CreatePays, payment)
		if task.HasSrcVw {
			vw, err := createPayDBSrcVw(task, contractParam)
			if err != nil {
				return model.Order{}, fmt.Errorf("e2031: createSrcVw err: %v", err)
			}
			srcChainTx.Vw = &vw
		}

		if task.HasDstVw {
			vw, err := createPayDBDstVw(task, contractParam)
			if err != nil {
				return model.Order{}, fmt.Errorf("e2032: createDstVw err: %v", err)
			}
			dstChainTx.Vw = &vw
		}
		dstChainTx.ExecutePays = append(dstChainTx.ExecutePays, payment)
	}

	// construct order data
	order := model.Order{}
	orderHash, err := util.GenRandomHash()
	if err != nil {
		return model.Order{}, fmt.Errorf("e2016: order hash generate fail:%s", err.Error())
	}
	if !task.HasSrcVw {
		orderHash, err = genOrderHash(
			sender,
			"",
			vwParam.Code.String(),
			srcChainTx.CreatePays,
			dstChainTx.ExecutePays)
		if err != nil {
			return model.Order{}, fmt.Errorf("e2017: order hash generate fail:%s", err.Error())
		}
	}
	order.Sender = account.Address
	order.Txs = append(order.Txs, srcChainTx)
	order.Txs = append(order.Txs, dstChainTx)
	order.OrderHash = orderHash

	return order, nil
}

func callPayDBContract(
	task serviceModel.Task,
	contractParam serviceModel.ContractParam,
) (string, error) {
	var err error
	// eth client
	ethClient, err := service.GetEvmClient(service.GetRPC(task.SrcChainId), task.SrcChainId)
	if err != nil {
		return "", err
	}
	logger.Info("PayDBTask(3/9): 🔐 get client successfully...")

	// get payDB contract addr
	dbAddr := network.GetChainInfo(task.SrcChainId).Contract.PayDB

	// account info
	account := service.GetAccount()
	sender := account.Address
	owner := sender
	wallet := account.Address
	receiver := account.Address
	if task.HasDstVw {
		receiver = service.GetVW(task.DstChainId)
	}

	// 4. approve
	opts, err := bind.NewKeyedTransactorWithChainID(account.PrivateKey, big.NewInt(int64(task.SrcChainId)))
	if err != nil {
		return "", err
	}

	optValue := big.NewInt(0)
	for i, name := range task.TokenIns {

		// token info
		tokenIn := network.GetTokenInfo(task.NodeId, task.SrcChainId, name)
		tokenOut := network.GetTokenInfo(task.NodeId, task.DstChainId, task.TokenOuts[i])

		// amountIn/out
		amountIn, _, err := calculateAmountInOut(task, task.Amounts[i], tokenIn.TokenName, tokenOut.TokenName)
		if err != nil {
			return "", err
		}

		logger.InfoF("========,chainId:%d tokenName:%v-tokenAddr:%v-payDB:%v", task.SrcChainId, tokenIn.TokenName, tokenIn.TokenAddress, dbAddr)

		if tokenIn.TokenAddress == constant.ZeroAddr {
			optValue.Add(optValue, amountIn)

		} else {
			err = approve(task.SrcChainId, tokenIn.TokenAddress, dbAddr, ethClient)
			if err != nil {
				return "", err
			}
			logger.InfoF("PayDBTask(4/9): 🔋  Wallet(%s) approve transfer successfully...\n", account.Address)
		}
	}

	// 5. call payDB contract
	opts.Value = optValue
	var tx *types.Transaction

	// set contract data
	contractParam.Wallet = common.HexToAddress(wallet)
	contractParam.Receiver = common.HexToAddress(receiver)
	contractParam.Owner = common.HexToAddress(owner)
	contractParam.PayDB = common.HexToAddress(dbAddr)
	contractParam.Amount = optValue
	contractParam.ChainId = task.SrcChainId

	// call contract
	tx, err = service.SendTransaction(contractParam)
	if err != nil {
		logger.ErrorF("e2019: createSrcOrderETH err: %v", err.Error())
		return "", fmt.Errorf("e2019: createSrcOrderETH err: %v", err)
	}
	logger.Info("PayDBTask(5/9): 🏆️ Call function createSrcOrder of payDBContract successfully and wait for receive txHash...")

	// 6. wait for tx receipt
	_, _ = bind.WaitMined(context.Background(), ethClient, tx)
	logger.InfoF("PayDBTask(6/9): ⛳️ Transaction(%s) on chain  successfully...\n", tx.Hash().String())

	return tx.Hash().String(), nil
}

func createPayDBSrcVw(
	task serviceModel.Task,
	contractParam serviceModel.ContractParam,
) (model.Vw, error) {

	// address info
	account := service.GetAccount()
	wallet := service.GetVW(task.SrcChainId)
	vwManager := network.GetChainInfo(task.SrcChainId).Contract.Manager
	receiver := service.GetAccount().Address
	if task.HasDstVw {
		receiver = service.GetVW(task.DstChainId)
	}
	// service data
	data, _ := payDBHybridServiceData(
		contractParam.CreatePays,
		contractParam.CallData,
		common.HexToAddress(account.Address),
		common.HexToAddress(receiver),
	)

	// service code
	code := getPayOrderId(task.SrcChainId, task.SrcChainId, task.Type)

	// vw data
	srcVw := model.Vw{
		Code:          code.String(),
		Wallet:        wallet,
		Owner:         account.Address,
		Sender:        wallet,
		Service:       network.GetChainInfo(task.SrcChainId).Contract.HybridService,
		Data:          data,
		Hash:          "",
		WorkFlowHash:  "",
		ChainId:       task.SrcChainId,
		IsGateWay:     0,
		GasToken:      network.GetTokenInfo(task.NodeId, task.SrcChainId, "USDT").TokenAddress,
		GasTokenPrice: "55075000",
		GasLimit:      "85075000",
		PriorityFee:   getChainPriorityFee(task.NodeId, task.SrcChainId),
		Manager:       vwManager,
		Signature:     "",
		App:           "Wallet",
		Text:          "Withdraw",
	}

	// sign data
	srcSignature, err := transferServiceSign(srcVw, task.SrcChainId)
	if err != nil {
		return model.Vw{}, err
	}
	srcVw.Signature = srcSignature
	return srcVw, nil
}

func createPayDBDstVw(
	task serviceModel.Task,
	contractParam serviceModel.ContractParam,
) (model.Vw, error) {

	// address info
	account := service.GetAccount()

	wallet := service.GetVW(task.DstChainId)
	vwManager := network.GetChainInfo(task.DstChainId).Contract.Manager

	// service data
	createPay := contractParam.CreatePays[0]
	tokenOut := network.GetTokenInfo(task.NodeId, task.DstChainId, createPay.TokenOut.String())
	data, err := transferServiceData(tokenOut.TokenAddress, service.GetAccount().Address, createPay.AmountOut)
	if err != nil {
		return model.Vw{}, err
	}

	// service code
	code := getPayOrderId(task.DstChainId, task.DstChainId, task.Type)

	// vw data
	vw := model.Vw{
		Code:          code.String(),
		Wallet:        wallet,
		Owner:         account.Address,
		Sender:        wallet,
		Service:       network.GetChainInfo(task.DstChainId).Contract.TokenTransferService,
		Data:          data,
		Hash:          "",
		WorkFlowHash:  "",
		ChainId:       task.DstChainId,
		IsGateWay:     0,
		GasToken:      network.GetTokenInfo(task.NodeId, task.DstChainId, "USDT").TokenAddress,
		GasTokenPrice: "35075000",
		GasLimit:      "45075000",
		PriorityFee:   getChainPriorityFee(task.NodeId, task.DstChainId),
		Manager:       vwManager,
		Signature:     "",
		App:           "Wallet",
		Text:          "Withdraw",
	}

	// sign data
	signature, err := transferServiceSign(vw, task.DstChainId)
	if err != nil {
		return model.Vw{}, err
	}
	vw.Signature = signature
	return vw, nil
}
