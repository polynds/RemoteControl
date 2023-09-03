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

	startCapture := make(chan bool)

	hub := ws.NewHub()
	go hub.Run()
	r.GET("/ws", func(c *gin.Context) {
		ws.WsClient(hub, c.Writer, c.Request, startCapture)
	})

	r.GET("/stream", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "image/png")
		// 设置刷新响应的间隔时间
		flusher, ok := c.Writer.(http.Flusher)
		if !ok {
			http.Error(c.Writer, "Streaming not supported", http.StatusInternalServerError)
			return
		}
		flusher.Flush()

		go device.PushFlow(c.Writer, startCapture)
	})

	//sp := device.NewScreenshot(startCapture)
	//go sp.StartCapture()

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
