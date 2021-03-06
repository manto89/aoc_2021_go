package aoc_2021_01

import (
	"errors"
	"github.com/manto89/aoc_2021_go/utils"
	"log"
)

var fileName = "./input/01"

func getMeasurementsReport(windowSize int) {
	ints, err := utils.ReadInts(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if len(ints) < 1 {
		log.Fatalf("File doesn't contain any number")
	}
	lastMeasurement := 0
	increasingCounter := 0
	decreasingCounter := 0
	for i := 0; i < len(ints); i++ {
		//declared a function so we can use defer to update lastMeasurement
		func() {
			defer func() {
				lastMeasurement, _ = getWindowSum(ints, i, windowSize)
			}()
			if i < 1 {
				return
			}
			newInt, err := getWindowSum(ints, i, windowSize)
			//this happens only when there aren't enough measurements
			if err != nil {
				return
			}
			if newInt > lastMeasurement {
				increasingCounter++
			} else {
				decreasingCounter++
			}
		}()
	}
	log.Printf("Increasing %d, Decreasing %d", increasingCounter, decreasingCounter)
}

// Gets measurement for the selected interval, using the current index and the following # elements (windowSize)
func getWindowSum(ints []int, counter int, windowSize int) (int, error) {
	var ret = 0
	if counter+windowSize > len(ints) {
		return 0, errors.New("Not enough measurements")
	}
	for i := 0; i < windowSize; i++ {
		ret += ints[counter+i]
	}
	return ret, nil
}

func SolvePart1() {
	log.Printf("**** DAY 01 Part01****")
	log.Printf("Given a list of number, count how many numbers are bigger than the previous one")
	getMeasurementsReport(1)
	log.Printf("****")
}

func SolvePart2() {
	log.Printf("**** DAY 01 Part02 ****")
	log.Printf("Given a list of number, count how many numbers are bigger than the previous one using a three-measurement sliding window")
	getMeasurementsReport(3)
	log.Printf("****")
}
