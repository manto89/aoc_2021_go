package utils

import (
	"errors"
	"fmt"
	"strconv"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Vector struct {
	Direction Direction
	Length    int
}

func ParseVector(dir string, length string) (Vector, error) {
	var ret Vector
	l, err := strconv.Atoi(length)
	if err != nil {
		return Vector{}, err
	}
	ret.Length = l
	switch dir {
	case "up":
		ret.Direction = Up
	case "down":
		ret.Direction = Down
	case "forward":
		ret.Direction = Right
	default:
		return Vector{}, errors.New(fmt.Sprintf("Unable to parse %s", dir))
	}
	return ret, nil
}
