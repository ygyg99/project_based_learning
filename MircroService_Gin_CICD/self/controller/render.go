package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 根据c中的请求头信息判断需要将数据渲染成哪种形式(HTML、JSON、XML)
// 在每个路由处理函数中添加如下函数
func Render(c *gin.Context, data gin.H, templateName string){
	// 在Gin中，传递给路由处理程序的Context 包含一个名为Request 的字段。
	// 这个字段包含了包含所有请求头的Header 字段。
	// 我们可以在Header 上使用Get 方法来提取Accept 头信息
	t := c.Request.Header.Get("Accept")
	switch t{
	case "application/json":
		// 只出传入data中key为payload的值
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK,data["payload"])
	default:
		c.HTML(http.StatusOK,templateName, data)
	}
}