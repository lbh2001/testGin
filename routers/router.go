package routers

import "github.com/gin-gonic/gin"

/**
 * @Author: lbh
 * @Date: 2021/4/7
 * @Description:
 */

//func HelloHandler(c *gin.Context)  {
//	c.JSON(http.StatusOK,gin.H{
//		"message": "hello handler~",
//	})
//}
//
//func Startup() *gin.Engine{
//	r := gin.Default()
//	r.GET("/handler",HelloHandler)
//	return r
//}
type Option func(e *gin.Engine)

var options = []Option{}

func Include(opt ...Option) {
	options = append(options, opt...)
}

func Init() *gin.Engine {
	r := gin.New()
	for _, opt := range options {
		opt(r)
	}
	return r
}
