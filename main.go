package main

import (
	"Learning-Gin/models/forms"
	"Learning-Gin/models/uri"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
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

	r.GET("/hello/:name/:times", func(c *gin.Context) {
		var data uri.HelloUri
		err := c.ShouldBindUri(&data)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Bad request",
			})
			return
		}

		msg := "Hello " + data.Name + "! "
		regards := strings.Repeat(msg, data.Times)
		c.String(http.StatusOK, regards)
	})

	r.GET("/params/:param/*optional", func(c *gin.Context) {
		param := c.Param("param")
		optional := c.Param("optional")
		c.String(http.StatusOK, fmt.Sprintf("Param: %s\nOptional: %s", param, optional))
	})

	r.GET("/query", func(c *gin.Context) {
		id := c.Query("id")
		limit := c.DefaultQuery("limit", "10")
		offset := c.DefaultQuery("offset", "0")
		c.String(http.StatusOK, fmt.Sprintf("id: %s; limit: %s; offset: %s", id, limit, offset))
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
		var data forms.MessageForm
		err := c.Bind(&data)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Bad request",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": data.Message,
			"nick":    data.Nick,
		})
	})

	r.GET("/google", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://google.com")
	})

	r.Run(":8080")
}
