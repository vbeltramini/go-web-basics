package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
)

type user struct {
	Id           int     `json:"id"`
	Name         string  `json:"name" validate:"required"`
	LastName     string  `json:"last_name" validate:"required"`
	Email        string  `json:"email" validate:"required"`
	Age          int     `json:"age" validate:"required"`
	Height       float32 `json:"height" validate:"required"`
	Active       bool    `json:"active" validate:"required"`
	CreationDate string  `json:"creation_date"`
}

var usersData []user

func main() {
	router := gin.Default()
	group := router.Group("users")
	{
		group.GET("/", listUsers(&usersData))
		group.POST("/", saveUser(&usersData))
	}
	router.Run()
}

func saveUser(users *[]user) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var u user

		validate := validator.New()

		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := validate.Struct(u)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				fmt.Println(err)
				return
			}
			for _, err := range err.(validator.ValidationErrors) {
				if err != nil {
					s := fmt.Sprintf("Field: %s cannot be empty", err.StructField())
					log.Println(s)
					c.JSON(http.StatusBadRequest, gin.H{"error": s})
					return
				}
			}
		}

		u.Id = getNewId()
		u.CreationDate = time.Now().Format("2006-01-02")
		*users = append(*users, u)
		c.JSON(http.StatusOK, gin.H{"content": u})
	}
	return fn
}

func getNewId() int {
	var newID int
	if len(usersData) <= 0 {
		newID = 1
	} else {
		newID = usersData[len(usersData)-1].Id + 1
	}
	return newID
}

func listUsers(users *[]user) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"content": &users})
	}
	return fn
}
