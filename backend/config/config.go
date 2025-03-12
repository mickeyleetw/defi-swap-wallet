package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SepoliaRPCURL string
	PrivateKey    string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	SepoliaRPCURL = os.Getenv("SEPOLIA_RPC_URL")
	PrivateKey = os.Getenv("WALLET_PRIVATE_KEY")
}
