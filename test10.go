package main

// 学习Gin框架
// Gin是一个高性能的Web框架，使用Go语言编写，具有极快的速度和灵活性。它提供了路由、参数解析、中间件、JSON序列化等功能，适合构建RESTful API和Web应用程序。

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test10() {
	engine := gin.Default() //创建gin引擎
	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	if err := engine.Run(); err != nil { //开启服务器，默认监听localhost:8080
		panic(err) // 如果启动失败，输出错误并退出程序
	}
}
