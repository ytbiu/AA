package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载 .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	// 初始化 Gin
	r := gin.Default()

	// 基础健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 启动服务
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server starting on port", port)
	r.Run(":" + port)
}
