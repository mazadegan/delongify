package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.String(200, "Hello, world!")
}

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()

	r.GET("/hello", HelloHandler)

	api := r.Group("/api")
	api.GET("/ping", PingHandler)

	r.Use(static.Serve("/", static.LocalFile("./delongify_frontend/dist", true)))

	r.Run()
}
