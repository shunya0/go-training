package main

import "fmt"

func main() {
	num1 := 2
	var num2 float64 = 4.2
	sum := float64(num1) + num2

	fmt.Println("Addition of 2 nums (3 , 4): ", sum)
	fmt.Println("Subtraction of 2 nums (3 , 4): ", float64(num1)-num2)
	fmt.Println("Multiplication of 2 nums (3 , 4): ", float64(num1)*num2)
	fmt.Println("Division of 2 nums (3 , 4): ", float64(num1)/num2)
}
