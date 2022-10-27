package routers

import (
	"assignment2/controllers/order"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/order", order.CreateOrder)

	router.GET("/order/:id", order.GetOrder)
	router.GET("/orders", order.GetOrders)

	router.PUT("/order/:id", order.UpdateOrder)
	router.DELETE("/order/:id", order.DeleteOrder)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
