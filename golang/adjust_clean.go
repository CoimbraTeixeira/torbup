package main

import (
	"fmt"
)

var countB = 0
var countSB = 0
var iB = 0
var iSB = 0
var ISBsumm = 1

func main() {
	rest := 0
	for i := 2; i <= 10000000; i++ {
		n := i
		count := 1
		for n > 1 {
			rest = n % 2
			if rest == 0 {
				n = n / 2
				count++

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
