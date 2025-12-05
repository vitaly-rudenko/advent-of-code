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

	freshIngredientRanges := []IngredientRange{}
	availableIngredients := []int{}

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

			freshIngredientRanges = append(freshIngredientRanges, ingredientRange)
		case "available":
			availableIngredient, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}

			availableIngredients = append(availableIngredients, availableIngredient)
		default:
			log.Fatal("Invalid mode")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	freshAndAvailableIngredients := []int{}

	for _, availableIngredient := range availableIngredients {
		for _, freshIngredientRange := range freshIngredientRanges {
			if availableIngredient >= freshIngredientRange.min && availableIngredient <= freshIngredientRange.max {
				freshAndAvailableIngredients = append(freshAndAvailableIngredients, availableIngredient)
				break
			}
		}
	}

	log.Printf("Found %v fresh ingredient ranges and %v available ingredients: "+
		"out of them %v are both fresh and available",
		len(freshIngredientRanges), len(availableIngredients),
		len(freshAndAvailableIngredients))
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
