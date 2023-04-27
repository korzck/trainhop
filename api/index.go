package api

import (
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
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func init(){
	App = gin.Default()
	App.Use(static.Serve("/", static.LocalFile("./pages", true)))
	App.GET("/admin", admin)
	// App.GET("/", root)
	App.GET("/html", html)
}

func Handler(w http.ResponseWriter , r *http.Request){
	App.ServeHTTP(w,r)
}