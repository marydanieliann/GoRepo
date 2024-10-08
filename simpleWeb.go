package main

import (
	"github.com/gin-gonic/gin"
	"leetcodeTasks/entity"
	"leetcodeTasks/product"
)

func main() {
	product.Products = []product.Product{
		{Id: "1", Name: "milk", Price: 3.99},
		{Id: "2", Name: "chocolate", Price: 2.5},
		{Id: "3", Name: "bishop", Price: 5.99},
	}
	entity.Humans = []entity.Human{
		{Id: "1", Name: "ann", Surname: "ann2", Age: 30},
		{Id: "2", Name: "foo", Surname: "foo2", Age: 20},
		{Id: "3", Name: "john", Surname: "doe", Age: 15},
		{Id: "4", Name: "Human4", Surname: "Human4", Age: 3},
	}

	router := gin.Default()
	router.GET("/human", entity.GetHuman)
	router.POST("/human", entity.SetHuman)
	router.POST("/getById/:id", entity.GetHumanByID)
	router.GET("/products", product.GetProducts)
	router.POST("/products", product.SetProduct)
	router.POST("/getProductById/:id", product.GetProductByID)
	router.Run(":8080")

}
