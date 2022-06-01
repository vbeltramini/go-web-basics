package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vbeltramini.com/go-web-basics/web_class_II_part_2/exercise_1/internal/products"
)

type request struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}

type Product struct {
	service products.Service
}

func (productService *Product) GetAll() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		if validateToken(c) {
			return
		}
		p, err := productService.service.GetAll()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, p)
	}
	return fn
}

func (productService *Product) Store() gin.HandlerFunc {
	fn := func(context *gin.Context) {
		if validateToken(context) {
			return
		}
		var req request
		if err := context.Bind(&req); err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		newProduct, err := productService.service.Store(req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, newProduct)
	}
	return fn
}

func validateToken(context *gin.Context) bool {
	token := context.Request.Header.Get("token")
	if token != "csgo" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized token"})
		return true
	}
	return false
}

func NewProduct(productService products.Service) *Product {
	return &Product{service: productService}
}
