package main

import (
	"github.com/codescalersinternships/Linktree-RawanMostafa/controllers"
	"github.com/codescalersinternships/Linktree-RawanMostafa/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	publicRoutes := r.Group("/public")
	{
		publicRoutes.POST("/login", controllers.Login)
		publicRoutes.POST("/register", controllers.Signup)
	}

	protectedRoutes := r.Group("/protected")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		protectedRoutes.POST("/addlink", controllers.AddLink)
		protectedRoutes.PUT("/editlink/:link_id", controllers.EditLink)
	}

	r.Run(":8080")
}
