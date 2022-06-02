package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Exercício 1 - Vamos filtrar nosso endpoint

// Dependendo do tema escolhido, precisamos adicionar filtros ao nosso endpoint, ele deve ser
// capaz de filtrar todos os campos.
// 1. Dentro do manipulador de endpoint, recebi os valores para filtrar do contexto.
// 2. Em seguida, ele gera a lógica do filtro para nossa matriz.
// 3. Retorne a matriz filtrada por meio do endpoint.

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
		group.GET("/", getAllUsers(u))
		group.GET("/filter", getUserById())
	}
	router.Run()
}

func getUserById() gin.HandlerFunc {
	return func(context *gin.Context) {
		users := createUsersData()
		id := context.Query("id")
		id2, err := strconv.Atoi(id)
		if err != nil && id != "" {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": "id is not a number",
			})
			log.Println(err)
			return
		}
		for _, u := range users {
			fmt.Println(u.Id)
			if u.Id == id2 {
				context.JSON(200, u)
				return
			}
		}
		context.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
	}
}

func getAllUsers(usuarios []user) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		for _, user := range usuarios {
			c.JSON(200, user)
		}
	}
	return fn
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
