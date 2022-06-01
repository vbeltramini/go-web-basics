package main

import (
	"github.com/gin-gonic/gin"

	"vbeltramini.com/go-web-basics/web_class_II_part_2/exercise_1/cmd/server/handler"
	"vbeltramini.com/go-web-basics/web_class_II_part_2/exercise_1/internal/products"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	r.Run()
}
