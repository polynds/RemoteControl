package device

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func PressKeyUp() {
	err := robotgo.KeyPress("up")
	if err != nil {
		fmt.Println("KeyPress err:", err)
		return
	}

	//timer := time.After(100 * time.Millisecond)
	//err := robotgo.KeyPress("up")
	//if err != nil {
	//	return
	//}
	//<-timer
	//err := robotgo.KeyTap("up", "up")
	//if err != nil {
	//	return
	//}
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
