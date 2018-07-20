package main

import (
	"fmt"
	"log"
	"time"
)

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
	start := time.Now()
	for i := 2; i <= 10000000; i++ {
		n := i
		count := 1
		/*
		  runtime improvements can be done here:
		  1. Discard processing base 2 numbers -> never going to be on top
		  2. By observation 1: Top 10 come from original odd numbers. Discard processing original even numbers can be done?
		  3. By observation 2: Greater values are always create by the second half (Ex: i=10000, can start with i>500. reduce 50% )
		*/

		for n >= 2 {
			/*
				Just count dont care about the slices values
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
	fmt.Println("\nInformational only:")
	fmt.Println("-------------------")
	fmt.Println("Biggest numer:", iB)
	fmt.Println("Biggest number length:", countB)
	fmt.Println("Second biggest numer:", iSB)
	fmt.Println("Second biggest number length:", countSB, "\n")
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
	fmt.Println("Summ of slices for the second biggest length:", ISBsumm, "\n")
	elapsed := time.Since(start)
	log.Printf("Runtime %s", elapsed)
}
