package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("C:\\Users\\yachmenev.av\\Downloads\\input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var (
		increaseCount int
		previousDepth int
	)
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if previousDepth < depth {
			increaseCount += 1
		}

		previousDepth = depth
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(increaseCount)
}
