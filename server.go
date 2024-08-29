package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/super_secret_app/apis"
	"github.com/super_secret_app/database"
)

func main() {
	r := gin.Default()
	database.Init()
	defer database.End()
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	r.LoadHTMLGlob(dir + "/views/*")
	r.Static("/public/assets", dir+"/public/assets/")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", apis.GenericHandler)
	r.GET("/db", apis.ViewLastUser)
	r.POST("/api/user", apis.InsertUser)
	r.Run() // listen and serve on 0.0.0.0:8080
}
