package main

import (
	"log"

	"github.com/vbeltramini/go-web-basics/web_class_III_part_2/exercise_1/cmd/server/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("failed to load .env")
	}
	productHandler := handler.NewProduct()
	gin := gin.Default()
	productsRouterGroup := gin.Group("/products")
	{

		productsRouterGroup.Use(productHandler.AuthToken)

		productsRouterGroup.POST("/", productHandler.Save())
		productsRouterGroup.GET("/", productHandler.GetAll())
		productsRouterGroup.GET("/:id", productHandler.GetById())
		productsRouterGroup.PUT("/:id", productHandler.Update())
		productsRouterGroup.DELETE("/:id", productHandler.Delete())
		productsRouterGroup.PATCH("/:id", productHandler.PatchNamePrice())
	}
	gin.Run()
}
