package main

import (
	"fmt"
)

const (
	fizzNumber = 3
	buzzNumber = 5
)

func isFizz(n int) bool {
	return n%fizzNumber == 0
}

func isBuzz(n int) bool {
	return n%buzzNumber == 0
}

func Convert(n int) string {
	if isFizz(n) && isBuzz(n) {
		return "FizzBuzz"
	} else if isFizz(n) {
		return "Fizz"
	} else if isBuzz(n) {
		return "Buzz"
	} else {
		return ""
	}
}

func main() {
	fmt.Println(Convert(3))  // Output: Fizz
	fmt.Println(Convert(5))  // Output: Buzz
	fmt.Println(Convert(15)) // Output: FizzBuzz
	fmt.Println(Convert(7))  // Output: ""
}
