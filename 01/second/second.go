package second

import (
	"bufio"
	"os"
	"strconv"
)

func Second(file *os.File) int {
	scanner := bufio.NewScanner(file)

	var depths [3]int
	fill(scanner, &depths)

	var increaseCount int
	for scanner.Scan() {
		previousDepth := depths[0]
		currentDepth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}

		if previousDepth < currentDepth {
			increaseCount += 1
		}

		moveLeftByOne(&depths)

		depths[2] = currentDepth
	}

	return increaseCount
}

// Fills slice with first three lines from bufio.Scanner
func fill(scanner *bufio.Scanner, s *[3]int) bool {
	var index int
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return false
		}

		s[index] = depth

		index += 1

		if index == 3 {
			break
		}
	}

	return index == 3
}

func moveLeftByOne(lst *[3]int) {
	lst[0] = lst[1]
	lst[1] = lst[2]
}
