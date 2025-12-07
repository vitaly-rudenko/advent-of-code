package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDigitAt_Zero(t *testing.T) {
	var result int
	var err error

	result, err = GetDigitAt(0, 0)
	assert.Equal(t, 0, result)
	assert.NoError(t, err)

	result, err = GetDigitAt(0, 1)
	assert.Equal(t, -1, result)
	assert.NoError(t, err)
}

func TestGetDigitAt_Positive(t *testing.T) {
	var result int
	var err error

	result, err = GetDigitAt(456, 0)
	assert.Equal(t, 4, result)
	assert.NoError(t, err)

	result, err = GetDigitAt(456, 1)
	assert.Equal(t, 5, result)
	assert.NoError(t, err)

	result, err = GetDigitAt(456, 2)
	assert.Equal(t, 6, result)
	assert.NoError(t, err)

	result, err = GetDigitAt(456, 3)
	assert.Equal(t, -1, result)
	assert.NoError(t, err)
}

func TestGetDigitAt_Negative(t *testing.T) {
	var result int
	var err error

	result, err = GetDigitAt(-456, 0)
	assert.Equal(t, 4, result)
	assert.NoError(t, err)

	result, err = GetDigitAt(-456, 1)
	assert.Equal(t, 5, result)
	assert.NoError(t, err)

	result, err = GetDigitAt(-456, 2)
	assert.Equal(t, 6, result)
	assert.NoError(t, err)

	result, err = GetDigitAt(-456, 3)
	assert.Equal(t, -1, result)
	assert.NoError(t, err)
}

func TestGetDigitAt_InvalidDigitPosition(t *testing.T) {
	var err error

	_, err = GetDigitAt(456, -1)
	assert.EqualError(t, err, "Invalid digit position")
}

func TestBuildNumber(t *testing.T) {
	var result int
	var err error

	result, err = BuildNumber([]int{0})
	assert.Equal(t, 0, result)
	assert.NoError(t, err)

	result, err = BuildNumber([]int{1, 0, 0, 1})
	assert.Equal(t, 1001, result)
	assert.NoError(t, err)

	result, err = BuildNumber([]int{4, 5, 6, 0})
	assert.Equal(t, 4560, result)
	assert.NoError(t, err)

	_, err = BuildNumber([]int{})
	assert.EqualError(t, err, "Empty digits list")

	_, err = BuildNumber([]int{-1})
	assert.EqualError(t, err, "Invalid digit")

	_, err = BuildNumber([]int{10})
	assert.EqualError(t, err, "Invalid digit")
}
