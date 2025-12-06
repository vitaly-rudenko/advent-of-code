package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"sort"
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

	// V1
	freshAndAvailable := []int{}
	for _, ingredient := range available {
		for _, ingredientRange := range freshRanges {
			if ingredient >= ingredientRange.min && ingredient <= ingredientRange.max {
				freshAndAvailable = append(freshAndAvailable, ingredient)
				break
			}
		}
	}

	// V2
	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i].min < freshRanges[j].min
	})

	freshTotal := 0
	start := 0
	maxMax := freshRanges[0].max
	for i := 1; i <= len(freshRanges); i++ {
		if i < len(freshRanges) && freshRanges[i].min <= maxMax {
			maxMax = max(freshRanges[i].max, maxMax)
			continue
		}

		freshTotal += maxMax - freshRanges[start].min + 1

		if i < len(freshRanges) {
			start = i
			maxMax = freshRanges[i].max
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
