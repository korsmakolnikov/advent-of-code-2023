package domain

import (
	"fmt"
	"unicode"
)

func FilterExtremitiesDigits(input string) int {
	var partialRes, partialLast int
	for _, c := range input {
		if unicode.IsNumber(c) {
			fmt.Println("CHECK 2")
			if partialRes < 10 {
				partialRes = int(c-'0') * 10
				continue

			}
			partialLast = int(c - '0')
		}
	}

	return partialRes + partialLast
}
