package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindInvalidIdsInRangeV2_1(t *testing.T) {
	invalidIds, err := FindInvalidIdsInRangeV2(0, 1015)
	if err != nil {
		t.Fatal(err)
	}

	assert.ElementsMatch(t, MapToInts(invalidIds), []int{
		11, 22, 33, 44, 55, 66, 77, 88, 99,
		111, 222, 333, 444, 555, 666, 777, 888, 999, 1010,
	})
}

func TestFindInvalidIdsInRangeV2_2(t *testing.T) {
	invalidIds, err := FindInvalidIdsInRangeV2(33, 177)
	if err != nil {
		t.Fatal(err)
	}

	assert.ElementsMatch(t, MapToInts(invalidIds), []int{
		33, 44, 55, 66, 77, 88, 99,
		111,
	})
}

func TestFindInvalidIdsInRangeV2_3(t *testing.T) {
	invalidIds, err := FindInvalidIdsInRangeV2(0, 10)
	if err != nil {
		t.Fatal(err)
	}

	assert.ElementsMatch(t, MapToInts(invalidIds), []int{})
}

func TestFindInvalidIdsInRangeV2_4(t *testing.T) {
	invalidIds, err := FindInvalidIdsInRangeV2(9525, 11195)
	if err != nil {
		t.Fatal(err)
	}

	assert.ElementsMatch(t, MapToInts(invalidIds), []int{
		9595, 9696, 9797, 9898, 9999,
		11111,
	})
}

func MapToInts(input map[int]bool) []int {
	result := []int{}

	for k, _ := range input {
		result = append(result, k)
	}

	return result
}
