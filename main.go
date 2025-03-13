package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的 Gin 路由
	router := gin.Default()

	// 定义 /ping 接口
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pang")
	})

	// 定义 /fuck 接口
	router.GET("/fuck", func(c *gin.Context) {
		c.String(http.StatusOK, "motherfuker")
	})

	// 启动服务器，监听 8080 端口
	router.Run("0.0.0.0:8080") // 监听所有可用接口，包括 IPv4

}
