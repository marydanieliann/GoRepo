package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/secure_json", func(c *gin.Context) {
		names := []string{"Alice", "Bob", "Charlie"}
		c.SecureJSON(200, names)
		fmt.Println(names)
	})

	router.Run(":8080")
}
