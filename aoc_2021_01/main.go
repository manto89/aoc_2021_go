package aoc_2021_01

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var fileName = "./input/01"

func readLines() {
	f, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("File %s doesn't exist", fileName)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("Unable to close file")
		}
	}(f)

	scanner := bufio.NewScanner(f)
	var lastInt int
	for {
		scanner.Scan()
		firstString := scanner.Text()
		if len(firstString) > 0 {
			lastInt, err = strconv.Atoi(firstString)
			if err != nil {
				continue
			}
			break
		}
	}
	increasingCounter := 0
	decreasingCounter := 0
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) < 1 {
			continue
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		if i > lastInt {
			increasingCounter++
		} else {
			decreasingCounter++
		}
		lastInt = i

	}
	log.Printf("Increasing %d, Decreasing %d", increasingCounter, decreasingCounter)
}

func Solve() {
	log.Printf("**** DAY 01 ****")
	log.Printf("Given a list of number, count how many numbers are bigger than the previous one")
	readLines()
	log.Printf("****")
}
