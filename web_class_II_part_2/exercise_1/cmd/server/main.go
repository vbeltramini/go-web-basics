package main

import (
	"github.com/gin-gonic/gin"

	"vbeltramini.com/go-web-basics/web_class_II_part_2/exercise_1/cmd/server/handler"
	"vbeltramini.com/go-web-basics/web_class_II_part_2/exercise_1/internal/products"
)

func main() {
	productRepository := products.NewRepository()
	service := products.NewService(productRepository)
	productHandler := handler.NewProduct(service)

	router := gin.Default()
	productRouterGroup := router.Group("/products")
	{
		productRouterGroup.POST("/", productHandler.Store())
		productRouterGroup.GET("/", productHandler.GetAll())
	}
	router.Run()
}
