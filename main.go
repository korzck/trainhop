package main

import (
	"trainhop/api"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	api.App.Use(static.Serve("/", static.LocalFile("./testing", true)))
	
	// Setup homepage route
	api.App.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api.App.Run()

	

}