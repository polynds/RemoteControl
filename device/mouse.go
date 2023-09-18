package device

import "github.com/go-vgo/robotgo"

func TouchMove(x int, y int) {

	robotgo.MoveRelative(x, y)

	//fmt.Println("MoveLocation:1:", x, y)
	//currentX, currentY := robotgo.Location()
	//fmt.Println("MoveLocation:2:", currentX, currentY)
	//offsetX, offsetY := x-currentX, y-currentY
	//fmt.Println("MoveLocation:3:", offsetX, offsetY)
	//fmt.Println("MoveLocation:4:", currentX+offsetX, currentY+offsetY)
	//robotgo.Move(currentX+offsetX, currentY+offsetY)
}
