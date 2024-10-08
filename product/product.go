package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Product struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var Products []Product

func GetProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Products)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	for _, product := range Products {
		if product.Id == id {
			c.IndentedJSON(http.StatusOK, product)
			return
		} else {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		}
	}
}

func SetProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBind(&product); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	Products = append(Products, product)
	status := "New product has been successfully created"
	c.IndentedJSON(http.StatusOK, status)
}
