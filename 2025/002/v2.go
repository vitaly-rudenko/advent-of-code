package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Invalid: repeated at least twice (55, 999, 646464, 123123)
// Task: find all invalid IDs in given ranges
//       then add up all invalid ranges (sum)

func V2() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	invalidIds := map[int]bool{}
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

		invalidIdsInRange, err := FindInvalidIdsInRangeV2(min, max)
		if err != nil {
			log.Fatal(err)
		}

		for invalidId, _ := range invalidIdsInRange {
			invalidIds[invalidId] = true
		}

		if readerErr == io.EOF {
			break
		}
	}

	sumOfInvalidIds := 0
	for invalidId, _ := range invalidIds {
		sumOfInvalidIds += invalidId
	}

	log.Printf("[V2] Found %v invalid IDs, sum: %v.", len(invalidIds), sumOfInvalidIds)
}

func FindInvalidIdsInRangeV2(min int, max int) (map[int]bool, error) {
	maxLength := len(strconv.Itoa(max))

	invalidIdsInRange := map[int]bool{}

	for partLength := 1; partLength <= maxLength/2; partLength++ {
		for part := int(math.Pow10(partLength - 1)); part <= int(math.Pow10(partLength))-1; part++ {
			for repeat := 2; repeat <= maxLength/partLength; repeat++ {
				stringifiedPart := strconv.Itoa(part)

				invalidId, err := strconv.Atoi(strings.Repeat(stringifiedPart, repeat))
				if err != nil {
					return map[int]bool{}, err
				}

				if invalidId >= min && invalidId <= max {
					invalidIdsInRange[invalidId] = true
				}
			}
		}
	}

	return invalidIdsInRange, nil
}
