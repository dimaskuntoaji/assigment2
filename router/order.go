package router

import (
	ct "restapi2/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRouter(router gin.Engine) *gin.Engine {
	router.POST("/orders", ct.PostOrder)
	router.GET("/orders", ct.GetOrders)
	router.PUT("/orders/:orderId", ct.PutOrder)
	router.DELETE("/orders/:orderId", ct.DeleteOrder)

	return &router
}