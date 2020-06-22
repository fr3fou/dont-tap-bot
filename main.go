package main

import (
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
	robotgo.CaptureScreen(x, y, w, h)
}
