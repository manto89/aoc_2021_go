package aoc_2021_01

import (
	"github.com/manto89/aoc_2021_go/utils"
	"log"
)

var fileName = "./input/01"

func readLines() {
	ints, err := utils.ReadInts(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if len(ints) < 1 {
		log.Fatalf("File doesn't contain any number")
	}
	lastInt := ints[0]
	increasingCounter := 0
	decreasingCounter := 0
	for i := 1; i < len(ints); i++ {
		if ints[i] > lastInt {
			increasingCounter++
		} else {
			decreasingCounter++
		}
		lastInt = ints[i]
	}
	log.Printf("Increasing %d, Decreasing %d", increasingCounter, decreasingCounter)
}

func Solve() {
	log.Printf("**** DAY 01 ****")
	log.Printf("Given a list of number, count how many numbers are bigger than the previous one")
	readLines()
	log.Printf("****")
}
