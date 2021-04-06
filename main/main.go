package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
			"data":    "gin go go ",
		})
	})

	r.POST("/book", func(c *gin.Context) {

		username := c.PostForm("username")
		password := c.PostForm("password")
		c.GetRawData()

		c.JSON(200, gin.H{
			"message":  "this is POST method",
			"username": username,
			"password": password,
		})
	})

	r.GET("/urlQuery", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default name")
		c.JSON(200, gin.H{
			"name": name,
		})
	})

	r.Run()
}
