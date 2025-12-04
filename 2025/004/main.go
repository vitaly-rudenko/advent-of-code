package main

import (
	"bufio"
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

	matrix := CreateMatrixFromLines(lines)
	// for i := range len(matrix) {
	// 	log.Print(matrix[i])
	// }
	//
	// log.Print("")

	PopulateMatrix(matrix)
	// for i := range len(matrix) {
	// 	log.Print(matrix[i])
	// }
	//
	// log.Print("")
	//
	// for i := range len(matrix) {
	// 	items := []int{}
	//
	// 	for j := range len(matrix[i]) {
	// 		if matrix[i][j] > 0 && matrix[i][j] <= 4 {
	// 			items = append(items, 1)
	// 		} else {
	// 			items = append(items, 0)
	// 		}
	// 	}
	//
	// 	log.Print(items)
	// }

	count := CountItemsBetween(matrix, 1, 4)

	log.Printf("Matrix %vx%v: %v", len(matrix), len(matrix[0]), count)
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
			if matrix[i][j] > 0 {
				for ii := -1; ii <= 1; ii++ {
					for jj := -1; jj <= 1; jj++ {
						if ii == 0 && jj == 0 {
							continue
						}

						x := i + ii
						y := j + jj

						if x >= 0 && y >= 0 && x < len(matrix) && y < len(matrix[i]) {
							if matrix[x][y] > 0 {
								matrix[x][y]++
							}
						}
					}
				}
			}
		}
	}
}

func CreateMatrixFromLines(lines []string) [][]int {
	matrix := [][]int{}

	for _, line := range lines {
		items := []int{}

		for _, char := range line {
			if char == '@' {
				items = append(items, 1)
			} else {
				items = append(items, 0)
			}
		}

		matrix = append(matrix, items)
	}

	return matrix
}
