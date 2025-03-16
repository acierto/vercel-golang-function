package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	app *gin.Engine
)

func myRoute(r *gin.RouterGroup) {
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from GoLang in Vercel")
	})
}

func init() {
	app = gin.New()
	r := app.Group("/api")
	myRoute(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
