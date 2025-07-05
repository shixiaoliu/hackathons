package routes

import (
	"eth-for-babies-backend/internal/api/handlers"
	"eth-for-babies-backend/internal/api/middleware"
	"eth-for-babies-backend/internal/config"
	"eth-for-babies-backend/internal/repository"
	"eth-for-babies-backend/internal/services"
	"eth-for-babies-backend/internal/utils"
	"eth-for-babies-backend/pkg/blockchain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, cfg *config.Config, contractManager *blockchain.ContractManager) *gin.Engine {
	router := gin.New()

	// 创建JWT管理器
	jwtManager := utils.NewJWTManager(cfg.JWTSecret)

	// 全局中间件
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.Use(gin.Recovery())

	// 创建仓库
	rewardRepo := repository.NewRewardRepository(db)
	exchangeRepo := repository.NewExchangeRepository(db)
	childRepo := repository.NewChildRepository(db)
	familyRepo := repository.NewFamilyRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	// 创建服务
	contractService, _ := services.NewContractService(&cfg.Blockchain, contractManager)
	rewardService := services.NewRewardService(rewardRepo, exchangeRepo, childRepo, contractManager)
	childService := services.NewChildService(childRepo, familyRepo, taskRepo)

	// 创建处理器
	authHandler := handlers.NewAuthHandler(db, jwtManager)
	familyHandler := handlers.NewFamilyHandler(db)
	childHandler := handlers.NewChildHandler(db)
	taskHandler := handlers.NewTaskHandler(db, contractManager)
	contractHandler := handlers.NewContractHandler(db, contractService)
	rewardHandler := handlers.NewRewardHandler(rewardService)
	exchangeHandler := handlers.NewExchangeHandler(rewardService, childService)

	// API v1 路由组
	v1 := router.Group("/api/v1")
	{
		// 认证路由（无需认证）
		auth := v1.Group("/auth")
		{
			auth.GET("/nonce/:wallet_address", authHandler.GetNonce)
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
		}

		// 需要认证的路由
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware(jwtManager))
		{
			// 认证相关（需要认证）
			protected.POST("/auth/logout", authHandler.Logout)

			// 家庭管理路由
			families := protected.Group("/families")
			{
				families.POST("", middleware.RequireRole("parent"), familyHandler.CreateFamily)
				families.GET("", familyHandler.GetFamilies)
				families.GET("/:id", familyHandler.GetFamilyByID)
				families.PUT("/:id", middleware.RequireRole("parent"), familyHandler.UpdateFamily)
			}

			// 孩子管理路由
			children := protected.Group("/children")
			{
				children.POST("", middleware.RequireRole("parent"), childHandler.CreateChild)
				children.GET("/my", childHandler.GetChildren)
				children.GET("/:id", childHandler.GetChildByID)
				children.PUT("/:id", childHandler.UpdateChild)
				children.GET("/:id/progress", childHandler.GetChildProgress)
				children.DELETE("/:id", middleware.RequireRole("parent"), childHandler.DeleteChild)
			}

			// 任务管理路由
			tasks := protected.Group("/tasks")
			{
				tasks.POST("", middleware.RequireRole("parent"), taskHandler.CreateTask)
				tasks.GET("", taskHandler.GetTasks)
				tasks.GET("/:id", taskHandler.GetTaskByID)
				tasks.PUT("/:id", middleware.RequireRole("parent"), taskHandler.UpdateTask)
				tasks.POST("/:id/direct-assign", middleware.RequireRole("parent"), taskHandler.DirectAssignTask)
				tasks.POST("/:id/complete", middleware.RequireRole("child"), taskHandler.CompleteTask)
				tasks.POST("/:id/approve", middleware.RequireRole("parent"), taskHandler.ApproveTask)
				tasks.POST("/:id/reject", middleware.RequireRole("parent"), taskHandler.RejectTask)
			}

			// 智能合约路由
			contracts := protected.Group("/contracts")
			{
				contracts.GET("/balance/:address", contractHandler.GetBalance)
				contracts.POST("/transfer", contractHandler.Transfer)
				contracts.GET("/transactions/:hash", contractHandler.GetTransactionStatus)
			}

			// 奖品管理路由
			rewards := protected.Group("/rewards")
			{
				rewards.POST("/family/:family_id", middleware.RequireRole("parent"), rewardHandler.CreateReward)
				rewards.GET("/family/:family_id", rewardHandler.GetRewards)
				rewards.GET("/:id", rewardHandler.GetRewardByID)
				rewards.PUT("/:id", middleware.RequireRole("parent"), rewardHandler.UpdateReward)
				rewards.DELETE("/:id", middleware.RequireRole("parent"), rewardHandler.DeleteReward)
			}

			// 兑换管理路由
			exchanges := protected.Group("/exchanges")
			{
				exchanges.POST("", middleware.RequireRole("child"), exchangeHandler.ExchangeReward)
				exchanges.GET("/my", middleware.RequireRole("child"), exchangeHandler.GetChildExchanges)
				exchanges.GET("/:id", exchangeHandler.GetExchangeByID)
				exchanges.PUT("/:id/status", middleware.RequireRole("parent"), exchangeHandler.UpdateExchangeStatus)
			}

			// 家庭兑换记录路由
			protected.GET("/exchanges/family/:family_id", exchangeHandler.GetFamilyExchanges)
		}

		// 在 v1 路由组下添加健康检查路由
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Family Task Chain Backend is running",
			})
		})
	}

	return router
}
