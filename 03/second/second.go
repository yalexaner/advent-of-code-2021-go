package second

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const BitCount = 12

type Counter struct {
	zeros int
	ones  int
}

func (c *Counter) increase(digit rune) {
	if digit == '0' {
		c.zeros += 1
	} else if digit == '1' {
		c.ones += 1
	}
}

func (c Counter) getMax() rune {
	if c.zeros > c.ones {
		return '0'
	} else {
		return '1'
	}
}

func (c Counter) getMin() rune {
	if c.ones < c.zeros {
		return '1'
	} else {
		return '0'
	}
}

func Second(file *os.File) int64 {
	scanner := bufio.NewScanner(file)
	lines := readLines(scanner)

	generator, err := calculateNumber(lines, func(counter Counter) rune {
		return counter.getMax()
	})
	if err != nil {
		log.Fatal(err)
	}

	scrubber, err := calculateNumber(lines, func(counter Counter) rune {
		return counter.getMin()
	})
	if err != nil {
		log.Fatal(err)
	}

	return generator * scrubber
}

func readLines(scanner *bufio.Scanner) (lines []string) {
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func calculateNumber(lines []string, f func(Counter) rune) (gamma int64, err error) {
	for i := 0; i < BitCount; i++ {
		if len(lines) == 1 {
			break
		}

		counter := computeCounter(lines, i)

		lines = filter(lines, func(line string) bool {
			return rune(line[i]) == f(counter)
		})
	}

	binaryNumber := lines[0]

	gamma, err = strconv.ParseInt(binaryNumber, 2, 32)
	return
}

func computeCounter(lines []string, index int) (counter Counter) {
	for _, line := range lines {
		counter.increase(rune(line[index]))
	}
	return
}

func filter(lines []string, f func(string) bool) (result []string) {
	for _, line := range lines {
		if f(line) {
			result = append(result, line)
		}
	}
	return
}
