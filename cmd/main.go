package main

import (
	// "os"

	"github.com/codescalersinternships/Linktree-RawanMostafa/controllers"
	"github.com/codescalersinternships/Linktree-RawanMostafa/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// port := os.Getenv("PORT")

	r := gin.Default()

	publicRoutes := r.Group("/public")
	{
		publicRoutes.POST("/login", controllers.Login)
		publicRoutes.POST("/register", controllers.Signup)
	}

	protectedRoutes := r.Group("/protected")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		// Protected routes
	}


	r.Run(":8080")
}
