package main

import "fmt"

func collatz(n int) int {
	steps := 0
	if n == 1 {
		return steps
	}

	if n%2 == 0 {
		n /= 2
		steps++
	} else {
		n = (n * 3) + 1
		steps++
	}

	return steps + collatz(n)

}

func main() {
	n := 0
	fmt.Println("Enter n: ")
	fmt.Scanln(&n)

	stepsToReachOne := collatz(n)

	fmt.Println(stepsToReachOne)
}
