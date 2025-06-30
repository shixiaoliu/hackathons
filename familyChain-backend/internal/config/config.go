package config

import (
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Config struct {
	Port                  string
	Environment           string
	LogLevel              string
	JWTSecret             string
	BlockchainRPCURL      string
	PrivateKey            string
	TaskContractAddress   string
	FamilyContractAddress string
	TokenContractAddress  string
	RewardContractAddress string
	Database              DatabaseConfig
	Blockchain            BlockchainConfig
}

type DatabaseConfig struct {
	Driver string
	DSN    string
}

type BlockchainConfig struct {
	RPCURL                string
	PrivateKey            string
	TaskRegistryAddress   string
	FamilyRegistryAddress string
	RewardTokenAddress    string
	RewardRegistryAddress string
	ChainID               int64
	Client                *ethclient.Client
}

func Load() *Config {
	chainID, _ := strconv.ParseInt(getEnv("BLOCKCHAIN_CHAIN_ID", "1337"), 10, 64)

	return &Config{
		Port:                  getEnv("PORT", "8080"),
		Environment:           getEnv("ENVIRONMENT", "development"),
		LogLevel:              getEnv("LOG_LEVEL", "info"),
		JWTSecret:             getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		BlockchainRPCURL:      getEnv("BLOCKCHAIN_RPC_URL", "http://localhost:8545"),
		PrivateKey:            getEnv("BLOCKCHAIN_PRIVATE_KEY_PARENT", ""),
		TaskContractAddress:   getEnv("TASK_CONTRACT_ADDRESS", ""),
		FamilyContractAddress: getEnv("FAMILY_CONTRACT_ADDRESS", ""),
		TokenContractAddress:  getEnv("TOKEN_CONTRACT_ADDRESS", ""),
		RewardContractAddress: getEnv("REWARD_CONTRACT_ADDRESS", ""),
		Database: DatabaseConfig{
			Driver: getEnv("DB_DRIVER", "sqlite"),
			DSN:    getEnv("DB_DSN", "./data/family_task_chain.db"),
		},
		Blockchain: BlockchainConfig{
			RPCURL:                getEnv("BLOCKCHAIN_RPC_URL", "http://localhost:8545"),
			PrivateKey:            getEnv("BLOCKCHAIN_PRIVATE_KEY_PARENT", ""),
			TaskRegistryAddress:   getEnv("TASK_CONTRACT_ADDRESS", ""),
			FamilyRegistryAddress: getEnv("FAMILY_CONTRACT_ADDRESS", ""),
			RewardTokenAddress:    getEnv("TOKEN_CONTRACT_ADDRESS", ""),
			RewardRegistryAddress: getEnv("REWARD_CONTRACT_ADDRESS", ""),
			ChainID:               chainID,
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
