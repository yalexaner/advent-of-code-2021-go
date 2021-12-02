package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	depth  int
	length int
}

func (p *Position) move(direction string, steps int) error {
	if direction == "forward" {
		p.length += steps
	} else if direction == "down" {
		p.depth += steps
	} else if direction == "up" {
		p.depth -= steps
	} else {
		return errors.New("unknown direction")
	}

	return nil
}

func (p Position) computePoint() int {
	return p.depth * p.length
}

func main() {
	file, err := os.Open("C:\\Users\\yachmenev.av\\Downloads\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var position Position
	for scanner.Scan() {
		lineData := strings.Split(scanner.Text(), " ")
		direction := lineData[0]
		steps, err := strconv.Atoi(lineData[1])
		if err != nil {
			log.Fatal(err)
		}

		if err = position.move(direction, steps); err != nil {
			log.Fatal(err)
		}
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(position.computePoint())
}
