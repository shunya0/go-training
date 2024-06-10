package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func add(x, y int) int {
	return x + y
}

func main() {

	fmt.Println("Hello, world")
	// Printing Hello, World
	fmt.Println("Time right now is", time.Now()) // print current time

	fmt.Println("This number is a random number > ", rand.Intn(20)) // printing random number in a range of 20

	fmt.Println("Adding two number(4 and 5): ", add(4, 5))

	//Creating string variable
	var name string = "Suryansh"

	// var name string = "Suryansh"
	fmt.Println(name)

	//Declaring a constant
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

	num1 := 2
	var num2 float64 = 4.2
	sum := float64(num1) + num2

	fmt.Println("Addition of 2 nums (3 , 4): ", sum)
	fmt.Println("Subtraction of 2 nums (3 , 4): ", float64(num1)-num2)
	fmt.Println("Multiplication of 2 nums (3 , 4): ", float64(num1)*num2)
	fmt.Println("Division of 2 nums (3 , 4): ", float64(num1)/num2)

	//Conversions

	number := 52
	// fulLName := "SuryanshRohil"
	var floatNum float64 = 3.14

	floatInt := float64(number)
	intFloat := int(floatNum)
	strInt := strconv.Itoa(number)

	fmt.Println("float to int: ", floatInt)
	fmt.Println("int to float: ", intFloat)
	fmt.Println("int to string: ", strInt)

	//Using Boolean variable to perform logical operators

	shouldGo := true
	fmt.Println("Please enter your height (in cm): ")
	height := 0.0
	fmt.Scanln(&height)

	weight := 0.0
	isStaff := true

	fmt.Println("Are you are part of staff? (yes or no)")
	var staff string
	fmt.Scanln(&staff)
	if staff == "yes" {
		isStaff = true
	} else {
		isStaff = false
	}

	fmt.Println("Please enter your weight (in kgs): ")
	fmt.Scanln(&weight)
	if height > 152.4 && weight < 80 || isStaff {
		shouldGo = true

	} else {
		shouldGo = false
	}

	if shouldGo {
		fmt.Println("Can go")
	} else {
		fmt.Println("Go back")
	}

	//concatenating strings

	string1 := "hello"
	string2 := "world"
	fullString := string1 + string2
	fmt.Println(fullString)
	stringLength := 0
	// stringLength := len(fullString) finding length using library
	for range fullString {
		stringLength++
	}
	fmt.Println("length of the string is: ", stringLength)
	//max&min Values of each variable

	var int64Max int64 = math.MaxInt64
	var int64Min int64 = math.MinInt64
	var int32Max int64 = math.MaxInt32
	var int32Min int64 = math.MinInt32

	var float32Max float32 = math.MaxFloat32
	var float32Min float32 = -math.MaxFloat32
	var float64Max float64 = math.MaxFloat64
	var float64Min float64 = -math.MaxFloat64

	fmt.Println("int64 Max:", int64Max)
	fmt.Println("int64 Max:", int64Min)
	fmt.Println("int32 Max:", int32Min)
	fmt.Println("int32 Max:", int32Max)
	fmt.Println("float32 Max:", float32Max)
	fmt.Println("float32 Min:", float32Min)
	fmt.Println("float64 Max:", float64Max)
	fmt.Println("float64 Min:", float64Min)

}
