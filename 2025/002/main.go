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

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	allInvalidIds := []int{}
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

		log.Print("stringifiedRange ", stringifiedRange, " -> min ", min, " & max ", max)

		invalidIds, err := ExtractInvalidIds(min, max)
		if err != nil {
			log.Fatal(err)
		}

		log.Print("invalidIds: ", invalidIds)
		allInvalidIds = append(allInvalidIds, invalidIds...)

		if readerErr == io.EOF {
			break
		}
	}

	allInvalidIdsSum := 0
	for i := range allInvalidIds {
		allInvalidIdsSum += allInvalidIds[i]
	}

	log.Print("allInvalidIds: ", allInvalidIds, " -> allInvalidIdsSum: ", allInvalidIdsSum)
}

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

func ExtractInvalidIds(min int, max int) ([]int, error) {
	stringifiedMin := strconv.Itoa(min)
	stringifiedMax := strconv.Itoa(max)

	minLength := len(stringifiedMin)
	maxLength := len(stringifiedMax)

	for length := minLength; length <= maxLength; length++ {
		if length%2 == 0 {
			invalidIds := []int{}

			firstPartMin, err := strconv.Atoi(stringifiedMin[:length/2])
			if err != nil {
				return nil, err
			}

			firstPartMax, err := strconv.Atoi(stringifiedMax[:length/2])
			if err != nil {
				return nil, err
			}

			if firstPartMax < firstPartMin {
				temp := firstPartMin
				firstPartMin = firstPartMax
				firstPartMax = temp
			}

			for firstPart := firstPartMin; firstPart <= firstPartMax; firstPart++ {
				generatedNumber, err := strconv.Atoi(strings.Repeat(strconv.Itoa(firstPart), 2))
				if err != nil {
					return nil, err
				}

				if generatedNumber >= min && generatedNumber <= max {
					invalidIds = append(invalidIds, generatedNumber)
				}
			}

			return invalidIds, nil
		}
	}

	return []int{}, nil
}
