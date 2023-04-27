package api

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)


var(
	App *gin.Engine

)

func admin(c *gin.Context){
	c.String(http.StatusOK,"Hello from admin page of trainhop")
}

// func root(c *gin.Context){
// 	c.String(http.StatusOK,"Hello from root page of trainhop!")
// }

func html(c *gin.Context){
	data, err := ioutil.ReadFile("pages/index.html")
    if err != nil {
      c.Status(http.StatusNotFound)
      return
    }
  
    c.Writer.Header().Set("Content-Type", "text/html")
    c.Writer.Write(data)
}



func init(){
	App = gin.Default()
	App.Use(static.Serve("/", static.LocalFile("./pages", true)))
	App.GET("/admin", admin)
	App.GET("/html", html)
}

func Handler(w http.ResponseWriter , r *http.Request){
	App.ServeHTTP(w,r)
}