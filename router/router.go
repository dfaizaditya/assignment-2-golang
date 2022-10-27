package router

import (
	controllers "assignment-2/controller"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrders)
	router.DELETE("/orders/:OrderID", controllers.DeleteOrder)
	router.PUT("/orders/:OrderID", controllers.UpdateOrder)

	return router
}
