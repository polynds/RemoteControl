package device

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func TouchMove(x int, y int) {
	fmt.Println("MoveLocation:1:", x, y)
	currentX, currentY := robotgo.Location()
	fmt.Println("MoveLocation:2:", currentX, currentY)
	robotgo.Move(x, y)
	currentX, currentY = robotgo.Location()
	fmt.Println("MoveLocation:3:", currentX, currentY)
}
