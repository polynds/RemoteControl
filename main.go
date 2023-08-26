package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/press/:code", func(c *gin.Context) {
		code := c.Param("code")
		Press(code)
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.GET("/release/:code", func(c *gin.Context) {
		code := c.Param("code")
		Release(code)
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	port := ":9021"
	fmt.Println("http://127.0.0.1" + port)
	err := r.Run(port)
	if err != nil {
		return
	}
}
