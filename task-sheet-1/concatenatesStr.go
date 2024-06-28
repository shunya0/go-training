package main

import "fmt"

func main() {
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
}
