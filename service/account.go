// @Author fcihpy
// @Date 2024/7/29 17:38:00
// @Desc
//

package service

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"github.com/DappOSDao/node-service/network"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fcihpy001/chain-service/constant"
	"github.com/fcihpy001/chain-service/contract"
	"github.com/fcihpy001/chain-service/logger"
	"github.com/fcihpy001/chain-service/model"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"strings"
)

func init() {
	loadEnv()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		logger.PanicF("global setting(1/5): can't load .env file: %v", err)
	}
	logger.Info("global setting(1/5): env file load success...")
}

var account *model.Account

func GetAccount() *model.Account {
	if account != nil {
		return account
	}
	privKey := os.Getenv("PRIVATEKEY")
	if len(strings.TrimSpace(privKey)) == 0 {
		logger.Panic("Load private key Error")
	}

	// ecdsa privateKey
	ecdsaPrivKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		logger.Panic("convert private key Error")
	}

	// address
	address := PrivateKey2Address(ecdsaPrivKey)
	if len(address) != 42 {
		logger.Panic("generator account address Error")
	}

	// common vw addr
	vw, err := getVwAddr(10, address)
	if err != nil {
		logger.Panic("get vw Addr Error")
	}

	//manta vw addr
	mantaVwAddr, err := getVwAddr(169, address)
	if err != nil {
		logger.Error("Get manta Vw Addr error")
	}

	//manta vw addr
	zksyncVwAddr, err := getVwAddr(324, address)
	if err != nil {
		logger.Error("Get zksync Vw Addr error")
	}

	account = &model.Account{
		Address:    address,
		Vw:         vw,
		MantaVw:    mantaVwAddr,
		ZkSyncVw:   zksyncVwAddr,
		PrivateKey: ecdsaPrivKey,
	}
	logger.Debug("current user's account info:",
		zap.Any("address", account.Address),
		zap.Any("vw", account.Vw),
		zap.Any("mantaVw", account.MantaVw),
		zap.Any("zkVw", account.ZkSyncVw))
	return account
}

func GetVW(chainId int) string {
	ac := GetAccount()
	vw := ac.Vw
	if chainId == 169 {
		vw = ac.MantaVw
	} else if chainId == 324 {
		vw = ac.ZkSyncVw
	}
	return vw
}

func PrivateKey2Address(privateKey *ecdsa.PrivateKey) string {

	// private -> publicKey
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.Error("error casting public key to ECDSA")
	}

	// publicKey -> address
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address
}

func getVwAddr(chainId int, wallet string) (string, error) {
	client, err := GetEvmClient(GetRPC(chainId), chainId)
	if err != nil {
		return "", err
	}

	vwManager := network.GetChainInfo(chainId).Contract.Manager
	vwContract, err := contract.NewVWManager(common.HexToAddress(vwManager), client)
	if err != nil {
		return "", err
	}

	vwAddr, err := vwContract.OwnerWallet(nil, common.HexToAddress(wallet))
	if err != nil {
		return "", err
	}

	if vwAddr.String() == constant.ZeroAddr {
		return "", errors.New("vw not found")
	}
	return vwAddr.String(), nil
}

func PrivKey2wallet(privKey string) string {
	// ecdsa privateKey
	ecdsaPrivKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		logger.Panic("convert private key Error")
	}

	// address
	address := PrivateKey2Address(ecdsaPrivKey)
	if len(address) != 42 {
		logger.Panic("generator account address Error")
	}
	return address
}

func Keystore2wallet(passwd string, path string) string {
	keyByte, err := os.ReadFile(path)
	if err != nil {
		logger.PanicF("read file error: %s", err.Error())
	}

	privKey, err := keystore.DecryptKey(keyByte, passwd)
	if err != nil {
		logger.PanicF("key decode error: %s", err.Error())
	}

	// ecdsa privateKey
	keyBytes := crypto.FromECDSA(privKey.PrivateKey)
	if err != nil {
		logger.Panic("convert private key Error")
	}

	// privateKey
	return hex.EncodeToString(keyBytes)
}

func PriKey2keystore(privKey string, passwd string, path string) {
	// ecdsa privateKey
	ecdsaPrivKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		logger.Panic("convert private key Error")
	}

	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	addr, err := ks.ImportECDSA(ecdsaPrivKey, passwd)

	if err != nil {
		logger.PanicF("make keystore error:", err.Error())
	}
	logger.DebugF("wallet: %s", addr)
}

func Keystore2privateKey(password string, path string) (*ecdsa.PrivateKey, error) {
	keyByte, err := os.ReadFile(path)
	if err != nil {
		logger.PanicF("read file error: %s", err.Error())
	}

	privKey, err := keystore.DecryptKey(keyByte, password)
	if err != nil {
		logger.PanicF("key decode error: %s", err.Error())
	}

	// ecdsa privateKey
	keyBytes := crypto.FromECDSA(privKey.PrivateKey)
	if err != nil {
		logger.Panic("convert private key Error")
	}
	return crypto.HexToECDSA(hex.EncodeToString(keyBytes))
}
