package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-vgo/robotgo"
)

func main() {
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	w, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic(err)
	}
	h, err := strconv.Atoi(os.Args[4])
	if err != nil {
		panic(err)
	}

	black := robotgo.OpenBitmap("black.png")
	defer robotgo.FreeBitmap(black)

	fmt.Println("Press F10 to start...")
	start := robotgo.AddEvent("f10")
	if start {
		fmt.Println("Starting...")
	}

	hook := robotgo.EventStart()
	defer robotgo.EventEnd()

	for {
		select {
		case v := <-hook:
			if v.Rawcode == 120 {
				return
			}
		default:
			bmp := robotgo.CaptureScreen(x, y, w, h)
			defer robotgo.FreeBitmap(bmp)

			mouseX, mouseY := robotgo.FindBitmap(black, bmp)

			robotgo.MoveClick(x+mouseX+80, y+mouseY+80)
			robotgo.MilliSleep(25)
		}
	}
}
