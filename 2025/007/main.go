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

	beams := map[int]bool{}
	splits := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for pos, char := range line {
			if char == 'S' {
				beams[pos] = true
			}

			if char == '^' {
				if beam, ok := beams[pos]; ok && beam {
					beams[pos] = false
					splits++

					if beam, ok := beams[pos-1]; !ok || !beam {
						beams[pos-1] = true
					}

					if beam, ok := beams[pos+1]; !ok || !beam {
						beams[pos+1] = true
					}
				}
			}
		}
	}

	log.Printf("Splits: %v", splits)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
