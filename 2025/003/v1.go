package main

import (
	"bufio"
	"log"
	"math/big"
	"os"
)

func V1() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	maxJoltages := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		joltageRatings := scanner.Text()

		maxJoltage, err := FindMaxJoltage(joltageRatings, 2)
		if err != nil {
			log.Fatal(err)
		}

		maxJoltages = append(maxJoltages, maxJoltage)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	maxJoltagesSum := new(big.Int)
	for i := range maxJoltages {
		maxJoltageInt := new(big.Int)
		_, ok := maxJoltageInt.SetString(maxJoltages[i], 10)
		if !ok {
			log.Fatal("Failed to parse maxJoltage: ", maxJoltages[i])
		}

		maxJoltagesSum.Add(maxJoltagesSum, maxJoltageInt)
	}

	log.Printf("[V1] Found %v maxJoltages, sum: %v", len(maxJoltages), maxJoltagesSum.String())
}
