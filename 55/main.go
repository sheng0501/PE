package main

import (
	"fmt"
	"math"
)

func isP(number int) bool {
	return number == reverse(number)
}

func digit(n, i int) int {
	return n / int(math.Pow10(i-1)) % 10
}

func reverse(number int) int {
	a := int(math.Log10(float64(number))) + 1
	result := 0
	for i := 1; i <= a; i++ {
		result += digit(number, i) * int(math.Pow10(a-i))
	}
	return result
}

func isL(number int) bool {
	for i := 0; i < 50; i++ {
		number = number + reverse(number)
		if isP(number) {
			return false
		}
	}
	return true
}

func main() {
	count := 0
	for i := 0; i < 10000; i++ {
		if isL(i) {
			count++
		}
	}
	fmt.Println(count)
}
