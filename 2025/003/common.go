package main

import (
	"errors"
	"strconv"
)

func FindMaxJoltage(joltageRatings string, batteries int) (string, error) {
	marked := map[int]bool{}

	leftmostVal := -1
	leftmostPos := -1

	for range batteries {
		realMaxVal := -1
		realMaxPos := -1

		maxVal := -1
		maxPos := -1

		for currPos := 0; currPos < len(joltageRatings); currPos++ {
			_, seen := marked[currPos]
			if seen {
				continue
			}

			currVal, err := strconv.Atoi(string(joltageRatings[currPos]))
			if err != nil {
				return "", err
			}

			if currVal >= realMaxVal {
				realMaxVal = currVal
				realMaxPos = currPos
			}

			// skip if before leftmost, unless it's a larger rating
			if currPos <= leftmostPos && currVal < leftmostVal {
				continue
			}

			// ">=" to find rightmost, even if identical
			if currVal >= maxVal {
				maxVal = currVal
				maxPos = currPos
			}
		}

		if maxPos == -1 && realMaxPos != -1 {
			maxPos = realMaxPos
			maxVal = realMaxVal
		}

		if maxPos == -1 {
			return "", errors.New("Could not find max")
		}

		marked[maxPos] = true

		if leftmostPos == -1 || maxPos < leftmostPos {
			leftmostVal = maxVal
			leftmostPos = maxPos
		}
	}

	result := ""

	for pos, char := range joltageRatings {
		_, seen := marked[pos]
		if seen {
			result += string(char)
		}
	}

	return result, nil
}
