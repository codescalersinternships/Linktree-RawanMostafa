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

	protectedRoutes := r.Group("/links")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		protectedRoutes.POST("/add", controllers.AddLink)
		protectedRoutes.PUT("/edit/:link_id", controllers.EditLink)
		protectedRoutes.DELETE("/delete/:link_id", controllers.DeleteLink)

	}

	r.Run(":8080")
}
