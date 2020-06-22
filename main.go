package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-vgo/robotgo"
)

// Square is the square bitmap
const Square = "b50,50,eNrtwTEBAAAAwqD1T+1vBqAAAAAAAAAAeAMdTAAB"

// F keys
const (
	F9  = 120
	F10 = 121
)

// Events
const (
	KeyUp = 5
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

	fmt.Println("Press F10 to start / stop!")
	fmt.Println("Press F9 to quit!")

	// Start async event listener
	hook := robotgo.EventStart()
	defer robotgo.EventEnd()

	isRunning := false

	qc := make(chan struct{})

	for v := range hook {
		if v.Kind != KeyUp {
			continue
		}

		switch v.Rawcode {
		case F9:
			fmt.Println("Quitting!")
			qc <- struct{}{}
			return
		case F10:
			if isRunning {
				isRunning = false
				fmt.Println("Stopping!")
				qc <- struct{}{}
			} else {
				isRunning = true
				fmt.Println("Starting!")
				go play(qc, x, y, w, h)
			}
		}
	}
}

func play(qc chan struct{}, x, y, w, h int) {
	square := robotgo.BitmapFromStr(Square)
	defer robotgo.FreeBitmap(square)

	for {
		select {
		case <-qc:
			return
		default:
			bmp := robotgo.CaptureScreen(x, y, w, h)
			defer robotgo.FreeBitmap(bmp)

			mouseX, mouseY := robotgo.FindBitmap(square, bmp)
			if mouseX == -1 || mouseY == -1 {
				continue
			}

			robotgo.MoveClick(x+mouseX+80, y+mouseY+80)
			robotgo.MilliSleep(28)
		}
	}
}
