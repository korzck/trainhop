package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)


func main() {
	// var App *gin.Engine
	App := gin.Default()
	App.Use(static.Serve("/", static.LocalFile("./app", true)))
	
	App.GET("/", root)	
	App.GET("/admin", admin)
	App.Run()
	
}
