package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Exercício 1 - Estruturar um JSON
//
//Dependendo do tema escolhido, gere um JSON que atenda as seguintes chaves de
//acordo com o tema.
//
//Os produtos variam por id, nome, cor, preço, estoque, código (alfanumérico),
//publicação (sim-não), data de criação.
//
//Os usuários variam por id, nome, sobrenome, e-mail, idade, altura,
//ativo (sim-não), data de criação.
//
//Transações: id, código da transação (alfanumérico), moeda, valor,
//emissor (string), receptor (string), data da transação.
//
//1. Dentro da pasta go-web crie um arquivo theme.json, o nome tem que ser o tema
//escolhido, ex: products.json.
//
//2. Dentro dele escreva um JSON que permite ter uma matriz de produtos,
//usuários ou transações com todas as suas variantes.

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
	user := user{77, "Victor", "Hugo", "vhbeltra@gmail.com", 19, 1.80, true, "30/05/2022"}

	content, err := json.Marshal(user)
	check(err)

	file, err := os.Create("./data/users.json")
	check(err)
	defer file.Close()

	_, err = file.Write(content)
	check(err)

	fmt.Println("user successful writed")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
