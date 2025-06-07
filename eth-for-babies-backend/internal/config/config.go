package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port        string
	Environment string
	LogLevel    string
	JWTSecret   string
	Database    DatabaseConfig
	Blockchain  BlockchainConfig
}

type DatabaseConfig struct {
	Driver string
	DSN    string
}

type BlockchainConfig struct {
	RPCURL          string
	PrivateKey      string
	ContractAddress string
	ChainID         int64
}

func Load() *Config {
	chainID, _ := strconv.ParseInt(getEnv("BLOCKCHAIN_CHAIN_ID", "1337"), 10, 64)

	return &Config{
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		Database: DatabaseConfig{
			Driver: getEnv("DB_DRIVER", "sqlite"),
			DSN:    getEnv("DB_DSN", "./data/app.db"),
		},
		Blockchain: BlockchainConfig{
			RPCURL:          getEnv("BLOCKCHAIN_RPC_URL", "http://localhost:8545"),
			PrivateKey:      getEnv("BLOCKCHAIN_PRIVATE_KEY", ""),
			ContractAddress: getEnv("CONTRACT_ADDRESS", ""),
			ChainID:         chainID,
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}