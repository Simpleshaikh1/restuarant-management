package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/simpleshaik1/restuarant-management/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

var mmenuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		result, err := mmenuCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured while listing the menu item"})
		}
		var allMenu []bson.M
		if err = result.All(ctx, &allMenu); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allMenu)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateMenus() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateMenus() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
