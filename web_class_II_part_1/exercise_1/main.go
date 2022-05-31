package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u.Id = getNewId()
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
