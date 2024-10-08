package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "host=localhost user=user dbname=mydb port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	if err := DB.AutoMigrate(&User{}); err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}
}

func AddUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func main() {
	InitDB()
	r := gin.Default()
	r.POST("/users", AddUser)
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Failed to run server: %v", err)
	}
}
