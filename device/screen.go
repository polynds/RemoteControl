package device

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func Capture(send chan []byte, capture chan bool) {
	// 开始录制屏幕
	send <- []byte("capturing...")
	capture <- true
}

func CloseCapture(send chan []byte, capture chan bool) {
	send <- []byte("Capture closing...")
	capture <- false
}

func Size() (int, int) {
	return robotgo.GetScreenSize()
}

type ScreenSize struct {
	Width  int
	Height int
}

func NewScreenSize() *ScreenSize {
	w, h := Size()
	return &ScreenSize{
		Width:  w,
		Height: h,
	}
}

type Region struct {
	Width  int
	Height int
}

func NewClientSize(w int, h int) *Region {
	return &Region{
		Width:  w,
		Height: h,
	}
}

type Coordinates struct {
	X int
	Y int
}

func MapCoordinates(region *Region, screenSize *ScreenSize, coordinates Coordinates) Coordinates {
	xFactor := float64(screenSize.Width) / float64(region.Width)
	yFactor := float64(screenSize.Height) / float64(region.Height)
	fmt.Println(screenSize, region, xFactor, yFactor)
	mappedX := int(float64(coordinates.X) * xFactor)
	mappedY := int(float64(coordinates.Y) * yFactor)

	return Coordinates{
		X: mappedX,
		Y: mappedY,
	}
}
