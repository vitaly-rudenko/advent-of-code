package main

import (
	"errors"
	"strconv"
	"strings"
)

func ParseStringifiedRange(stringifiedRange string) (int, int, error) {
	parts := strings.Split(stringifiedRange, "-")
	if len(parts) != 2 {
		return 0, 0, errors.New("Range must consist of two IDs")
	}

	stringifiedMin := parts[0]
	min, err := strconv.Atoi(stringifiedMin)
	if err != nil {
		return 0, 0, err
	}

	stringifiedMax := parts[1]
	max, err := strconv.Atoi(stringifiedMax)
	if err != nil {
		return 0, 0, err
	}

	return min, max, nil
}
