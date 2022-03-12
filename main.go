package main

import (
	"github.com/gin-gonic/gin"

	"jwt-gin/controllers"
	"jwt-gin/middlewares"
	"jwt-gin/models"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")
	public.POST("/login", controllers.Login)

	// public.POST("/register", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})
	// })

	public.POST("/register", controllers.Register)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")

}
