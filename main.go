package main

import (
	"github.com/gin-gonic/gin"
	database "github.com/simpleshaik1/restuarant-management/database"
	middleware "github.com/simpleshaik1/restuarant-management/middleware"
	routes "github.com/simpleshaik1/restuarant-management/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	routes.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access Granted for api 1"})
	})

	err := router.Run(":" + port)
	if err != nil {
		log.Panic(err)
	}
}
