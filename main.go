package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"o-rest/config"
	handler2 "o-rest/handler"
	"o-rest/repository"
	service2 "o-rest/service"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	fmt.Println("Running...")

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUserName := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	db := config.DBInit(dbUserName,dbPassword,dbHost,dbPort,dbName)
	itemRepo := repository.NewItemRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	service := service2.NewService(orderRepo, itemRepo)
	handler := handler2.NewHandlerService(service)

	router := gin.Default()
	router.GET("/orders", handler.GetOrder)
	router.POST("/order", handler.CreateOrder)
	router.PUT("/order",handler.UpdateOrder)
	router.DELETE("/order/:id",handler.DeleteOrder)

	router.Run()

}
