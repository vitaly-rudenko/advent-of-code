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

	numbers := [][]int{}
	previousLine := ""
	line := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if previousLine != "" {
			extracted, err := ExtractNumbers(previousLine)
			if err != nil {
				log.Fatal(err)
			}

			numbers = append(numbers, extracted)
		}

		line = scanner.Text()

		previousLine = line
	}

	operators, err := ExtractOperators(line)
	results := []int{}

	for i, operator := range operators {
		if operator == "*" {
			results = append(results, 1)
		} else {
			results = append(results, 0)
		}

		for j := range len(numbers) {
			if operator == "*" {
				results[i] *= numbers[j][i]
			} else {
				results[i] += numbers[j][i]
			}
		}
	}

	sum := 0
	for _, result := range results {
		sum += result
	}

	log.Print(results, " ", sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func ExtractNumbers(line string) ([]int, error) {
	parts := strings.Split(line, " ")
	numbers := []int{}

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		number, err := strconv.Atoi(part)
		if err != nil {
			return []int{}, err
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func ExtractOperators(line string) ([]string, error) {
	parts := strings.Split(line, " ")
	operators := []string{}

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if part == "*" || part == "+" {
			operators = append(operators, part)
		} else {
			return []string{}, errors.New("Invalid operator detected")
		}
	}

	return operators, nil
}
