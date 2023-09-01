package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/polynds/RemoteControl/device"
	"github.com/polynds/RemoteControl/ip"
	"github.com/polynds/RemoteControl/qrcode"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", nil)
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
	qrcodeUrl := "http://127.0.0.1" + port
	fmt.Println("http://127.0.0.1" + port)
	if ip, err := ip.ClientIp(); err == nil {
		fmt.Println("http://" + ip + port)
		qrcodeUrl = "http://" + ip + port
	}
	if len(qrcodeUrl) > 0 {
		qr := qrcode.NewQRCode2ConsoleWithUrl(qrcodeUrl, true)
		qr.Output()
	}
	err := r.Run(port)
	if err != nil {
		return
	}
}
