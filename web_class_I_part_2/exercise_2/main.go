package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Exercício 2 - Get one endpoint
//Gere um novo endpoint que nos permita buscar um único resultado do array de temas.
//Usando parâmetros de caminho o endpoint deve ser /theme/:id (lembre-se que o tema sempre tem que ser plural). Uma vez que o id é recebido, ele retorna a posição correspondente.
// 1.Gere uma nova rota.
// 2.Gera um manipulador para a rota criada.
// 3.Dentro do manipulador, procure o item que você precisa.
// 4.Retorna o item de acordo com o id.
//Se você não encontrou nenhum elemento com esse id retorne como código de resposta 404.

type user struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	Age          int     `json:"age"`
	Height       float32 `json:"height"`
	Active       bool    `json:"active"`
	CreationDate string  `json:"creation_date"`
}

func main() {
	u := createUsersData()
	gin.SetMode("release")
	router := gin.Default()
	group := router.Group("users")
	{
		group.GET("/", getAll(u))
		group.GET("/:id", getById(u))
	}
	router.Run()
}

func getById(usuarios []user) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")
		id2, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "id is not a number",
			})
			log.Println(err)
			return
		}
		for _, u := range usuarios {
			fmt.Println(u.Id)
			if u.Id == id2 {
				c.JSON(200, u)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{
			"message": "id not found",
		})
	}
	return gin.HandlerFunc(fn)
}

func getAll(usuarios []user) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		for _, user := range usuarios {
			c.JSON(200, user)
		}
	}
	return gin.HandlerFunc(fn)
}

func createUsersData() []user {
	jsonData := []byte(`
    [
        {"id":77,"name":"Victor","lastname":"Hugo","email":"victor@gmail.com","age":19,"height":1.8,"active":true,"creationdate":"30/05/2022"},
        {"id":10,"name":"Victor","lastname":"Hugo","email":"victor@gmail.com","age":19,"height":1.8,"active":true,"creationdate":"30/05/2022"},
        {"id":15,"name":"Victor","lastname":"Hugo","email":"victor@gmail.com","age":19,"height":1.8,"active":true,"creationdate":"30/05/2022"},
        {"id":100,"name":"Victor","lastname":"Hugo","email":"victor@gmail.com","age":19,"height":1.8,"active":true,"creationdate":"30/05/2022"}
    ]`)

	var u []user

	if err := json.Unmarshal(jsonData, &u); err != nil {
		log.Fatal(err)
	}
	return u
}
