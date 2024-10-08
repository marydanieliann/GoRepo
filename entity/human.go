package entity

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Human struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

var Humans []Human

func GetHuman(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Humans)
}

func SetHuman(c *gin.Context) {
	var newHuman Human
	if err := c.ShouldBind(&newHuman); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	Humans = append(Humans, newHuman)
	status := "New human has been successfully created"
	c.IndentedJSON(http.StatusOK, status)
}
func GetHumanByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range Humans {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		} else {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Human not found"})
		}
	}
}
