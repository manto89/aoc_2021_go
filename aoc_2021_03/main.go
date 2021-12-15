package aoc_2021_03

import (
	"github.com/manto89/aoc_2021_go/utils"
	"log"
	"strconv"
)

var fileName = "./input/03"

func getPowerConsuption() (int64, int64) {
	array, err := utils.ReadBinaryNums(fileName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	var gamma, epsilon []bool
	for i := 0; i < len(array[0]); i++ {
		value := sumBools(func(bools [][]bool) []bool {
			var ret []bool
			for _, el := range bools {
				ret = append(ret, el[i])
			}
			return ret
		}(array)...)
		if value > len(array)/2 {
			gamma = append(gamma, true)
			epsilon = append(epsilon, false)
		} else {
			gamma = append(gamma, false)
			epsilon = append(epsilon, true)
		}
	}
	g, err := boolToInt(gamma)
	if err != nil {
		log.Fatalf(err.Error())
	}
	e, err := boolToInt(epsilon)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return g, e
}

func boolToInt(bools []bool) (int64, error) {
	i, err := strconv.ParseInt(func(bools []bool) string {
		var ret string
		for _, b := range bools {
			if b {
				ret += "1"
			} else {
				ret += "0"
			}
		}
		return ret
	}(bools), 2, 64)
	if err != nil {
		return 0, err
	}
	return i, nil

}

func sumBools(bools ...bool) int {
	res := 0
	for _, b := range bools {
		if b {
			res++
		}
	}
	return res
}

func SolvePart1() {
	log.Printf("**** DAY 01 Part01****")
	log.Printf("Given a list of binary sensor readings, return the power consumption")
	gamma, epsilon := getPowerConsuption()
	log.Printf("Power consumption: %d", gamma*epsilon)
	log.Printf("****")
}
