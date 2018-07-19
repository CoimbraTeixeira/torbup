/*
Task:
For any positive integer n we define two rules:
if even: divide by two
if odd: multiply by three, then add one, and repeat until the result is the number 1

This will generate sequences of numbers like below, converging to 1:

3, 10, 5, 16, 8, 4, 2, 1

7, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1

For each number n we can now count the number of steps in this sequence until we reach 1.

So the sequence above, starting with 3, has a length of 8 (including the starting point and the final one).
The second sequence has a length of 17.
*/
package main

import "fmt"

var rest = 0
var count = 0
var i = 0
var n = 0
var countB = 0
var countSB = 0
var iB = 0
var iSB = 0

func main() {

	for i := 2; i <= 10000000; i++ 
	{
		n := i
		count := 1

		for n >= 2 
		{

			rest = n % 2
			if rest == 0 
			{
				n = n / 2
				count++

			} else 
				  {
				  n = n*3 + 1
				  count++
			      }

		}

		fmt.Println(i, count)
		if	count > countSB
		    {
			countSB := countB 
			countB := count
			iSB := iB
			iB := i
			}

		

			
		
	}

}
