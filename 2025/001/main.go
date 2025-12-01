package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
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

	dial := 50
	zeroesAfterRotation := 0
	zeroesDuringDuration := 0
	turns := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		ticks, err := RotationToTicks(line)
		if err != nil {
			log.Fatal(err)
		}

		updatedDial, zeroes := AddTicks(dial, ticks)
		dial = updatedDial

		turns++
		zeroesDuringDuration += zeroes

		if dial == 0 {
			zeroesAfterRotation++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf(
		`Dial turned %v times, reaching zero %v times after the operation`+
			`and %v times during the operation, final dial position: %v`,
		turns, zeroesAfterRotation,
		zeroesDuringDuration, dial,
	)
}

// Extracts ticks from rotation string: "L39" => -39; "R13" => 13
func RotationToTicks(rotation string) (int, error) {
	sign := 0
	if strings.HasPrefix(rotation, "L") {
		sign = -1
	} else if strings.HasPrefix(rotation, "R") {
		sign = 1
	} else {
		return 0, errors.New("Invalid rotation prefix")
	}

	stringifiedTicks := rotation[1:]
	ticks, err := strconv.Atoi(stringifiedTicks)
	if err != nil {
		return 0, err
	}
	if ticks < 0 {
		return 0, errors.New("Unexpected negative ticks")
	}

	return ticks * sign, nil
}

// Adds ticks to the dial (accounting for the wrap around)
// Returns final dial position and amount of times it reached zero
func AddTicks(dial int, ticks int) (int, int) {
	zeroes := int(math.Abs(float64(ticks) / 100)) // -190 => 1; 359 => 3
	ticks = ticks - (ticks/100)*100               // -190 => -90; 359 => 59

	result := dial + ticks
	if result < 0 || result >= 100 {
		zeroes++
	}

	return (100 + result) % 100, zeroes
}
