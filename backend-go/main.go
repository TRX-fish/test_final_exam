package main

import (
	"backend-go/config"
	"backend-go/database"
	"backend-go/handlers"
	"backend-go/middleware"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	if err := config.LoadConfig(); err != nil {
		log.Fatal("加载配置失败:", err)
	}

	// 初始化数据库
	if err := database.InitDB(); err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 创建 Gin 实例
	r := gin.Default()

	// 配置 CORS
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	r.Use(cors.New(corsConfig))

	// 静态文件服务
	r.Static("/static", "./static")

	// API 路由
	api := r.Group("/api")
	{
		// 用户相关路由
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
		api.GET("/user/me", middleware.JWTAuth(), handlers.GetCurrentUser)

		// 试卷相关路由
		api.POST("/papers", handlers.AddPaper)
		api.PUT("/papers/:id/image", handlers.UpdatePaperImage)
		api.GET("/papers", handlers.GetPapers)
		api.GET("/stats", handlers.GetStats)

		// 题目相关路由
		api.POST("/questions", handlers.AddQuestion)
		api.PUT("/questions/:id/image", handlers.UpdateQuestionImage)
		api.GET("/questions", handlers.GetQuestions)
		api.GET("/questions/search", handlers.SearchQuestions)

		// 上传相关路由
		api.POST("/upload/paper-image", handlers.UploadPaperImage)
		api.POST("/upload/question-image", handlers.UploadQuestionImage)
		api.DELETE("/upload/delete-image", handlers.DeleteImage)
	}

	// 启动服务器
	log.Println("服务器启动在 http://localhost:5000")
	if err := r.Run(":5000"); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
} 