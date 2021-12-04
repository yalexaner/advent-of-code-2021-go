package main

import (
	"awesomeProject/02/first"
	"awesomeProject/02/second"
	"awesomeProject/core"
	"fmt"
	"io"
)

func main() {
	path := core.AbsPath("02/input.txt")
	file := core.OpenFile(path)

	firstPartResult := first.First(file)
	fmt.Printf("First part result: %d\n", firstPartResult)

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		core.CloseFile(file)
		file = core.OpenFile(path)
	}

	secondPartResult := second.Second(file)
	fmt.Printf("Second part result: %d\n", secondPartResult)

	core.CloseFile(file)
}
