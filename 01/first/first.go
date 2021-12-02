package first

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func First(file *os.File) int {
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

	return increaseCount
}
