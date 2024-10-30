package main

import (
	"github.com/codescalersinternships/Linktree-RawanMostafa/controllers"
	_ "github.com/codescalersinternships/Linktree-RawanMostafa/docs"
	"github.com/codescalersinternships/Linktree-RawanMostafa/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
//  @title Linktree Service API
//	@version		1.0
//	@description	This is a the server of the linktree application

//	@host		localhost:8080
func main() {

	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	publicRoutes := r.Group("/user")
	{
		publicRoutes.POST("/login", controllers.Login)
		publicRoutes.POST("/register", controllers.Signup)
	}

	protectedRoutes := r.Group("/api/v1/link")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		protectedRoutes.POST("/", controllers.AddLink)
		protectedRoutes.PUT("/:link_id", controllers.EditLink)
		protectedRoutes.DELETE("/:link_id", controllers.DeleteLink)
		protectedRoutes.GET("/:username", controllers.GetUserLinks)
	}

	r.Run(":8080")
}
