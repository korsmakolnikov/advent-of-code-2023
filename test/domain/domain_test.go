package domain_test

import (
	"korsmakolnikov/advent_of_code_2023_day_one/internal/domain"
	"testing"
)

func TestFilterExtremitiesDigits(t *testing.T) {
	input := "1abc2"
	result := domain.FilterExtremitiesDigits(input)
	expected := 12

	if expected != result {
		t.Errorf("filterExtremitiesDigits fails with %s, %d was expected but it returns %d", input, expected, result)
	}
}

func TestFilterExtremitiesDigitsRoundTwo(t *testing.T) {
	input := "pqr3stu8vwx "
	result := domain.FilterExtremitiesDigits(input)
	expected := 38

	if expected != result {
		t.Errorf("filterExtremitiesDigits fails with %s, %d was expected but it returns %d", input, expected, result)
	}
}

func TestFilterExtremitiesDigitsWithMultipleDigits(t *testing.T) {
	input := "a1b2c3d4e5f "
	result := domain.FilterExtremitiesDigits(input)
	expected := 15

	if expected != result {
		t.Errorf("filterExtremitiesDigits fails with %s, %d was expected but it returns %d", input, expected, result)
	}
}

func TestFilterExtremitiesDigitsWithASingleDigit(t *testing.T) {
	input := "treb7uchet "
	result := domain.FilterExtremitiesDigits(input)
	expected := 77

	if expected != result {
		t.Errorf("filterExtremitiesDigits fails with %s, %d was expected but it returns %d", input, expected, result)
	}
}
