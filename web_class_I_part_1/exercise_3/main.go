package main

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

// Exercício 3 - Listar Entidade
//
//Já tendo criado e testado nossa API que nos recebe, geramos uma rota que
//retorna uma lista do tema escolhido.
//
//1. Dentro do "main.go", crie uma estrutura de acordo com o tema com os campos
//correspondentes.
//
//2. Crie um endpoint cujo caminho é /thematic (plural). Exemplo: “/products”
//
//3. Crie uma handler para o endpoint chamada "GetAll".
//
//4. Crie um slice da estrutura e retorne-a por meio de nosso endpoint.

type user struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	LastName     string  `json:"lastname"`
	Email        string  `json:"email"`
	Age          int     `json:"age"`
	Height       float32 `json:"height"`
	Active       bool    `json:"active"`
	CreationDate string  `json:"creationdate"`
}

func main() {
	router := gin.Default()
	router.GET("users", getAll)
	router.Run()
}

func getAll(c *gin.Context) {

	jsonData := []byte(`
    [
        {"id":77,"name":"Victor","lastname":"Hugo","email":"victor@gmail.com","age":19,"height":1.8,"active":true,"creationdate":"30/05/2022"},
        {"id":77,"name":"Victor","lastname":"Hugo","email":"victor@gmail.com","age":19,"height":1.8,"active":true,"creationdate":"30/05/2022"},
        {"id":77,"name":"Victor","lastname":"Hugo","email":"victor@gmail.com","age":19,"height":1.8,"active":true,"creationdate":"30/05/2022"},
        {"id":77,"name":"Victor","lastname":"Hugo","email":"victor@gmail.com","age":19,"height":1.8,"active":true,"creationdate":"30/05/2022"}
    ]`)

	var u []user

	if err := json.Unmarshal(jsonData, &u); err != nil {
		log.Fatal(err)
	}
	for _, i := range u {
		c.JSON(200, i)
	}
}
