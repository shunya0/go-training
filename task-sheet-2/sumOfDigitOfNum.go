package main

import "fmt"

func main() {
	number := 0
	fmt.Println("Enter a number: ")
	fmt.Scanln(&number)

	digits := 0
	fmt.Println("Enter number of digit this number contains: ")
	fmt.Scanln(&digits)

	arr := make([]int, digits)

	for i := 0; i < digits; i++ {
		arr[i] = number % 10
		number = number / 10
	}
	sum := 0
	for a := 0; a < digits; a++ {
		sum += arr[a]
	}

	fmt.Println(sum)
}
