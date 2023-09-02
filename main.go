package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/polynds/RemoteControl/device"
	"github.com/polynds/RemoteControl/ip"
	"github.com/polynds/RemoteControl/ws"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	hub := ws.NewHub()
	go hub.Run()
	r.GET("/ws", func(c *gin.Context) {
		ws.WsClient(hub, c.Writer, c.Request)
	})
	r.GET("/press/:code", func(c *gin.Context) {
		code := c.Param("code")
		device.Press(code)
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.GET("/release/:code", func(c *gin.Context) {
		code := c.Param("code")
		device.Release(code)
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	port := ":9021"
	fmt.Println("http://127.0.0.1" + port)
	if ip, err := ip.ClientIp(); err == nil {
		fmt.Println("http://" + ip + port)
	}
	err := r.Run(port)
	if err != nil {
		return
	}
}
