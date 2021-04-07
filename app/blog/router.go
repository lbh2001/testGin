package blog

import "github.com/gin-gonic/gin"

/**
 * @Author: lbh
 * @Date: 2021/4/7
 * @Description:
 */

func LoadBlogRouter(e *gin.Engine) {
	e.GET("/blog/comment", Comment)
}
