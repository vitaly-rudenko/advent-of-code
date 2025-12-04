package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	matrix, err := CreateMatrixFromLines(lines, map[rune]int{'@': 1, '.': 0})
	if err != nil {
		log.Fatal(err)
	}

	PopulateMatrix(matrix)

	count := CountItemsBetween(matrix, 1, 4)

	removedCount := RecursivelyRemoveItemsBetween(matrix, 1, 4)

	log.Printf("Matrix %vx%v, count: %v, removedCount: %v", len(matrix), len(matrix[0]), count, removedCount)
}

func CountItemsBetween(matrix [][]int, min int, max int) int {
	count := 0

	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			if matrix[i][j] >= min && matrix[i][j] <= max {
				count++
			}
		}
	}

	return count
}

func PopulateMatrix(matrix [][]int) {
	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			if matrix[i][j] == 0 {
				continue
			}

			for ii := -1; ii <= 1; ii++ {
				for jj := -1; jj <= 1; jj++ {
					if ii == 0 && jj == 0 {
						continue
					}

					x := i + ii
					y := j + jj

					if x >= 0 && y >= 0 &&
						x < len(matrix) && y < len(matrix[i]) &&
						matrix[x][y] > 0 {
						matrix[x][y]++
					}
				}
			}
		}
	}
}

func RecursivelyRemoveItemsBetween(matrix [][]int, min int, max int) int {
	removedCountTotal := 0

	for {
		removedCountIteration := 0

		for i := range len(matrix) {
			for j := range len(matrix[i]) {
				if matrix[i][j] < min || matrix[i][j] > max {
					continue
				}

				matrix[i][j] = 0
				removedCountIteration++

				for ii := -1; ii <= 1; ii++ {
					for jj := -1; jj <= 1; jj++ {
						if ii == 0 && jj == 0 {
							continue
						}

						x := i + ii
						y := j + jj

						if x >= 0 && y >= 0 &&
							x < len(matrix) && y < len(matrix[i]) &&
							matrix[x][y] > 0 {
							matrix[x][y]--
						}
					}
				}
			}
		}

		if removedCountIteration == 0 {
			break
		}

		removedCountTotal += removedCountIteration
	}

	return removedCountTotal
}

func CreateMatrixFromLines(lines []string, mappings map[rune]int) ([][]int, error) {
	matrix := [][]int{}

	for _, line := range lines {
		items := []int{}

		for _, char := range line {
			for rune, value := range mappings {
				if char == rune {
					items = append(items, value)
					break
				}
			}
		}

		if len(items) != len(line) {
			return [][]int{}, errors.New("Unmapped characters detected")
		}

		matrix = append(matrix, items)
	}

	return matrix, nil
}
