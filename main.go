package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("start")
	robotgo.MouseClick("left", true)
	robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
	fmt.Println("end")
}
