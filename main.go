package main

import (
	"github.com/Timotius-Nugroho/learn-go-rest-api.git/controllers/categorycontroller"
	"github.com/Timotius-Nugroho/learn-go-rest-api.git/controllers/productcontroller"
	"github.com/Timotius-Nugroho/learn-go-rest-api.git/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.GetAll)
	r.GET("/api/product/:id", productcontroller.GetById)
	r.POST("/api/product", productcontroller.Create)
	r.PATCH("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product/:id", productcontroller.Delete)

	r.POST("/api/category", categorycontroller.Create)

	r.Run()
}
