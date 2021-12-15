package utils

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLines(filename string) ([]string, error) {

	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("File %s doesn't exist", filename))
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("Unable to close file")
		}
	}(f)

	scanner := bufio.NewScanner(f)

	var ret []string
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) < 1 {
			continue
		}
		ret = append(ret, s)
	}
	return ret, nil
}

func ReadInts(filename string) ([]int, error) {
	var ret []int
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		ret = append(ret, i)
	}
	return ret, nil
}

func ReadDirections(filename string) ([]Vector, error) {
	var ret []Vector
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}
		vector, err := ParseVector(parts[0], parts[1])
		if err != nil {
			return nil, err
		}
		ret = append(ret, vector)

	}
	return ret, nil
}

func ReadBinaryNums(filename string) ([][]bool, error) {
	var ret [][]bool
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}
	for _, line := range lines {
		var row []bool
		for _, char := range line {
			if char == '0' {
				row = append(row, false)
			} else if char == '1' {
				row = append(row, true)
			}
		}
		ret = append(ret, row)
	}
	return ret, nil
}
