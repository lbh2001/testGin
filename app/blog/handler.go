package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: lbh
 * @Date: 2021/4/7
 * @Description:
 */
func Comment(c *gin.Context) {
	c.String(http.StatusOK, "this is blog comment handler^")
}
