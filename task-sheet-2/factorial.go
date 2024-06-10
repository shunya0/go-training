package main

import "fmt"

func main() {
	n := 0
	fmt.Println("Enter n: ")
	fmt.Scanln(&n)
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	fmt.Println(result)
}
