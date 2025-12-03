package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxJoltage_12batteries(t *testing.T) {
	t.Parallel()

	tests := []struct {
		joltageRatings string
		maxJoltage12   string
		maxJoltage5    string
		maxJoltage2    string
	}{
		{"987654321111111", "987654321111", "98765", "98"},
		{"811111111111119", "811111111119", "81119", "89"},
		{"234234234234278", "434234234278", "44478", "78"},
		{"818181911112111", "888911112111", "92111", "92"},
		{"452749582798719", "749582798719", "99879", "99"},
		{"919191919191919", "999919191919", "99999", "99"},
		{"819181918191819", "981918191819", "99989", "99"},
		{"123456789112345", "456789112345", "92345", "95"},
		{"111122233344455", "122233344455", "44455", "55"},
		{"554433221155332", "554433255332", "55553", "55"},
		{"192837465192837", "987465192837", "99837", "99"},
	}

	for _, test := range tests {
		t.Run(test.joltageRatings, func(t *testing.T) {
			t.Parallel()
			t.Log(test.joltageRatings)

			maxJoltage2, err := FindMaxJoltage(test.joltageRatings, 2)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, test.maxJoltage2, maxJoltage2)

			maxJoltage5, err := FindMaxJoltage(test.joltageRatings, 5)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, test.maxJoltage5, maxJoltage5)

			maxJoltage12, err := FindMaxJoltage(test.joltageRatings, 12)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, test.maxJoltage12, maxJoltage12)
		})
	}
}
