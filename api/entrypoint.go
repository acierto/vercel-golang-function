package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
	"log"
	"net/http"
	"os"
)

var (
	app *gin.Engine
)

func myRoute(r *gin.RouterGroup) {
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from GoLang in Vercel")
	})
	r.GET("/employees", func(c *gin.Context) {
		fmt.Println(os.Environ())

		fmt.Printf("SUPABASE_URL %v", os.Getenv("SUPABASE_URL"))
		fmt.Printf("SUPABASE_ANON_KEY %v", os.Getenv("SUPABASE_ANON_KEY"))

		client, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_ANON_KEY"), &supabase.ClientOptions{})
		if err != nil {
			fmt.Println("cannot initialize client", err)
		}
		data, _, err := client.From("employees").Select("*", "exact", false).Execute()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprint(err))
		}
		if data != nil {
			c.String(http.StatusOK, string(data))
		}
	})
}

func init() {
	if os.Getenv("__VERCEL_DEV_RUNNING") == "1" {
		err := godotenv.Load("../../../.env.development.local")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	app = gin.New()
	r := app.Group("/api")
	myRoute(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
