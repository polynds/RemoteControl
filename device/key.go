package device

import (
	"github.com/go-vgo/robotgo"
	"time"
)

func PressKeyUp() {
	timer := time.After(100 * time.Millisecond)
	robotgo.KeyTap("up")
	<-timer
	robotgo.KeyTap("up", "up")
}

func PressKeyDown() {
	timer := time.After(100 * time.Millisecond)
	robotgo.KeyTap("down")
	<-timer
	robotgo.KeyTap("down", "down")
}

func PressKeyLeft() {
	timer := time.After(100 * time.Millisecond)
	robotgo.KeyTap("left")
	<-timer
	robotgo.KeyTap("left", "left")
}

func PressKeyRight() {
	timer := time.After(100 * time.Millisecond)
	robotgo.KeyTap("right")
	<-timer
	robotgo.KeyTap("right", "right")
}

func ClickLeft() {
	robotgo.Click()
}
func ClickRight() {
	robotgo.Click("right")
}
