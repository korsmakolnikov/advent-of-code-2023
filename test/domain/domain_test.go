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
