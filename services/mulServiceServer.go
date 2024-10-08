package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	userRouter := gin.Default()
	userRouter.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from User Service"})
	})

	userServer := &http.Server{
		Addr:         ":8080",
		Handler:      userRouter,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	productRouter := gin.Default()
	productRouter.GET("/product", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from Product Service"})
	})

	productServer := &http.Server{
		Addr:         ":8081",
		Handler:      productRouter,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := userServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	go func() {
		if err := productServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	select {}
}
