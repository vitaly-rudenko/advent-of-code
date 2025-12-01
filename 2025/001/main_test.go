package main

import (
	"testing"
)

// RotationToTicks()

func TestRotationToTicksLeft(t *testing.T) {
	rotation := "L31"
	want := -31
	ticks, err := RotationToTicks(rotation)
	if ticks != want || err != nil {
		t.Errorf(`RotationToTicks(%v) = %v, %v, want match for %v, nil`, rotation, ticks, err, want)
	}
}

func TestRotationToTicksRight(t *testing.T) {
	rotation := "R46"
	want := 46
	ticks, err := RotationToTicks(rotation)
	if ticks != want || err != nil {
		t.Errorf(`RotationToTicks(%v) = %v, %v, want match for %v, nil`, rotation, ticks, err, want)
	}
}

func TestRotationToTicksLeftZero(t *testing.T) {
	ticks, err := RotationToTicks("L0")
	if ticks != 0 || err != nil {
		t.Errorf(`RotationToTicks("L0") = %v, %v, want match for %v, nil`, ticks, err, 0)
	}

	ticks, err = RotationToTicks("R0")
	if ticks != 0 || err != nil {
		t.Errorf(`RotationToTicks("R0") = %v, %v, want match for %v, nil`, ticks, err, 0)
	}
}

func TestRotationToTicksInvalid(t *testing.T) {
	ticks, err := RotationToTicks("")
	if ticks != 0 || err == nil {
		t.Errorf(`RotationToTicks("") should have failed`)
	}

	ticks, err = RotationToTicks("U12")
	if ticks != 0 || err == nil {
		t.Errorf(`RotationToTicks("U12") should have failed`)
	}

	ticks, err = RotationToTicks("123")
	if ticks != 0 || err == nil {
		t.Errorf(`RotationToTicks("123") should have failed`)
	}

	ticks, err = RotationToTicks("L")
	if ticks != 0 || err == nil {
		t.Errorf(`RotationToTicks("L") should have failed`)
	}

	ticks, err = RotationToTicks("R-12")
	if ticks != 0 || err == nil {
		t.Errorf(`RotationToTicks("R-12") should have failed`)
	}

	ticks, err = RotationToTicks("L12.34")
	if ticks != 0 || err == nil {
		t.Errorf(`RotationToTicks("R-12") should have failed`)
	}
}

// AddTicks()

func TestAddTicksPositive(t *testing.T) {
	dial := 53
	ticks := 25
	want_result := 78
	want_zeroes := 0
	result, zeroes := AddTicks(dial, ticks)
	if result != want_result || zeroes != want_zeroes {
		t.Errorf(`AddTicks(%v, %v) = %v, want for %v`, dial, ticks, result, want_result)
	}
}

func TestAddTicksPositiveOverflow(t *testing.T) {
	dial := 52
	ticks := 48
	want_result := 0
	want_zeroes := 1
	result, zeroes := AddTicks(dial, ticks)
	if result != want_result || zeroes != want_zeroes {
		t.Errorf(`AddTicks(%v, %v) = %v, %v, want for %v, %v`, dial, ticks, result, zeroes, want_result, want_zeroes)
	}
}

func TestAddTicksPositiveCircle(t *testing.T) {
	dial := 43
	ticks := 100
	want_result := dial
	want_zeroes := 1
	result, zeroes := AddTicks(dial, ticks)
	if result != want_result || zeroes != want_zeroes {
		t.Errorf(`AddTicks(%v, %v) = %v, %v, want for %v, %v`, dial, ticks, result, zeroes, want_result, want_zeroes)
	}
}

func TestAddTicksNegative(t *testing.T) {
	dial := 53
	ticks := -21
	want_result := 32
	want_zeroes := 0
	result, zeroes := AddTicks(dial, ticks)
	if result != want_result || zeroes != want_zeroes {
		t.Errorf(`AddTicks(%v, %v) = %v, %v, want for %v, %v`, dial, ticks, result, zeroes, want_result, want_zeroes)
	}
}

func TestAddTicksNegativeOverflow(t *testing.T) {
	dial := 50
	ticks := -68
	want_result := 82
	want_zeroes := 1
	result, zeroes := AddTicks(dial, ticks)
	if result != want_result || zeroes != want_zeroes {
		t.Errorf(`AddTicks(%v, %v) = %v, %v, want for %v, %v`, dial, ticks, result, zeroes, want_result, want_zeroes)
	}
}

func TestAddTicksNegativeCircle(t *testing.T) {
	dial := 43
	ticks := -1000
	want_result := dial
	want_zeroes := 10
	result, zeroes := AddTicks(dial, ticks)
	if result != want_result || zeroes != want_zeroes {
		t.Errorf(`AddTicks(%v, %v) = %v, %v, want for %v, %v`, dial, ticks, result, zeroes, want_result, want_zeroes)
	}
}

func TestAddTicksEdgeCases(t *testing.T) {
	result, zeroes := AddTicks(50, 1000)
	if result != 50 || zeroes != 10 {
		t.Errorf(`AddTicks(50, 10) produced invalid result`)
	}

	result, zeroes = AddTicks(0, 999)
	if result != 99 || zeroes != 9 {
		t.Log(result, zeroes)
		t.Errorf(`AddTicks(0, 999) produced invalid result`)
	}

	result, zeroes = AddTicks(0, -999)
	if result != 1 || zeroes != 10 {
		t.Log(result, zeroes)
		t.Errorf(`AddTicks(0, -999) produced invalid result`)
	}
}
