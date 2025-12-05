package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type IngredientRange struct {
	min int
	max int
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mode := "fresh"

	freshRanges := []IngredientRange{}
	available := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if mode != "fresh" {
				log.Fatal("Invalid file format")
			}

			mode = "available"
			continue
		}

		switch mode {
		case "fresh":
			ingredientRange, err := ParseIngredientRange(line)
			if err != nil {
				log.Fatal(err)
			}

			freshRanges = append(freshRanges, ingredientRange)
		case "available":
			availableIngredient, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}

			available = append(available, availableIngredient)
		default:
			log.Fatal("Invalid mode")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	freshAndAvailable := []int{}
	for _, ingredient := range available {
		for _, ingredientRange := range freshRanges {
			if ingredient >= ingredientRange.min && ingredient <= ingredientRange.max {
				freshAndAvailable = append(freshAndAvailable, ingredient)
				break
			}
		}
	}

	removed := map[int]bool{}
	for {
		// TODO: Looks like a dirty trick to make flawed algorithm below work.
		//       For the given input, it retries once.
		changed := false

		for i := range freshRanges {
			if _, ok := removed[i]; ok {
				continue
			}

			for j := range i {
				if _, ok := removed[j]; ok {
					continue
				}

				if (freshRanges[j].min >= freshRanges[i].min &&
					freshRanges[j].max <= freshRanges[i].max) ||
					(freshRanges[i].min >= freshRanges[j].min &&
						freshRanges[i].max <= freshRanges[j].max) {
					// One includes another (absorb)

					freshRanges[i].min = min(freshRanges[i].min, freshRanges[j].min)
					freshRanges[i].max = max(freshRanges[i].max, freshRanges[j].max)
					removed[j] = true
					changed = true
				} else if (freshRanges[i].min <= freshRanges[j].min &&
					freshRanges[i].max >= freshRanges[j].min) ||
					(freshRanges[j].min <= freshRanges[i].min &&
						freshRanges[j].max >= freshRanges[i].min) {
					// One overlaps another (merge)

					freshRanges[i].min = min(freshRanges[i].min, freshRanges[j].min)
					freshRanges[i].max = max(freshRanges[i].max, freshRanges[j].max)
					removed[j] = true
					changed = true
				}
			}
		}

		if !changed {
			break
		}
	}

	freshTotal := 0
	for i := range freshRanges {
		if _, ok := removed[i]; !ok {
			freshTotal += freshRanges[i].max - freshRanges[i].min + 1
		}
	}

	log.Printf("Found %v fresh ingredient ranges and %v available ingredients: "+
		"out of them %v are both fresh and available, total fresh ingredients: %v",
		len(freshRanges), len(available),
		len(freshAndAvailable),
		freshTotal)
}

func ParseIngredientRange(line string) (IngredientRange, error) {
	items := strings.Split(line, "-")
	if len(items) != 2 {
		return IngredientRange{}, errors.New("Invalid range format")
	}

	min, err := strconv.Atoi(items[0])
	if err != nil {
		log.Fatal(err)
	}

	max, err := strconv.Atoi(items[1])
	if err != nil {
		log.Fatal(err)
	}

	return IngredientRange{min, max}, nil
}
