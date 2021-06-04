package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(Cors())
	upfile := r.Group("/upfile")
	{
		upfile.POST("/add", UpfileHandle.Upfile)
		upfile.Static(`/data`, "data")
	}
	lable := r.Group("/api/lable")
	{
		lable.POST("/add", LableHandle.LableAddHandle)
		lable.GET("/info/:id", LableHandle.LableInfoHandle)
		lable.GET("/list", LableHandle.LableListHandle)
		lable.POST("/edit", LableHandle.LableEditHandle)
		lable.POST("/del/:id", LableHandle.LableDelHandle)
	}
	category := r.Group("/api/category")
	{
		category.POST("/add", CategoryHandle.CategoryAddHandle)
		category.GET("/info/:id", CategoryHandle.CategoryInfoHandle)
		category.GET("/list", CategoryHandle.CategoryListHandle)
		category.GET("/listson", CategoryHandle.CategoryListSonHandle) //获取二级分类
		category.POST("/edit", CategoryHandle.CategoryEditHandle)
		category.POST("/del/:id", CategoryHandle.CategoryDelHandle)
	}
	text := r.Group("/api/text")
	{
		text.POST("/add", TextHandle.TextAddHandle)
		text.GET("/info/:id", TextHandle.TextInfoHandle)
		text.GET("/list", TextHandle.TextListHandle)
		text.POST("/edit", TextHandle.TextEditHandle)
		text.POST("/del/:id", TextHandle.TextDeletHandle)
	}
	r.Run(HandPort)
}
