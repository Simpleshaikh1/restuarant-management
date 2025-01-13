package routes

import (
	"fmt"
	"github.com/Simpleshaikh1/golang-jwt/controllers"
	"github.com/gin-gonic/gin"
	controller "github.com/simpleshaik1/restuarant-management/controllers"
	"net/http"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}
