package main

import (
	"github.com/gin-gonic/gin"

	"jwt-gin/controllers"
)

func main() {

	r := gin.Default()

	public := r.Group("/api")

	// public.POST("/register", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})
	// })

	public.POST("/register", controllers.Register)

	r.Run(":8080")

}
