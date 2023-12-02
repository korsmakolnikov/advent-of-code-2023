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
		t.Errorf("FilterExtremitiesDigits fails with %s, %d was expected but it returns %d", input, expected, result)
	}
}

func TestFilterExtremitiesDigitsRoundTwo(t *testing.T) {
	input := "pqr3stu8vwx "
	result := domain.FilterExtremitiesDigits(input)
	expected := 38

	if expected != result {
		t.Errorf("FilterExtremitiesDigits fails with %s, %d was expected but it returns %d", input, expected, result)
	}
}

func TestFilterExtremitiesDigitsWithMultipleDigits(t *testing.T) {
	input := "a1b2c3d4e5f "
	result := domain.FilterExtremitiesDigits(input)
	expected := 15

	if expected != result {
		t.Errorf("FilterExtremitiesDigits fails with %s, %d was expected but it returns %d", input, expected, result)
	}
}

func TestFilterExtremitiesDigitsWithASingleDigit(t *testing.T) {
	input := "treb7uchet "
	result := domain.FilterExtremitiesDigits(input)
	expected := 77

	if expected != result {
		t.Errorf("FilterExtremitiesDigits fails with %s, %d was expected but it returns %d", input, expected, result)
	}
}

func TestFilterExtremitiesDigitsWithTwoAndNineInLetters(t *testing.T) {
	input := "two1nine"
	result := domain.FilterExtremitiesDigits(input)
	expected := 29

	if expected != result {
		t.Errorf("FilterExtremitiesDigits fails with %s, %d was expected but it returns %d", input, expected, result)
	}
}

func TestFilterExtremitiesDigitsWithEightAndThreeInLetters(t *testing.T) {
	input := "eightwothree"
	result := domain.FilterExtremitiesDigits(input)
	expected := 83

	if expected != result {
		t.Errorf("FilterExtremitiesDigits fails with %s, %d was expected but it returns %d", input, expected, result)
	}
}

func TestFilterExtremitiesDigitsWithOneAndFourInLetters(t *testing.T) {
	input := "xtwone3four"
	result := domain.FilterExtremitiesDigits(input)
	expected := 24

	if expected != result {
		t.Errorf("FilterExtremitiesDigits fails with %s, %d was expected but it returns %d", input, expected, result)
	}
}

func TestFilterExtremitiesDigitsWithFourAndTwoInNumbers(t *testing.T) {
	input := "4nineeightseven2"
	result := domain.FilterExtremitiesDigits(input)
	expected := 42

	if expected != result {
		t.Errorf("FilterExtremitiesDigits fails with %s, %d was expected but it returns %d", input, expected, result)
	}
}
func TestFilterExtremitiesDigitsWithNumbersInLetters(t *testing.T) {

	var input = [...]string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	var expected = [...]int{29, 83, 13, 24, 42, 14, 76}
	var finalResult int
	for n, i := range input {

		result := domain.FilterExtremitiesDigits(i)
		e := expected[n]

		if e != result {
			t.Errorf("FilterExtremitiesDigits fails with %s, %d was expected but it returns %d", i, e, result)
		}
		finalResult += result
	}

	if 281 != finalResult {
		t.Errorf("FilterExtremitiesDigits fails with 281 was expected but it returns %d", finalResult)
	}
}

func TestEdgeCase(t *testing.T) {
	input := "9rnjqlpq"
	result := domain.FilterExtremitiesDigits(input)
	expected := 99

	if expected != result {
		t.Errorf("FilterExtremitiesDigits fails with %s, %d was expected but it returns %d", input, expected, result)
	}
}
