package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxJoltage_2batteries(t *testing.T) {
	t.Parallel()

	tests := []struct {
		joltageRatings string
		maxJoltage     string
	}{
		{"987654321111111", "98"},
		{"811111111111119", "89"},
		{"234234234234278", "78"},
		{"818181911112111", "92"},
	}

	for _, test := range tests {
		t.Run(test.joltageRatings, func(t *testing.T) {
			t.Parallel()
			t.Log(test.joltageRatings)

			maxJoltage, err := FindMaxJoltage(test.joltageRatings, 2)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, maxJoltage, test.maxJoltage)
		})
	}
}

func TestFindMaxJoltage_12batteries(t *testing.T) {
	t.Parallel()

	tests := []struct {
		joltageRatings string
		maxJoltage     string
	}{
		{"987654321111111", "987654321111"},
		{"811111111111119", "811111111119"},
		{"234234234234278", "434234234278"},
		{"818181911112111", "888911112111"},
		{"452749582798719", "749582798719"},
		{"919191919191919", "999919191919"},
		{"819181918191819", "981918191819"},
		{"123456789112345", "456789112345"},
		{"111122233344455", "122233344455"},
		{"554433221155332", "554433255332"},
		{"192837465192837", "983746592837"},
	}

	for _, test := range tests {
		t.Run(test.joltageRatings, func(t *testing.T) {
			t.Parallel()
			t.Log(test.joltageRatings)

			maxJoltage, err := FindMaxJoltage(test.joltageRatings, 12)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, maxJoltage, test.maxJoltage)
		})
	}
}
