package aoc_2021_03

import (
	"github.com/manto89/aoc_2021_go/utils"
	"log"
	"strconv"
	"time"
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
func getLifeSupportReading() (int64, int64) {
	array, err := utils.ReadBinaryNums(fileName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	exclusionListOxygen := make([]bool, len(array), len(array))
	exclusionListCo2 := make([]bool, len(array), len(array))
	for i := 0; i < len(array[0]); i++ {
		valueOxygen := sumBools(func(list [][]bool) []bool {
			var ret []bool
			for r, row := range list {
				if exclusionListOxygen[r] {
					continue
				}
				ret = append(ret, row[i])
			}
			return ret
		}(array)...)
		valueCo2 := sumBools(func(list [][]bool) []bool {
			var ret []bool
			for r, row := range list {
				if exclusionListCo2[r] {
					continue
				}
				ret = append(ret, row[i])
			}
			return ret
		}(array)...)
		//TODO don't loop over excluded rows
		if sumBools(exclusionListOxygen...) < len(array)-1 {
			exclusionListOxygenLength := len(exclusionListOxygen) - sumBools(exclusionListOxygen...)
			for j, row := range array {
				if exclusionListOxygen[j] {
					continue
				}
				//most of the elements are 1
				if float64(valueOxygen) >= float64(exclusionListOxygenLength)/2 {
					if !row[i] {
						exclusionListOxygen[j] = true
					}
				} else { //most of the elements are 0
					if row[i] {
						exclusionListOxygen[j] = true
					}
				}
			}
		}

		if sumBools(exclusionListCo2...) < len(array)-1 {
			exclusionListCo2Length := len(exclusionListCo2) - sumBools(exclusionListCo2...)
			for j, row := range array {
				if exclusionListCo2[j] {
					continue
				}
				//most of the elements are 1
				if float64(valueCo2) >= float64(exclusionListCo2Length)/2 {
					if row[i] {
						exclusionListCo2[j] = true
					}
				} else { //most of the elements are 0
					if !row[i] {
						exclusionListCo2[j] = true
					}
				}
			}

		}
		if sumBools(exclusionListOxygen...) == len(array)-1 && sumBools(exclusionListCo2...) == len(array)-1 {
			break
		}
	}

	oxygen := func() int64 {
		for i, excludedItem := range exclusionListOxygen {
			if !excludedItem {
				ret, err := boolToInt(array[i])
				if err != nil {
					log.Fatalf(err.Error())
				}
				return ret
			}
		}
		return 0
	}()
	co2 := func() int64 {
		for i, excludedItem := range exclusionListCo2 {
			if !excludedItem {
				ret, err := boolToInt(array[i])
				if err != nil {
					log.Fatalf(err.Error())
				}
				return ret
			}
		}
		return 0
	}()
	return oxygen, co2
}

func SolvePart1() {
	log.Printf("**** DAY 03 Part01****")
	log.Printf("Given a list of binary sensor readings, return the power consumption")
	start := time.Now()
	gamma, epsilon := getPowerConsuption()
	elapsed := time.Since(start)
	log.Printf("Power consumption: %d", gamma*epsilon)
	log.Printf("Execution took %s", elapsed)
	log.Printf("****")
}

func SolvePart2() {
	log.Printf("**** DAY 03 Part02****")
	log.Printf("Given a list of binary sensor readings, return the power consumption")
	start := time.Now()
	oxigen, co2 := getLifeSupportReading()
	elapsed := time.Since(start)
	log.Printf("Life support: %d", oxigen*co2)
	log.Printf("Execution took %s", elapsed)
	log.Printf("****")
}
