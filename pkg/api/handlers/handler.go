package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mostafa20220/ORGMANAGEMENTAPP/pkg/api/routes"
	"github.com/mostafa20220/ORGMANAGEMENTAPP/pkg/controllers"
)


func Handler(g * gin.Engine){

	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})


	g.POST("/signup", controllers.Signup)

	g.POST("/signin", controllers.Signin)

	g.POST("/logout", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	g.POST("/invite", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	g.POST("/organizations", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	g.GET("/organizations/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	g.PUT("/organizations/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	g.DELETE("/organizations/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

}