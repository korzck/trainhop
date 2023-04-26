package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


var(
	App *gin.Engine

)

func f(c *gin.Context){
	c.String(http.StatusOK,"Hello from golang in vercel")
}

func myRoute(r *gin.RouterGroup){
	r.GET("/admin", f)
}

func init(){
	App = gin.New()
	r := App.Group("/api")
	myRoute(r)

}

func Handler(w http.ResponseWriter , r *http.Request){
	App.ServeHTTP(w,r)
}