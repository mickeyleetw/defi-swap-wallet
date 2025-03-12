package services

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"sync"

	"defi-swap-backend/config"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactionService struct {
	client *ethclient.Client
	mu     sync.Mutex
}

func NewTransactionService() *TransactionService {
	client, err := ethclient.Dial(config.SepoliaRPCURL)
	if err != nil {
		log.Fatal(err)
	}
	return &TransactionService{client: client}
}

func (ts *TransactionService) SendTransaction(to string, amount int64) (string, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	nonce, err := ts.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	value := big.NewInt(amount)
	gasLimit := uint64(21000)
	gasPrice, err := ts.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	toAddress := common.HexToAddress(to)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := ts.client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	err = ts.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}
