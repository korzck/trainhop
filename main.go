package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)


func main() {
	App := gin.Default()
	App.Use(static.Serve("/", static.LocalFile("./frontend", true)))
	
	App.GET("/", root)	
	App.POST("/upload", getFile)
	App.GET("/admin", admin)
	App.Run()
}