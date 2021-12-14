package aoc_2021_02

import (
	"github.com/manto89/aoc_2021_go/utils"
	"log"
)

var fileName = "./input/02"

func getFinalPosition() (int, int) {

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

func SolvePart1() {
	log.Printf("**** DAY 01 Part01****")
	log.Printf("Given a list of number, count how many numbers are bigger than the previous one")
	posX, posY := getFinalPosition()
	log.Printf("Final position horizontal pos: %d, depth: %d (Mult:%d)", posX, -posY, posX*(-posY))
	log.Printf("****")
}
