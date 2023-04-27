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