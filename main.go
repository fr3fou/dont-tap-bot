package main

import (
	"fmt"
	"os"
)

func main() {
	coords := os.Args[1:]
	fmt.Println(coords)
}
