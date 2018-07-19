package main

import "fmt"

var rest = 0
var count = 0
var countB = 0
var countSB = 0
var i = 0
var iB = 0
var iSB = 0
var ISBsumm = 1
var n = 0

func main() {

	for i := 2; i <= 10000000; i++ {
		n := i
		count := 1

		for n >= 2 {
			/*
				Just count to find the first and second biggest ones
			*/
			rest = n % 2
			if rest == 0 {
				n = n / 2
				count++

			} else {
				n = n*3 + 1
				count++

			}
		}
		/*
			Debug purpose
			fmt.Println(i, count)
		*/
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
	fmt.Println(iB, countB, iSB, countSB)
	/*
		Looking and summ all values for the second biggest one
	*/
	for iSB > 1 {
		ISBsumm += iSB
		rest = iSB % 2
		if rest == 0 {
			iSB = iSB / 2

		} else {
			iSB = iSB*3 + 1

		}

	}
	fmt.Println("Summ of slices for the second big length:", ISBsumm)
}
