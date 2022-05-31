package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

// Exercício 2 - Olá {nome}
//
//1. Crie dentro da pasta go-web um arquivo chamado main.go
//
//2. Crie um servidor web com Gin que retorne um JSON que tenha uma chave
//“mensagem” e diga Olá seguido do seu nome.
//
//3. Acesse o end-point para verificar se a resposta está correta.

func helloHandler(c *gin.Context) {
	user := os.Getenv("USER")
	str := "Hello " + user
	c.JSON(200, gin.H{"mensagem": str})
}

func main() {
	router := gin.Default()
	router.GET("hello", helloHandler)
	router.Run()
}
