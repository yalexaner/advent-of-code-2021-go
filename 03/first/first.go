package first

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

type Counter struct {
	zeros int
	ones  int
}

func (c *Counter) increase(digit string) error {
	if digit == "0" {
		c.zeros += 1
	} else if digit == "1" {
		c.ones += 1
	} else {
		return errors.New("only allowed 0 or 1")
	}

	return nil
}

func (c Counter) getMax() string {
	if c.zeros < c.ones {
		return "1"
	} else {
		return "0"
	}
}

func (c Counter) getMin() string {
	if c.ones < c.zeros {
		return "1"
	} else {
		return "0"
	}
}

type Sequence [12]Counter

func (s *Sequence) increase(digits string) error {
	for index, digit := range digits {
		if err := s[index].increase(string(digit)); err != nil {
			return err
		}
	}

	return nil
}

func (s Sequence) toString(f func(counter Counter) string) (result string) {
	for _, counter := range s {
		result += f(counter)
	}
	return
}

func First(file *os.File) int64 {
	scanner := bufio.NewScanner(file)
	counters, err := createSequence(scanner)
	if err != nil {
		log.Fatal(err)
	}

	gamma, err := calculateGamma(counters)
	if err != nil {
		log.Fatal(err)
	}

	epsilon, err := calculateEpsilon(counters)
	if err != nil {
		log.Fatal(err)
	}

	return gamma * epsilon
}

func createSequence(scanner *bufio.Scanner) (counters Sequence, err error) {
	for scanner.Scan() {
		digits := scanner.Text()
		if err = counters.increase(digits); err != nil {
			return
		}
	}

	return
}

func calculateGamma(counters Sequence) (gamma int64, err error) {
	binaryNumber := counters.toString(func(counter Counter) string {
		return counter.getMax()
	})

	gamma, err = strconv.ParseInt(binaryNumber, 2, 32)

	return
}

func calculateEpsilon(counters Sequence) (epsilon int64, err error) {
	binaryNumber := counters.toString(func(counter Counter) string {
		return counter.getMin()
	})

	epsilon, err = strconv.ParseInt(binaryNumber, 2, 32)

	return
}
