package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vbeltramini/go-web-basics/web_class_III_part_1/exercise_1/cmd/server/handler"
)

func main() {
	productHandler := handler.NewProduct()
	gin := gin.Default()
	productsRouterGroup := gin.Group("/products")
	{
		productsRouterGroup.POST("/", productHandler.Save())
		productsRouterGroup.GET("/", productHandler.GetAll())
		productsRouterGroup.PUT("/:id", productHandler.Update())
	}
	gin.Run()
}
