package main

import (
	"awesomeProject/02/first"
	"awesomeProject/02/second"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	path, err := filepath.Abs("02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	firstPartResult := first.First(file)
	fmt.Printf("First part result: %d\n", firstPartResult)

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
		file, err = os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
	}

	secondPartResult := second.Second(file)
	fmt.Printf("Second part result: %d\n", secondPartResult)

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
