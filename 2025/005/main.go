package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mode := "fresh"

	freshIngredientRanges := [][]int{}
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
			if availableIngredient >= freshIngredientRange[0] && availableIngredient <= freshIngredientRange[1] {
				freshAndAvailableIngredients = append(freshAndAvailableIngredients, availableIngredient)
				break
			}
		}
	}

	log.Print(freshIngredientRanges)
	log.Print(availableIngredients)

	log.Printf("Found %v fresh ingredient ranges and %v available ingredients: "+
		"out of them %v are both fresh and available",
		len(freshIngredientRanges), len(availableIngredients),
		len(freshAndAvailableIngredients))
}

func ParseIngredientRange(line string) ([]int, error) {
	items := strings.Split(line, "-")
	if len(items) != 2 {
		return []int{}, errors.New("Invalid range format")
	}

	from, err := strconv.Atoi(items[0])
	if err != nil {
		log.Fatal(err)
	}

	to, err := strconv.Atoi(items[1])
	if err != nil {
		log.Fatal(err)
	}

	return []int{from, to}, nil
}
