package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


var(
	App *gin.Engine

)

func admin(c *gin.Context){
	c.String(http.StatusOK,"Hello from admin page of trainhop")
}

func root(c *gin.Context){
	c.String(http.StatusOK,"Hello from root page of trainhop!")
}

// func myRoute(r *gin.RouterGroup){
// 	r.GET("/admin", f)
// }

func init(){
	App = gin.New()
	// r := App.Group("/api")
	// myRoute(r)
	App.GET("/admin", admin)
	App.GET("/", root)
}

func Handler(w http.ResponseWriter , r *http.Request){
	App.ServeHTTP(w,r)
}