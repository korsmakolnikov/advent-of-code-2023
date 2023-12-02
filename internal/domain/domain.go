package domain

import (
	"unicode"
)

func FilterExtremitiesDigits(input string) int {
	var partialRes, partialLast int
	for _, c := range input {
		if unicode.IsNumber(c) {
			if partialRes < 10 {
				partialRes = int(c-'0') * 10
				partialLast = int(c - '0')
				continue
			}
			partialLast = int(c - '0')
		}
	}

	return partialRes + partialLast
}
