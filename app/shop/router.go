package shop

import "github.com/gin-gonic/gin"

/**
 * @Author: lbh
 * @Date: 2021/4/7
 * @Description:
 */
func LoadShopRouter(e *gin.Engine) {
	e.GET("/shop/checkout", Checkout)
}
