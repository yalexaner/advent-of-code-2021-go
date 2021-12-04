package main

import (
	"awesomeProject/03/first"
	"awesomeProject/core"
	"fmt"
)

func main() {
	path := core.AbsPath("03/input.txt")
	file := core.OpenFile(path)

	firstPartResult := first.First(file)
	fmt.Printf("First part result: %d\n", firstPartResult)
}
