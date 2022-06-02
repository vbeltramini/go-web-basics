package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vbeltramini/go-web-basics/web_class_III_part_1/exercise_2/cmd/server/handler"
)

func main() {
	productHandler := handler.NewProduct()
	gin := gin.Default()
	productsRouterGroup := gin.Group("/products")
	{
		productsRouterGroup.POST("/", productHandler.Save())
		productsRouterGroup.GET("/", productHandler.GetAll())
		productsRouterGroup.GET("/:id", productHandler.GetById())
		productsRouterGroup.PUT("/:id", productHandler.Update())
		productsRouterGroup.DELETE("/:id", productHandler.Delete())
		productsRouterGroup.PATCH("/:id", productHandler.PatchNamePrice())
	}
	gin.Run()
}
