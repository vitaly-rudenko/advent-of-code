package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	maxJoltages := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		joltageRatings := scanner.Text()

		maxJoltage, err := FindMaxJoltage(joltageRatings)
		if err != nil {
			log.Fatal(err)
		}

		maxJoltages = append(maxJoltages, maxJoltage)

		log.Print(joltageRatings, " -> ", maxJoltage)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	maxJoltagesSum := 0
	for i := range maxJoltages {
		maxJoltagesSum += maxJoltages[i]
	}

	log.Printf("Found %v maxJoltages, sum: %v", len(maxJoltages), maxJoltagesSum)
}

func FindMaxJoltage(joltageRatings string) (int, error) {
	maxJoltage := 0

	for i := 0; i < len(joltageRatings) - 1; i++ {
		joltageRating1, err := strconv.Atoi(string(joltageRatings[i]))
		if err != nil {
			return 0, err
		}

		highest := 0

		for j := i + 1; j < len(joltageRatings); j++ {
			joltageRating2, err := strconv.Atoi(string(joltageRatings[j]))
			if err != nil {
				return 0, err
			}

			if joltageRating2 > highest {
				highest = joltageRating2
			}
		}

		if highest == 0 {
			return 0, errors.New("Error in algorithm")
		}

		newMaxJoltage, err := strconv.Atoi(strconv.Itoa(joltageRating1) + strconv.Itoa(highest))
		if err != nil {
			return 0, err
		}

		if newMaxJoltage > maxJoltage {
			maxJoltage = newMaxJoltage
		}
	}

	return maxJoltage, nil
}
