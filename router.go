package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func admin(c *gin.Context){
	c.String(http.StatusOK,"Hello from admin page of trainhop")
}

func root(c *gin.Context){
	data, err := ioutil.ReadFile("app/index.html")
    if err != nil {
      c.Status(http.StatusNotFound)
      return
    }
    c.Writer.Header().Set("Content-Type", "text/html")
    c.Writer.Write(data)
}

func getFile(c *gin.Context){
	file, err := c.FormFile("file")
    if err != nil {
        c.JSON(409, gin.H {
			"error" : "couldn't get users file",
		})
    }

	err = processFile(c, file)
	if err != nil {
        c.JSON(409, gin.H {
			"error" : err.Error(),
		})
    } else {
		c.JSON(200, gin.H {
			"success" : "file saved",
		})
	}
	return

    
}