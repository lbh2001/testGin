package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: lbh
 * @Date: 2021/4/7
 * @Description:
 */
func Checkout(c *gin.Context) {
	c.String(http.StatusOK, "this is shop checkout handler^")
}
