package main

import (
	"bufio"
	"errors"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./example_input.txt")
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

	log.Print(sum)

	results2 := []int{}
	for i, operator := range operators {
		if operator == "*" {
			results2 = append(results, 1)
		} else {
			results2 = append(results, 0)
		}

		maxDigits := 0
		for j := range numbers {
			maxDigits = max(maxDigits, GetDigits(numbers[j][i]))
		}

		for pos := range maxDigits {
			digits := []int{}
			for j := range numbers {
				digit, err := GetDigitAt(numbers[j][i], pos)
				if err != nil {
					log.Fatal(err)
				}

				if digit != -1 {
					digits = append(digits, digit)
				}
			}

			log.Printf("maxDigits: %v, digits: %v at pos: %v", maxDigits, digits, pos)

			number, err := BuildNumber(digits)
			if err != nil {
				log.Fatal(err)
			}

			if operator == "*" {
				results2[i] *= number
			} else {
				results2[i] += number
			}
		}
	}

	log.Print(results2)

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

func GetDigitAt(number int, pos int) (int, error) {
	if pos < 0 {
		return -1, errors.New("Invalid digit position")
	}

	digits := GetDigits(number)
	if pos >= digits {
		return -1, nil
	}

	digit := int(math.Trunc(math.Abs(float64(number))/math.Pow10(digits-pos-1))) % 10
	return digit, nil
}

func GetDigits(number int) int {
	if number == 0 {
		return 1
	}

	return int(math.Floor(math.Log10(math.Abs(float64(number))))) + 1.0
}

func BuildNumber(digits []int) (int, error) {
	if len(digits) == 0 {
		return 0, errors.New("Empty digits list")
	}

	number := 0
	for i, digit := range digits {
		if digit < 0 || digit > 9 {
			return 0, errors.New("Invalid digit")
		}

		number += digit * int(math.Pow10(len(digits)-i-1))
	}

	return number, nil
}
