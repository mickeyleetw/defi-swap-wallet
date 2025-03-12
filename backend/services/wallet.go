package services

import (
	"context"
	"log"
	"math/big"

	"defi-swap-backend/config"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type WalletService struct {
	client *ethclient.Client
}

func NewWalletService() *WalletService {
	client, err := ethclient.Dial(config.SepoliaRPCURL)
	if err != nil {
		log.Fatal(err)
	}
	return &WalletService{client: client}
}

func (ws *WalletService) GetBalance(address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := ws.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
