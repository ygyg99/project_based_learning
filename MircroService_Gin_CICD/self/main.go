package main

import (
	"MicroService/controller"

	"github.com/gin-gonic/gin"
)

func main(){
	path := "D:/code/project_based_learning/MircroService_Gin_CICD/self"
	// 注册路由
	router := gin.Default()
	// 加载模板
	router.LoadHTMLGlob(path+"/templates/*")
	// 定义路由响应程序
	router.GET("/",controller.Home_new)
	router.GET("/article/view/:article_id",controller.GetArticle)
	router.Run(":8084") // 指定运行的端口为8084
}