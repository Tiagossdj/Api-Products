package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)

	// camada usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// camada de controllers
	productController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)

	server.Run(":8000")
}
