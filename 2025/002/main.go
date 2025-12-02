package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"slices"
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

	invalidIdsV1 := []int{}
	invalidIdsV2 := []int{}
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

		invalidIdsInRangeV1, err := FindInvalidIdsInRangeV1(min, max)
		invalidIdsInRangeV2, err := FindInvalidIdsInRangeV2(min, max)
		if err != nil {
			log.Fatal(err)
		}

		log.Print("Range: ", min, "-", max, " => ", invalidIdsInRangeV2)

		invalidIdsV1 = append(invalidIdsV1, invalidIdsInRangeV1...)
		invalidIdsV2 = append(invalidIdsV2, invalidIdsInRangeV2...)

		if readerErr == io.EOF {
			break
		}
	}

	sumOfInvalidIdsV1 := 0
	for i := range invalidIdsV1 {
		sumOfInvalidIdsV1 += invalidIdsV1[i]
	}

	// TODO: Current V2 algorithm returns duplicates (e.g. 222220-222224)
	invalidIdsV2 = slices.Compact(invalidIdsV2)

	sumOfInvalidIdsV2 := 0
	for i := range invalidIdsV2 {
		sumOfInvalidIdsV2 += invalidIdsV2[i]
	}

	log.Printf("[V1] Found %v invalid IDs, sum: %v.", len(invalidIdsV1), sumOfInvalidIdsV1)
	log.Printf("[V2] Found %v invalid IDs, sum: %v.", len(invalidIdsV2), sumOfInvalidIdsV2)
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

func FindInvalidIdsInRangeV2(min int, max int) ([]int, error) {
	minLength := len(strconv.Itoa(min))
	maxLength := len(strconv.Itoa(max))

	invalidIdsInRange := []int{}

	iter := 0

	for pairLength := 1; pairLength <= maxLength/2; pairLength++ {
		halfMin, err := PartOfInteger(min, pairLength)
		if err != nil {
			return []int{}, err
		}

		halfMax, err := PartOfInteger(max, pairLength)
		if err != nil {
			return []int{}, err
		}

		if halfMax < halfMin {
			halfMin = halfMax
			halfMax = halfMax*10 - 1
		}

		for half := halfMin; half <= halfMax; half++ {
			for repeat := minLength / pairLength; repeat <= maxLength/pairLength; repeat++ {
				stringifiedHalf := strconv.Itoa(half)

				iter++

				invalidId, err := strconv.Atoi(strings.Repeat(stringifiedHalf, repeat))
				if err != nil {
					return []int{}, err
				}

				if invalidId >= min && invalidId <= max {
					invalidIdsInRange = append(invalidIdsInRange, invalidId)
				}
			}
		}
	}

	return invalidIdsInRange, nil
}

func PartOfInteger(num int, length int) (int, error) {
	stringified := strconv.Itoa(num)
	if len(stringified) <= 0 {
		return 0, errors.New("Invalid input number")
	}

	if len(stringified) == 1 {
		return num, nil
	}

	part, err := strconv.Atoi(stringified[:length])
	if err != nil {
		return 0, err
	}

	return part, nil
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
