package main

import "github.com/gin-gonic/gin"

func HelloHandler(c *gin.Context) {
	c.String(200, "Hello, world!")
}

func main() {
	r := gin.Default()

	r.GET("/", HelloHandler)

	r.Run(":3000")
}
