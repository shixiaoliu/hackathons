package routes

import (
	"eth-for-babies-backend/internal/api/handlers"
	"eth-for-babies-backend/internal/api/middleware"
	"eth-for-babies-backend/internal/config"
	"eth-for-babies-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, cfg *config.Config) *gin.Engine {
	router := gin.New()

	// 创建JWT管理器
	jwtManager := utils.NewJWTManager(cfg.JWTSecret)

	// 全局中间件
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.Use(gin.Recovery())

	// 创建处理器
	authHandler := handlers.NewAuthHandler(db, jwtManager)
	familyHandler := handlers.NewFamilyHandler(db)
	childHandler := handlers.NewChildHandler(db)
	taskHandler := handlers.NewTaskHandler(db)
	contractHandler := handlers.NewContractHandler(db)

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
			}

			// 任务管理路由
			tasks := protected.Group("/tasks")
			{
				tasks.POST("", middleware.RequireRole("parent"), taskHandler.CreateTask)
				tasks.GET("", taskHandler.GetTasks)
				tasks.GET("/:id", taskHandler.GetTaskByID)
				tasks.PUT("/:id", middleware.RequireRole("parent"), taskHandler.UpdateTask)
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
		}
	}

	// 健康检查路由
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Family Task Chain Backend is running",
		})
	})

	return router
}