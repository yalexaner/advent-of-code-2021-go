package main

import (
	"awesomeProject/01/first"
	"awesomeProject/01/second"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	path, err := filepath.Abs("01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	file := openFile(path)

	firstPartResult := first.First(file)
	fmt.Printf("First part result: %d\n", firstPartResult)

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}

		file = openFile(path)
	}

	secondPartResult := second.Second(file)
	fmt.Printf("Second part result: %d\n", secondPartResult)

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
