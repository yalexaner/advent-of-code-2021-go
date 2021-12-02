package second

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	depth  int
	length int
	aim    int
}

func (p *Position) move(direction string, steps int) error {
	if direction == "forward" {
		p.length += steps
		p.depth += p.aim * steps
	} else if direction == "down" {
		p.aim += steps
	} else if direction == "up" {
		p.aim -= steps
	} else {
		return errors.New("unknown direction")
	}

	return nil
}

func (p Position) computePoint() int {
	return p.depth * p.length
}

func Second(file *os.File) int {
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

	return position.computePoint()
}
