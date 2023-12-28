package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	r := gin.Default()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World! :)")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
