package domain

import "fmt"

var numMap = [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

type Digits struct {
	first int
	last  int
}

func (d *Digits) upsert(n int) {
	if d.first < 10 {
		d.first = n * 10
	}
	d.last = n
}

func (d *Digits) res() int {
	if d.first == 0 {
		panic("Something very wrong happened here")
	}

	return d.first + d.last
}

func (d *Digits) String() string {
	return fmt.Sprintf("%d", d.first+d.last)
}

func FilterExtremitiesDigits(input string) int {
	var d Digits

	for i, c := range input {
		for n, needle := range numMap {
			if n == int(c-'0') {
				d.upsert(n)
				break
			}

			needleLen := len(needle)
			if len(input) >= i+needleLen {
				partial := input[i : i+needleLen]

				if partial == needle {
					d.upsert(n)
				}
			}

		}
	}

	return d.res()
}
