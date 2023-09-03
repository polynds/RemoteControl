package device

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Screenshot struct {
	LasData []byte
	Ticker  *time.Ticker
	Start   chan bool
	status  bool
	Done    chan struct{}
}

func NewScreenshot(startCapture chan bool) *Screenshot {
	return &Screenshot{
		Start:  startCapture,
		status: false,
		Done:   make(chan struct{}),
		Ticker: time.NewTicker(time.Second),
	}
}

func (s *Screenshot) StartCapture() {
	defer func() {
		s.Ticker.Stop()
	}()

	for {
		select {
		case <-s.Done:
			return
		case start := <-s.Start:
			s.status = start
			s.Ticker.Reset(time.Second)
		case <-s.Ticker.C:
			fmt.Println("capture", s.status)
			if s.status {
				go s.capture()
			} else {
				s.Ticker.Stop()
			}
		}
	}
}

func (s *Screenshot) capture() {
	bitmap := robotgo.CaptureScreen()
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	path := filepath.Join(currentDir, "screenshot.png")
	robotgo.SaveBitmap(bitmap, path)
}

func PushFlow(w http.ResponseWriter, stop chan bool) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	path := filepath.Join(currentDir, "screenshot.png")
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening image file:", err)
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	for {
		_, err = io.Copy(w, file)
		if err != nil {
			fmt.Println("Error writing image data to response writer:", err)
			return
		}

		select {
		case <-stop:
			return
		}
	}
}
