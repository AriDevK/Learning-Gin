package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	gin.ForceConsoleColor()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World! :)")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		c.String(http.StatusOK, fmt.Sprintf("Cookie value: %s", cookie))
	})

	r.GET("/form", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.POST("/form", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	r.Run(":8080")
}
