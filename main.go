// Solution to problem : https://projecteuler.net/problem=37
/* The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right,
   and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.
   Find the sum of the only eleven primes that are both truncatable from left to right and right to left.
   NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes. */

package main

import (
	"fmt"
	"strconv"
)

// Prime function return true if it's a prime number
func prime(number int) bool {
	nbDiv := 0
	for div := 2; div <= number; div++ {
		if number%div == 0 {
			nbDiv++
			if nbDiv > 1 {
				return false
			}
		}
	}
	return true
}

// truncate Prime function return true if it's a truncate prime
func truncatePrime(number int) bool {
	strings := string(strconv.Itoa(number))
	// Loop start to 1 to len of the number
	for digit := 1; digit <= len(strings); digit++ {
		// Select the X digits by right and by left
		right, _ := strconv.Atoi(strings[0:digit])
		left, _ := strconv.Atoi(strings[digit-1:])
		// if  right of left numbers not a prime number we quit the function
		if !prime(left) || !prime(right) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Début")
	counter := 0
	// loop start to 8 for avoid the Prime number before 8
	for i := 8; i < 100000; i++ {
		// If the number it's prime number we going to analyse if it's a truncate prime
		if prime(i) {
			if truncatePrime(i) {
				fmt.Println(i)
				counter++
			}
		}
		if counter >= 23 {
			break
		}
	}
}
