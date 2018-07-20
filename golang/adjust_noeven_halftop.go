package main

/* Only process odd numbers on the top half i>=limit/2
From 4.5 sec to  1.9 sec */
import (
	"fmt"
)

var countB = 0
var countSB = 0
var iB = 0
var iSB = 0
var ISBsumm = 1
var rest = 0
var limit = 10000000

func main() {
	for i := (limit / 2); i <= limit; i++ {
		n := i
		count := 1
		rest = n % 2
		if rest == 0 {
			continue
		}
		for n > 1 {
			rest = n % 2
			if rest == 0 {
				n = n / 2
			} else {
				n = n*3 + 1
				count++
			}
		}

		if count > countSB && count > countB {
			countSB = countB
			countB = count
			iSB = iB
			iB = i
		}
		if countB > count && count > countSB {
			countSB = count
			iSB = i
		}
	}
	for iSB > 1 {
		ISBsumm += iSB
		rest = iSB % 2
		if rest == 0 {
			iSB = iSB / 2

		} else {
			iSB = iSB*3 + 1

		}

	}
	fmt.Println("\n", "Summ of slices for the second biggest length:", ISBsumm, "\n")
}
