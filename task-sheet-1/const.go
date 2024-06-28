package main

import "fmt"

func main() {
	const pi = 3.14
	const gravitationConst = 6.67430e-11
	//Creating multiple variable at once

	name, lastName, age, learningGO := "Suryansh", "Rohil", 18, true // Covers String , int , bool
	fmt.Println("Full name and age is: ", name, lastName, age)
	fmt.Println("Learning go? ", learningGO)
	fmt.Println("Value of pi is: ", pi) // covers constant and also float
	fmt.Println("Value of Gravitation Constant is: ", gravitationConst)
	var num float64 = 6.1234567
	fmt.Println(num) // << float64
}
