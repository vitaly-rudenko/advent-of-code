package main

import (
	"errors"
	"strconv"
)

func FindMaxJoltage(joltageRatings string, batteries int) (string, error) {
	seenMap := map[int]bool{}

	for range batteries {
		// Find rightmost gap
		gapPos := -1

		for pos := len(joltageRatings) - 1; pos >= 0; pos-- {
			_, seen := seenMap[pos]

			if seen {
				if gapPos != -1 {
					break
				}
			} else {
				gapPos = pos
			}
		}

		if gapPos == -1 {
			return "", errors.New("Could not find gap")
		}

		// Find leftmost max value in that gap
		maxPos := -1
		maxVal := -1

		for pos := gapPos; pos < len(joltageRatings); pos++ {
			_, seen := seenMap[pos]
			if seen {
				continue
			}

			val, err := strconv.Atoi(string(joltageRatings[pos]))
			if err != nil {
				return "", err
			}

			if val > maxVal {
				maxPos = pos
				maxVal = val
			}
		}

		if maxPos == -1 {
			return "", errors.New("Could not find max in gap")
		}

		// Mark is as seen
		seenMap[maxPos] = true
	}

	// Extract digits as seen positions
	maxJoltage := ""

	for pos, char := range joltageRatings {
		_, seen := seenMap[pos]
		if seen {
			maxJoltage += string(char)
		}
	}

	return maxJoltage, nil
}
