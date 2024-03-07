package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vpcmps/finanseed/docs"
	"github.com/vpcmps/finanseed/handler"
	"github.com/vpcmps/finanseed/handler/transaction"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	docs.SwaggerInfo.Title = "Finanseed"
	docs.SwaggerInfo.Description = "Finanseed API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group("/api/v1")
	{
		v1.GET("/transactions", handler.ListOpeningsHandler)
		v1.GET("/transaction", handler.GetOpeningHandler)
		v1.POST("/transaction", transaction.CreateTransactionHandler)
		v1.DELETE("/transaction", handler.DeleteOpeningHandler)
		v1.PUT("/transaction", handler.UpdateOpeningHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
