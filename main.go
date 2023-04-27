package main

import (
	"io/ioutil"
	"net/http"
	"trainhop/api"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Html(c *gin.Context){
	data, err := ioutil.ReadFile("api/pages/index.html")
    if err != nil {
      c.Status(http.StatusNotFound)
      return
    }

    // Set content-type header and return HTML file
    c.Writer.Header().Set("Content-Type", "text/html")
    c.Writer.Write(data)
}

func main() {
	
	// Setup homepage route
	api.App.GET("/", Html)
	// api.App.Run()
	// api.Init()
	api.App.Use(static.Serve("/", static.LocalFile("./api/pages", true)))
	api.App.Run()

	

}