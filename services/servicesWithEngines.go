package services

import "github.com/gin-gonic/gin"

func UserServiceRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/human", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from user Service",
		})
	})
	return router
}

func ProductServiceRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/product", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from product Service",
		})
	})
	return router
}
