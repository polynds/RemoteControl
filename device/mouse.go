package device

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

var (
	pressingLeft        = false
	pressingRight       = false
	pressingUp          = false
	pressingDown        = false
	pressingLeftPress   = false
	pressingRightPress  = false
	pressingScrollUp    = false
	pressingScrollDown  = false
	prsssingMiddlePress = false
)

func Press(code string) {
	fmt.Println("Press:" + code)
	switch code {
	case "left":
		if !pressingLeft {
			pressingLeft = true
			go MoveLeft()
			fmt.Println("按下鼠标左移")
		}
	case "right":
		if !pressingRight {
			pressingRight = true
			go MoveRight()
			fmt.Println("按下鼠标右移")
		}
	case "up":
		if !pressingUp {
			pressingUp = true
			go MoveUp()
			fmt.Println("按下鼠标上移")
		}
	case "down":
		if !pressingDown {
			pressingDown = true
			go MoveDown()
			fmt.Println("按下鼠标下移")
		}
	case "left_press":
		if !pressingLeftPress {
			pressingLeftPress = true
			leftPress()
			fmt.Println("按下鼠标左键")
		}
	case "right_press":
		if !pressingRightPress {
			pressingRightPress = true
			rightPress()
			fmt.Println("按下鼠标右键")
		}
	case "double_click":
		doubleClick()
		fmt.Println("按下鼠标右键")
	case "scroll_up":
		if !pressingScrollUp {
			pressingScrollUp = true
			go scrollUp()
			fmt.Println("按下滑轮上滑")
		}
	case "scroll_down":
		if !pressingScrollDown {
			pressingScrollDown = true
			go scrollDown()
			fmt.Println("按下滑轮下滑")
		}
	case "middle_press":
		if !prsssingMiddlePress {
			prsssingMiddlePress = true
			middlePress()
			fmt.Println("按下鼠标中键")
		}
	}
}

func Release(code string) {
	switch code {
	case "left":
		if pressingLeft {
			pressingLeft = false
			fmt.Println("松开鼠标左移")
		}
	case "right":
		if pressingRight {
			pressingRight = false
			fmt.Println("松开鼠标右移")
		}
	case "up":
		if pressingUp {
			pressingUp = false
			fmt.Println("松开鼠标上移")
		}

	case "down":
		if pressingDown {
			pressingDown = false
			fmt.Println("松开鼠标下移")
		}
	case "left_press":
		if pressingLeftPress {
			pressingLeftPress = false
			leftRelease()
			fmt.Println("松开鼠标左键")
		}
	case "right_press":
		if pressingRightPress {
			pressingLeftPress = false
			rightRelease()
			fmt.Println("松开鼠标右键")
		}
	case "scroll_up":
		if pressingScrollUp {
			pressingScrollUp = false
			fmt.Println("松开滑轮上滑")
		}
	case "scroll_down":
		if pressingScrollDown {
			pressingScrollDown = false
			fmt.Println("松开滑轮下滑")
		}
	case "middle_press":
		if prsssingMiddlePress {
			prsssingMiddlePress = false
			middleRelease()
			fmt.Println("松开鼠标中键")
		}
	}
}

func MoveLeft() {
	for pressingLeft {
		time.Sleep(10 * time.Millisecond)
		x, y := robotgo.GetMousePos()
		x -= 2
		robotgo.Move(x, y)
	}
}

func MoveRight() {
	for pressingRight {
		time.Sleep(10 * time.Millisecond)
		x, y := robotgo.GetMousePos()
		x += 2
		robotgo.Move(x, y)

	}
}

func MoveUp() {
	for pressingUp {
		time.Sleep(10 * time.Millisecond)
		x, y := robotgo.GetMousePos()
		y -= 2
		robotgo.Move(x, y)
	}
}

func MoveDown() {
	for pressingDown {
		time.Sleep(10 * time.Millisecond)
		x, y := robotgo.GetMousePos()
		y += 2
		robotgo.Move(x, y)
	}
}

func leftPress() {
	robotgo.Toggle("down", "left")
}

func rightPress() {
	robotgo.Toggle("down", "right")
}

func leftRelease() {
	robotgo.Toggle("up", "left")
}

func rightRelease() {
	robotgo.Toggle("up", "right")
}

func doubleClick() {
	robotgo.Click("left", true)
}

func scrollUp() {
	for pressingScrollUp {
		time.Sleep(100 * time.Millisecond)
		robotgo.Scroll(0, 10)
	}
}

func scrollDown() {
	for pressingScrollDown {
		time.Sleep(100 * time.Millisecond)
		robotgo.Scroll(0, 10)
	}
}

func middlePress() {
	robotgo.Toggle("down", "center")
}

func middleRelease() {
	robotgo.Toggle("up", "center")
}
