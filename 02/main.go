package main

import (
	"awesomeProject/02/first"
	"fmt"
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

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
