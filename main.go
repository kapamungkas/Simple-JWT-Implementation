package main

import (
	"fmt"
	"simple-jwt-golang/handlers"
	"simple-jwt-golang/middlewares"
	"simple-jwt-golang/repositories"
	"simple-jwt-golang/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	auth_repo := repositories.NewAuthRepository()
	auth_service := services.NewAuthService(auth_repo)
	auth_handler := handlers.NewAuthHandler(auth_service)
	r.POST("/login", auth_handler.AuthLogin)

	r.GET("/admin", middlewares.UserMiddleware(), func(c *gin.Context) {
		username, _ := c.Get("username")
		text := fmt.Sprintf("%v", username)
		c.JSON(200, gin.H{
			"message": "Hello " + text,
		})
	})
	r.Run()
}
