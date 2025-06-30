package main

import (
	"log"
	"os"

	"eth-for-babies-backend/internal/api/routes"
	"eth-for-babies-backend/internal/config"
	"eth-for-babies-backend/pkg/blockchain"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting Family Task Chain Backend...")

	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}
	log.Println("Environment variables loaded")

	// 初始化配置
	cfg := config.Load()
	log.Println("Configuration loaded")

	// 初始化数据库
	log.Println("Initializing database...")
	db, err := config.InitDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	log.Println("Database initialized successfully")

	// 初始化区块链客户端和合约管理器
	var contractManager *blockchain.ContractManager
	if cfg.Blockchain.RPCURL != "" && cfg.Blockchain.PrivateKey != "" {
		log.Println("Initializing blockchain client...")

		// 创建以太坊客户端
		ethClient, err := blockchain.NewEthClient(cfg.Blockchain.RPCURL, cfg.Blockchain.PrivateKey)
		if err != nil {
			log.Printf("Warning: Failed to initialize blockchain client: %v", err)
			log.Println("Continuing without blockchain functionality...")
		} else {
			// 初始化合约管理器，使用自定义的合约地址
			contractAddresses := map[string]string{
				"TaskRegistry":   cfg.Blockchain.TaskRegistryAddress,
				"FamilyRegistry": cfg.Blockchain.FamilyRegistryAddress,
				"RewardToken":    cfg.Blockchain.RewardTokenAddress,
				"RewardRegistry": cfg.Blockchain.RewardRegistryAddress,
			}

			contractManager, err = blockchain.NewContractManager(ethClient, contractAddresses)
			if err != nil {
				log.Printf("Warning: Failed to initialize contract manager: %v", err)
				log.Println("Continuing without blockchain functionality...")
				contractManager = nil
			} else {
				log.Println("Contract addresses:")
				log.Printf("  - Task Registry: %s", cfg.Blockchain.TaskRegistryAddress)
				log.Printf("  - Family Registry: %s", cfg.Blockchain.FamilyRegistryAddress)
				log.Printf("  - Reward Token: %s", cfg.Blockchain.RewardTokenAddress)
				log.Printf("  - Reward Registry: %s", cfg.Blockchain.RewardRegistryAddress)
				log.Println("Blockchain client and contract manager initialized successfully")
			}
		}
	} else {
		log.Println("Blockchain configuration not found, running without blockchain functionality")
	}

	// Set Gin mode based on configuration
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 初始化路由
	log.Println("Setting up routes...")
	router := routes.SetupRoutes(db, cfg, contractManager)
	log.Println("Routes setup completed")

	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Port
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
