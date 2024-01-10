package handlers

import (
	"github.com/fdrt29/product-app/pkg/services"
	"github.com/gin-gonic/gin"
)

type Layer struct {
	services *services.Layer
}

func NewLayer(services *services.Layer) *Layer {
	return &Layer{services: services}
}

func (l *Layer) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/v1")
	{
		products := api.Group("/products")
		{
			products.POST("/", l.createProduct)
			products.GET("/", l.getAllProducts)
			products.GET("/:id", l.getProductById)
			products.PUT("/:id", l.updateProduct)
			products.DELETE("/:id", l.deleteProduct)
		}
	}
	return router
}
