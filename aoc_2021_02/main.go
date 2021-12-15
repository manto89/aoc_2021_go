package aoc_2021_02

import (
	"github.com/manto89/aoc_2021_go/utils"
	"log"
)

var fileName = "./input/02"

func getFinalPositionOrtho() (int, int) {

	var posX, posY = 0, 0
	dirs, err := utils.ReadDirections(fileName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, dir := range dirs {
		switch dir.Direction {
		case utils.Up:
			posY += dir.Length
		case utils.Down:
			posY -= dir.Length
		case utils.Right:
			posX += dir.Length
		case utils.Left:
			posX -= dir.Length
		}
	}
	return posX, posY
}

func getFinalPositionPolar() (int, int) {
	var posX, posY = 0, 0
	dirs, err := utils.ReadDirections(fileName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	var aim = 0
	for _, dir := range dirs {
		switch dir.Direction {
		case utils.Up:
			aim += dir.Length
		case utils.Down:
			aim -= dir.Length
		case utils.Right:
			posX += dir.Length
			posY += dir.Length * aim
		}
	}
	return posX, posY
}

func SolvePart1() {
	log.Printf("**** DAY 01 Part01****")
	log.Printf("Given a list of directions (orthogonal) with length, get the final position in x,y (where y is depth)")
	posX, posY := getFinalPositionOrtho()
	log.Printf("Final position horizontal pos: %d, depth: %d (Mult:%d)", posX, -posY, posX*(-posY))
	log.Printf("****")
}
func SolvePart2() {
	log.Printf("**** DAY 01 Part02****")
	log.Printf("Given a list of directions (up,down change aim; right given amount of movement), get the final position in x,y (where y is depth)")
	posX, posY := getFinalPositionPolar()
	log.Printf("Final position horizontal pos: %d, depth: %d (Mult:%d)", posX, -posY, posX*(-posY))
	log.Printf("****")
}
