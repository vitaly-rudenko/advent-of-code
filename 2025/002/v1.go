package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Invalid: repeated twice (55, 6464, 123123)
// Task: find all invalid IDs in given ranges
//       then add up all invalid ranges (sum)

func V1() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	invalidIds := []int{}
	reader := bufio.NewReader(file)
	for {
		string, readerErr := reader.ReadString(',')
		if readerErr != nil && readerErr != io.EOF {
			log.Fatal(readerErr)
		}

		stringifiedRange := strings.TrimRight(string, ",\n")
		min, max, err := ParseStringifiedRange(stringifiedRange)
		if err != nil {
			log.Fatal(err)
		}

		invalidIdsInRange, err := FindInvalidIdsInRangeV1(min, max)
		if err != nil {
			log.Fatal(err)
		}

		invalidIds = append(invalidIds, invalidIdsInRange...)

		if readerErr == io.EOF {
			break
		}
	}

	sumOfInvalidIds := 0
	for i := range invalidIds {
		sumOfInvalidIds += invalidIds[i]
	}

	log.Printf("[V1] Found %v invalid IDs, sum: %v.", len(invalidIds), sumOfInvalidIds)
}

func FindInvalidIdsInRangeV1(min int, max int) ([]int, error) {
	halfMin, err := HalfOfInteger(min)
	if err != nil {
		return []int{}, err
	}

	halfMax, err := HalfOfInteger(max)
	if err != nil {
		return []int{}, err
	}

	if halfMax < halfMin {
		halfMin = halfMax
		halfMax = halfMax*10 - 1
	}

	invalidIdsInRange := []int{}

	for half := halfMin; half <= halfMax; half++ {
		stringifiedHalf := strconv.Itoa(half)

		invalidId, err := strconv.Atoi(stringifiedHalf + stringifiedHalf)
		if err != nil {
			return []int{}, err
		}

		if invalidId >= min && invalidId <= max {
			invalidIdsInRange = append(invalidIdsInRange, invalidId)
		}
	}

	return invalidIdsInRange, nil
}

// 1 => 1, 23 => 2, 460 => 4, 1234 => 12
func HalfOfInteger(num int) (int, error) {
	stringified := strconv.Itoa(num)
	if len(stringified) <= 0 {
		return 0, errors.New("Invalid input number")
	}

	if len(stringified) == 1 {
		return num, nil
	}

	half, err := strconv.Atoi(stringified[:len(stringified)/2])
	if err != nil {
		return 0, err
	}

	return half, nil
}
