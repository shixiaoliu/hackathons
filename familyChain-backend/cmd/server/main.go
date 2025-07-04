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
			// 检查合约地址配置
			checkContractAddresses(cfg)

			// 初始化合约管理器，使用自定义的合约地址
			contractAddresses := map[string]string{
				"task":   cfg.Blockchain.TaskRegistryAddress,
				"family": cfg.Blockchain.FamilyRegistryAddress,
				"token":  cfg.Blockchain.RewardTokenAddress,
				"reward": cfg.Blockchain.RewardRegistryAddress,
			}

			contractManager, err = blockchain.NewContractManager(ethClient, contractAddresses)
			if err != nil {
				log.Printf("Warning: Failed to initialize contract manager: %v", err)
				log.Println("Continuing without blockchain functionality...")
				contractManager = nil
			} else {
				// 获取合约地址
				addresses := contractManager.GetContractAddresses()

				log.Println("Contract addresses:")
				log.Printf("  - Task Registry: %s", addresses["task"])
				log.Printf("  - Family Registry: %s", addresses["family"])
				log.Printf("  - Reward Token: %s", addresses["token"])
				log.Printf("  - Reward Registry: %s", addresses["reward"])

				// 验证RewardToken合约是否正确初始化
				if contractManager.RewardToken == nil {
					log.Printf("Warning: RewardToken contract not initialized, trying to initialize...")
					if addresses["token"] != "0x0000000000000000000000000000000000000000" && addresses["token"] != "" {
						err := contractManager.InitRewardToken(addresses["token"])
						if err != nil {
							log.Printf("Error initializing RewardToken: %v", err)
						} else {
							log.Printf("Successfully initialized RewardToken contract!")
						}
					} else {
						log.Printf("Error: RewardToken address is empty or zero address")
					}
				}

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

// 检查合约地址配置
func checkContractAddresses(cfg *config.Config) {
	if cfg.Blockchain.TaskRegistryAddress == "" {
		log.Printf("Warning: Task Registry address not configured. Set TASK_CONTRACT_ADDRESS env variable.")
	}
	if cfg.Blockchain.FamilyRegistryAddress == "" {
		log.Printf("Warning: Family Registry address not configured. Set FAMILY_CONTRACT_ADDRESS env variable.")
	}
	if cfg.Blockchain.RewardTokenAddress == "" {
		log.Printf("Warning: Reward Token address not configured. Set TOKEN_CONTRACT_ADDRESS env variable.")
	}
	if cfg.Blockchain.RewardRegistryAddress == "" {
		log.Printf("Warning: Reward Registry address not configured. Set REWARD_CONTRACT_ADDRESS env variable.")
	}
}
