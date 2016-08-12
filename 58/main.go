package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	if n < 2 || n%2 == 0 {
		return false
	}
	if n < 9 {
		return true
	}
	if n%3 == 0 {
		return false
	}
	r := math.Sqrt(float64(n))
	f := 5
	for float64(f) <= r {
		if n%f == 0 {
			return false
		}
		if n%(f+2) == 0 {
			return false
		}
		f += 6
	}
	return true
}

func main() {
	//	diagonals := []int{}
	primeDiagonals := 0.0
	nonPrimeDiagonals := 1.0
	//	diagonals = append(diagonals, 1)
	n := 1
	ratio := 100.0
	for ratio > 0.1 {
		ur := 4*n*n - 2*n + 1
		if isPrime(ur) {
			fmt.Println(ur)
			primeDiagonals += 1
		} else {
			nonPrimeDiagonals += 1
		}
		ul := 4*n*n + 1
		if isPrime(ul) {
			primeDiagonals += 1
		} else {
			nonPrimeDiagonals += 1
		}
		bl := 4*n*n + 2*n + 1
		if isPrime(bl) {
			primeDiagonals += 1
		} else {
			nonPrimeDiagonals += 1
		}
		nonPrimeDiagonals += 1
		ratio = primeDiagonals / (nonPrimeDiagonals + primeDiagonals)
		fmt.Println(ratio * 10)
		fmt.Println(primeDiagonals, nonPrimeDiagonals)
		if primeDiagonals/(nonPrimeDiagonals+primeDiagonals) < 0.1 {
			fmt.Println("ANSWER FOUND")
			fmt.Println(n)
			break
		}
		n += 1
	}

}
