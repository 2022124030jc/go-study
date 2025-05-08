package main

// 测试swagger接口文档生成
// swagger是一个用于描述和文档化RESTful API的工具。它使用OpenAPI规范来定义API的结构，包括端点、请求参数、响应格式等。Swagger提供了可视化界面，方便开发者和用户理解和测试API。
// 使用swag init生成接口文档
// 访问http://localhost:80/swagger/index.html查看接口文档
// 访问http://localhost:80/api/v1/ping?name=world查看接口返回结果

import (
	"fmt"
	// 匿名导入生成的接口文档包

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go-study/docs" // Removed as the package is not available
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @BasePath  /api/v1
func Test11() {
	engine := gin.Default()
	// 注册swagger静态文件路由
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("/api/v1/ping", Ping)
	engine.Run(":80")
}

// Ping godoc
// @Summary      say hello world
// @Description  return hello world json format content
// @param       name query    string  true  "name"
// @Tags         system
// @Produce      json
// @Router       /ping [get]
func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": fmt.Sprintf("Hello World!%s", ctx.Query("name")),
	})
}
