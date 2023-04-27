package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)


func main() {
	// var App *gin.Engine

	// cmd := exec.Command("ls", "-l")
    // // cmd.Stdout = os.Stdout
    // // cmd.Stderr = os.Stderr
    // err := cmd.Run()
	// recover()
	// fmt.Println(err.Error())
	// fmt.Println("hello")
	App := gin.Default()
	App.Use(static.Serve("/", static.LocalFile("./frontend", true)))
	
	App.GET("/", root)	
	App.POST("/upload", getFile)
	App.GET("/download", giveFile)
	App.GET("/admin", admin)
	App.Run()
	
}
