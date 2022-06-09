package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vbeltramini/go-web-basics/web_class_IV_part_2/exercise_1/internal/products"
	"github.com/vbeltramini/go-web-basics/web_class_IV_part_2/exercise_1/pkg/web"
)

type productRequest struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}

const (
	ERROR_NAME  = "o nome do produto é obrigatório"
	ERROR_TYPE  = "o tipo do produto é obrigatório"
	ERROR_COUNT = "a quantidade do produto é obrigatória"
	ERROR_PRICE = "o preço do produto é obrigatório"
	ERROR_TOKEN = "token inválido"
	ERROR_ID    = "id inválido"
)

type Product struct {
	service products.Service
}

func (p *Product) AuthToken(ctx *gin.Context) {
	privateToken := os.Getenv("TOKEN")

	providedToken := ctx.GetHeader("token")

	if providedToken != privateToken {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token provided"})
		return
	}

	ctx.Next()
}

func (p *Product) GetAll(ctx *gin.Context) {
	products, err := p.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound,
			web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, products, ""))
}

// GetById Update UpdateProducts godoc
// @Summary Get products by ID
// @Tags Products
// @Description get product by id
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "Some ID"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "We need ID"
// @Failure 404 {object} web.Response "Can not find ID"
// @Router /products/{:id} [Get]
func (p *Product) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	prod, err := p.service.GetById(id)
	if err != nil {
		web.NewResponse(http.StatusNotFound, nil, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, prod)
}

// Save godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body productRequest true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (p *Product) Save(ctx *gin.Context) {
	var req productRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusNotFound,
			web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	products, err := p.service.Save(req.Name, req.Type, req.Count, req.Price)
	if err != nil {
		ctx.JSON(http.StatusNotFound,
			web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, products, ""))
}

// Update UpdateProducts godoc
// @Summary Update products by ID
// @Tags Products
// @Description update products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param {:id} path int true "Some ID"
// @Param product body productRequest true "Product to update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "We need ID"
// @Failure 404 {object} web.Response "Can not find ID"
// @Router /products/{:id} [PUT]
func (p *Product) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			web.NewResponse(http.StatusBadRequest, nil, ERROR_ID))
		return
	}
	var req productRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,
			web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	if req.Name == "" {
		ctx.JSON(http.StatusBadRequest,
			web.NewResponse(http.StatusBadRequest, nil, ERROR_NAME))
		return
	}
	if req.Type == "" {
		ctx.JSON(http.StatusBadRequest,
			web.NewResponse(http.StatusBadRequest, nil, ERROR_TYPE))
		return
	}
	if req.Count == 0 {
		ctx.JSON(http.StatusBadRequest,
			web.NewResponse(http.StatusBadRequest, nil, ERROR_COUNT))
		return
	}
	if req.Price == 0 {
		ctx.JSON(http.StatusBadRequest,
			web.NewResponse(http.StatusBadRequest, nil, ERROR_PRICE))
		return
	}
	newProduct, err := p.service.Update(int(id), req.Name, req.Type, req.Count, req.Price)
	if err != nil {
		ctx.JSON(http.StatusNotFound,
			web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, newProduct, ""))
}

// PatchNamePrice UpdateNameProducts godoc
// @Summary Update name products by ID
// @Tags Products
// @Description update the name of the products by ID
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param some_id path int true "Some ID"
// @Param product body productRequest true "Product to update name"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "We need ID"
// @Failure 404 {object} web.Response "Can not find ID"
// @Router /products/{:id} [PATCH]
func (p *Product) PatchNamePrice(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		web.NewResponse(http.StatusBadRequest, nil, "Invalid id")
		return
	}
	var req productRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	switch true {
	case req.Name == "":
	case req.Price == 0:
		web.NewResponse(http.StatusBadRequest, nil, ERROR_NAME)
		return
	}

	product, err := p.service.PatchNamePrice(id, req.Name, req.Price)
	if err != nil {
		context.JSON(404, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, product)
}

// Delete DeleteProducts godoc
// @Summary Delete products by ID
// @Tags Products
// @Description delete products by ID
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param some_id path int true "Some ID"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "We need ID"
// @Failure 404 {object} web.Response "Can not find ID"
// @Router /products/{:id} [DELETE]
func (p *Product) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			web.NewResponse(http.StatusBadRequest, nil, ERROR_ID))
		return
	}
	err = p.service.Delete(int(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound,
			web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, fmt.Sprintf("o produto %d foi removido", id), ""))
}

func NewProduct() *Product {
	return &Product{service: products.NewService()}
}
