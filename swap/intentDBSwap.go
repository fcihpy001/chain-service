// @Author fcihpy
// @Date 2024/7/29 18:02:00
// @Desc
//

package swap

//
//import (
//	"context"
//	"fmt"
//	nodeModel "github.com/DappOSDao/node-service/model"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/crypto"
//	"github.com/fcihpy001/chain-service/contract"
//	"github.com/fcihpy001/chain-service/model"
//	"github.com/fcihpy001/chain-service/service"
//	"math/big"
//)
//
//func IntentDBSwap(tasks []model.Task, force bool) (string, error) {
//	if force {
//		if err := checkNodeStatus(tasks); err != nil {
//			return "", err
//		}
//	}
//
//	var order nodeModel.Order
//	chainTxMap := make(map[int]*nodeModel.Tx)
//	for i, task := range tasks {
//		if task.Type != constant.TaskTypeIntentDB {
//			return "", errors.New("e2045: task type mismatch")
//		}
//		if !ContainNodeId(task.Type, task.NodeId) {
//			return "", errors.New("e2046: nodeId is not support")
//		}
//		if task.DstChainId != 0 {
//			return "", errors.New("e2049: dst chainId must be zero")
//		}
//
//		tempOrder, err := getIntentDBOrder(task)
//		if err != nil {
//			return "", err
//		}
//		order.OrderHash = tempOrder.OrderHash
//		order.Sender = tempOrder.Sender
//
//		// srcChainTx
//		tempTx := tempOrder.Txs[0]
//		srcChainTx := model.Tx{Id: i}
//
//		// createPay
//		srcChainTx.CreatePays = append(srcChainTx.CreatePays, tempTx.CreatePays...)
//		srcChainTx.Vw = tempTx.Vw
//		order.Txs = append(order.Txs, &srcChainTx)
//
//		// executePay
//		index := len(tasks) + i
//		dstChainTx, ok := chainTxMap[task.DstChainId]
//		if !ok {
//			dstChainTx = &model.Tx{Id: index}
//			chainTxMap[task.DstChainId] = dstChainTx
//		}
//		for _, pay := range tempOrder.Txs[1].ExecutePays {
//			if pay.TokenIn == constant.ZeroAddr {
//				ethDstChainTx, eok := chainTxMap[1]
//				if !eok {
//					ethDstChainTx = &model.Tx{Id: len(tasks) + 1}
//					chainTxMap[1] = ethDstChainTx
//				}
//				chainTxMap[1] = ethDstChainTx
//				ethDstChainTx.ExecutePays = append(ethDstChainTx.ExecutePays, pay)
//				ethDstChainTx.DependenceIds = append(ethDstChainTx.DependenceIds, i)
//			} else {
//				dstChainTx.ExecutePays = append(dstChainTx.ExecutePays, pay)
//			}
//		}
//		dstChainTx.DependenceIds = append(dstChainTx.DependenceIds, i)
//	}
//
//	for _, tx := range chainTxMap {
//		// sign service data
//		if tx.Vw != nil {
//			transferData, _ := multiTransferServiceData(tx.ExecutePays)
//			vw := tx.Vw
//			vw.Data = transferData
//			sign, err := transferServiceSign(*vw, tx.ExecutePays[0].DstChain)
//			vw.Signature = sign
//			if err != nil {
//				return "", err
//			}
//			tx.Vw = vw
//		}
//		order.Txs = append(order.Txs, tx)
//	}
//
//	// 8.send IntentTask order to superNode
//	err := service.PostIntentOrder(order)
//	if err != nil {
//		return "", err
//	}
//	logger.Info("IntentTask(8/9): 💯 Congratulate 🎊 Send to supperNode successfully")
//
//	// 9. verify transaction receipt
//	err = service.VerifyTransaction(order)
//	if err != nil {
//		return "", fmt.Errorf("e2012: verifyTransaction err: %v", err)
//	}
//	logger.InfoF("IntentTask(9/9): ⛳️execute successfully: %s%s", service.GetQueryOrderHashApi(), order.OrderHash)
//	return order.OrderHash, nil
//}
//
//func getIntentDBOrder(task model.Task) (nodeModel.Order, error) {
//	logger.InfoF("IntentTask(0/9): ⭐  %s", taskMsg(task))
//
//	// 2. package params
//	contractParam, err := getIntentDBContractParam(task)
//	if err != nil {
//		return model.Order{}, err
//	}
//	logger.Info("IntentTask(2/9): 🧮 contract param package successfully...")
//
//	// 3. send transaction
//	txHash := ""
//	err = errors.New("")
//	if !task.HasSrcVw {
//		txHash, err = callIntentDBContract(task, contractParam)
//		if err != nil {
//			return model.Order{}, err
//		}
//	}
//
//	// 7. create order
//	order, err := createIntentDBOrder(task, contractParam, txHash)
//	if err != nil {
//		return model.Order{}, err
//	}
//
//	return order, nil
//}
//
//func getIntentDBContractParam(task model.Task) (model.ContractParam, error) {
//
//	// verify param length
//	if len(task.TokenIns) != len(task.Amounts) {
//		return model.ContractParam{}, errors.New("slice length mismatch")
//	}
//
//	// token info
//	srcChainId := task.SrcChainId
//	dstChainId := task.DstChainId
//
//	// node addr
//	nodeAddr := service.GetNodeVault(task.NodeId, task.SrcChainId)
//
//	var createPayParams []contract.IIntentDBCreateIntentOrderParam
//	for i, amount := range task.Amounts {
//
//		// 1. generate payOrderId
//		payOrderId := getPayOrderId(task.SrcChainId, task.DstChainId, task.Type)
//		logger.InfoF("IntentTask(1/9): 💹 pays(%d/%d) Generate payOrderId(%d) successfully...", i+1, len(task.Amounts), payOrderId)
//
//		tokenIn := service.GetTokenInfo(task.NodeId, srcChainId, task.TokenIns[i])
//		tokenOut := service.GetTokenInfo(task.NodeId, dstChainId, task.TokenOuts[i])
//		if task.Type == constant.TaskTypeIntentDB {
//			tokenOutName := fmt.Sprintf("Intent %s", tokenIn.TokenName)
//			tokenOutName = tokenOutName[:len(tokenOutName)-1]
//			tokenOut = service.GetTokenInfo(task.NodeId, srcChainId, tokenOutName)
//		}
//
//		// amountIn/amountOut
//		amountIn, amountOut, err := calculateAmountInOut(task, amount, tokenIn.TokenName, tokenOut.TokenName)
//		if err != nil {
//			return model.ContractParam{}, err
//		}
//		logger.Debug("amount info:", zap.Any("amountIn", amountIn), zap.Any("amountOut", amountOut))
//
//		// cretePayParam
//		createPayParam := contract.IIntentDBCreateIntentOrderParam{
//			SrcIntentToken:         common.HexToAddress(tokenOut.TokenAddress),
//			UnderlyingAssetsAmount: amountIn,
//			UnderlyingAssets:       common.HexToAddress(tokenIn.TokenAddress),
//			PayOrderId:             payOrderId,
//			BridgeFee:              big.NewInt(0),
//			Node:                   common.HexToAddress(nodeAddr),
//		}
//		createPayParams = append(createPayParams, createPayParam)
//	}
//
//	// vwParam
//	vwParam := contract.IPayDBVwOrderDetail{
//		Code:     big.NewInt(0),
//		Service:  common.HexToAddress("0x"),
//		DataHash: crypto.Keccak256Hash(common.FromHex("")),
//	}
//
//	// callDataParam
//	callParam := contract.IIntentDBCallParam{
//		Node:         common.HexToAddress(nodeAddr),
//		NodeCallData: common.FromHex(service.GetNodeCallData(task.NodeId, task.SrcChainId)),
//	}
//
//	return model.ContractParam{
//		IntentCreatePays: createPayParams,
//		Vw:               vwParam,
//		IntentCallParam:  callParam,
//	}, nil
//}
//
//func createIntentDBOrder(
//	task model.Task,
//	contractParam model.ContractParam,
//	txHash string,
//) (model.Order, error) {
//
//	// contract param
//	vwParam := contractParam.Vw
//
//	// account info
//	account := service.GetAccount()
//	sender := account.Address
//	owner := account.Address
//	receiver := account.Address
//	wallet := account.Address
//	if task.HasDstVw {
//		receiver = service.GetVW(task.DstChainId)
//	}
//	if task.HasSrcVw {
//		sender = service.GetVW(task.SrcChainId)
//	}
//	if task.HasDstVw && task.HasSrcVw {
//		wallet = service.GetVW(task.DstChainId)
//	}
//
//	vwManager := service.GetChainInfo(task.SrcChainId).Contract.Manager
//	workflowHash := genWorkFlowHash(
//		wallet,
//		vwParam.Code.String(),
//		vwParam.Service.String(),
//		vwParam.DataHash,
//		true)
//
//	srcChainTx := &model.Tx{Id: 0}
//	dstChainTx := &model.Tx{Id: len(contractParam.CreatePays) - 1, DependenceIds: []int{0}}
//	for _, createParam := range contractParam.IntentCreatePays {
//		// createPay construct
//		payment := &model.HybridPayment{
//			VaultAddress: createParam.Node.String(),
//			PayOrderId:   createParam.PayOrderId.String(),
//			AmountIn:     createParam.UnderlyingAssetsAmount.String(),
//			AmountOut:    "0",
//			TokenIn:      createParam.UnderlyingAssets.String(),
//			TokenOut:     "",
//			Code:         vwParam.Code.String(),
//			WorkFlowHash: workflowHash,
//			Sender:       sender,
//			Receiver:     receiver,
//			Owner:        owner,
//			Manager:      vwManager,
//			SrcHash:      txHash,
//			SrcChain:     task.SrcChainId,
//			DstChain:     task.DstChainId,
//		}
//
//		// construct createPayTx and executePayTx
//		srcChainTx.CreatePays = append(srcChainTx.CreatePays, payment)
//		if task.HasSrcVw {
//			vw, err := createIntentDBSrcVw(task, contractParam)
//			if err != nil {
//				return model.Order{}, err
//			}
//			srcChainTx.Vw = &vw
//		}
//
//		// executePay
//		dstChainTx.ExecutePays = append(dstChainTx.ExecutePays, payment)
//	}
//
//	// construct order data
//	order := model.Order{}
//	orderHash, err := util.GenRandomHash()
//	if err != nil {
//		return order, err
//	}
//	if !task.HasSrcVw {
//		orderHash, err = genOrderHash(
//			sender,
//			"",
//			vwParam.Code.String(),
//			srcChainTx.CreatePays,
//			dstChainTx.ExecutePays)
//		if err != nil {
//			return order, err
//		}
//	}
//	order.Sender = account.Address
//	order.Txs = append(order.Txs, srcChainTx)
//	order.Txs = append(order.Txs, dstChainTx)
//	order.OrderHash = orderHash
//
//	return order, nil
//}
//
//func callIntentDBContract(
//	task model.Task,
//	contractParam model.ContractParam,
//) (string, error) {
//
//	// eth client
//	ethClient, err := service.GetEthClientByChainId(task.SrcChainId)
//	if err != nil {
//		return "", err
//	}
//	logger.Info("IntentTask(3/9): 🔐 get client successfully...")
//
//	// init intentDB contract
//	dbAddr := service.GetChainInfo(task.SrcChainId).Contract.IntentDB
//	//dbContract, err := contract.NewIntentDB(common.HexToAddress(dbAddr), ethClient)
//	//if err != nil {
//	//	return "", err
//	//}
//	//logger.Info("IntentTask(3/9): 🔐 Init payDBContract successfully...")
//
//	// account info
//	account := service.GetAccount()
//	sender := account.Address
//	owner := sender
//	wallet := account.Address
//	receiver := account.Address
//	if task.HasDstVw {
//		receiver = service.GetVW(task.DstChainId)
//	}
//
//	// 4. approve
//	opts, err := bind.NewKeyedTransactorWithChainID(account.PrivateKey, big.NewInt(int64(task.SrcChainId)))
//	if err != nil {
//		return "", err
//	}
//
//	optValue := big.NewInt(0)
//	for i, name := range task.TokenIns {
//		// token info
//		tokenIn := service.GetTokenInfo(task.NodeId, task.SrcChainId, name)
//		tokenOut := service.GetTokenInfo(task.NodeId, task.DstChainId, task.TokenOuts[i])
//
//		// amountIn/out
//		amountIn, _, err := calculateAmountInOut(task, task.Amounts[i], tokenIn.TokenName, tokenOut.TokenName)
//		if err != nil {
//			return "", err
//		}
//
//		if tokenIn.TokenAddress == constant.ZeroAddr {
//			optValue = optValue.Add(optValue, amountIn)
//
//		} else {
//			err = approve(task.SrcChainId, tokenIn.TokenAddress, dbAddr, ethClient)
//			if err != nil {
//				return "", err
//			}
//			logger.InfoF("IntentTask(4/9): 🔋  Wallet(%s) approve transfer successfully...\n", account.Address)
//		}
//	}
//
//	logger.Debug("optInfo", zap.Any("opt", opts.Value), zap.Any("value", optValue))
//
//	// 5. call DB contract
//	opts.Value = optValue
//	var tx *types.Transaction
//	//tx, err := dbContract.CreateSrcIntentOrderETH(
//	//	opts,
//	//	common.HexToAddress(owner),
//	//	common.HexToAddress(receiver),
//	//	contractParam.IntentCreatePays,
//	//	contractParam.IntentCallParam,
//	//)
//	// set contract data
//	contractParam.Wallet = common.HexToAddress(wallet)
//	contractParam.Receiver = common.HexToAddress(receiver)
//	contractParam.Owner = common.HexToAddress(owner)
//	contractParam.PayDB = common.HexToAddress(dbAddr)
//	contractParam.Amount = optValue
//	contractParam.ChainId = task.SrcChainId
//
//	// call contract
//	tx, err = service.SendIntentTransaction(contractParam)
//
//	if err != nil {
//		return "", fmt.Errorf("CreateSrcIntentOrderETH err: %v", err)
//	}
//	logger.Info("IntentTask(5/9): 🏆️ Call function createSrcOrder of payDBContract successfully and wait for receive txHash...")
//
//	// 6. wait for tx receipt
//	_, _ = bind.WaitMined(context.Background(), ethClient, tx)
//	logger.InfoF("IntentTask(6/9): ⛳️ Transaction(%s) on chain  successfully...\n", tx.Hash().String())
//
//	return tx.Hash().String(), nil
//}
//
//func createIntentDBSrcVw(
//	task model.Task,
//	contractParam model.ContractParam,
//) (model.Vw, error) {
//
//	// address info
//	account := service.GetAccount()
//	wallet := service.GetVW(task.SrcChainId)
//	vwManager := service.GetChainInfo(task.SrcChainId).Contract.Manager
//	receiver := service.GetAccount().Address
//	if task.HasDstVw {
//		receiver = service.GetVW(task.DstChainId)
//	}
//	// service data
//	data, _ := intentHybridServiceData(
//		contractParam.IntentCreatePays,
//		contractParam.IntentCallParam,
//		common.HexToAddress(account.Address),
//		common.HexToAddress(receiver),
//	)
//	code := getPayOrderId(task.SrcChainId, task.SrcChainId, task.Type)
//
//	// vw data
//	srcVw := model.Vw{
//		Code:          code.String(),
//		Wallet:        wallet,
//		Owner:         account.Address,
//		Sender:        wallet,
//		Service:       service.GetChainInfo(task.SrcChainId).Contract.IntentHybridPayService,
//		Data:          data,
//		Hash:          "",
//		WorkFlowHash:  "",
//		ChainId:       task.SrcChainId,
//		IsGateWay:     0,
//		GasToken:      service.GetTokenInfo(task.NodeId, task.SrcChainId, "USDT").TokenAddress,
//		GasTokenPrice: "35075000",
//		GasLimit:      "45075000",
//		PriorityFee:   getChainPriorityFee(task.NodeId, task.SrcChainId),
//		Manager:       vwManager,
//		Signature:     "",
//		App:           "Wallet",
//		Text:          "IntentHybrid",
//	}
//
//	// sign data
//	srcSignature, err := transferServiceSign(srcVw, task.SrcChainId)
//	if err != nil {
//		return model.Vw{}, err
//	}
//	srcVw.Signature = srcSignature
//	return srcVw, nil
//}
