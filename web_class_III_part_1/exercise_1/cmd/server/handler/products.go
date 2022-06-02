package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vbeltramini/go-web-basics/web_class_III_part_1/exercise_1/internal/products"
)

type request struct {
	Name        string  `json:"name"`
	ProductType string  `json:"type"`
	Count       int     `json:"count"`
	Price       float64 `json:"price"`
}

type Product struct {
	service products.Service
}

func (product *Product) GetAll() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		if validateToken(c) {
			return
		}
		p, err := product.service.GetAll()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, p)
	}
	return fn
}

func (product *Product) Save() gin.HandlerFunc {
	fn := func(context *gin.Context) {
		if validateToken(context) {
			return
		}
		var req request
		if err := context.Bind(&req); err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		newProduct, err := product.service.Save(req.Name, req.ProductType, req.Count, req.Price)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, newProduct)
	}
	return fn
}

func (product *Product) Update() gin.HandlerFunc {
	fn := func(context *gin.Context) {
		if validateToken(context) {
			return
		}

		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(400, gin.H{"error": "id inválido"})
			return
		}
		var req request
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		switch true {
		case req.Name == "":
		case req.ProductType == "":
		case req.Count == 0:
		case req.Price == 0:
			context.JSON(400, gin.H{"error": "All fields need to have a valid content"})
			return
		}

		p, err := product.service.Update(id, req.Name, req.ProductType, req.Count, req.Price)
		if err != nil {
			context.JSON(404, gin.H{"error": err.Error()})
			return
		}
		context.JSON(200, p)
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

func NewProduct() *Product {
	return &Product{service: products.NewService()}
}
