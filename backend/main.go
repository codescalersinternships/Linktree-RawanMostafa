package main

import (
	"github.com/codescalersinternships/Linktree-RawanMostafa/controllers"
	_ "github.com/codescalersinternships/Linktree-RawanMostafa/docs"
	"github.com/codescalersinternships/Linktree-RawanMostafa/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-contrib/cors"

)
//  @title Linktree Service API
//	@version		1.0
//	@description	This is a the server of the linktree application

//	@host		localhost:8083
func main() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"POST", "GET", "OPTIONS","PUT","DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	publicRoutes := r.Group("/user")
	{
		publicRoutes.POST("/login", controllers.Login)
		publicRoutes.POST("/register", controllers.Signup)
		publicRoutes.GET("/:username", controllers.GetUserInfo)
	}
	r.Use(middleware.AuthenticationMiddleware()).PUT("/user/update-bio/:username", controllers.UpdateBio)

	r.GET("/api/v1/link/:username", controllers.GetUserLinks)

	protectedRoutes := r.Group("/api/v1/link")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		protectedRoutes.POST("/", controllers.AddLink)
		protectedRoutes.PUT("/:link_id", controllers.EditLink)
		protectedRoutes.DELETE("/:link_id", controllers.DeleteLink)
	}

	r.Run(":8083")
}
