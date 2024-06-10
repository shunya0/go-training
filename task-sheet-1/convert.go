package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 52
	var floatNum float64 = 3.14

	floatInt := float64(number)
	intFloat := int(floatNum)
	strInt := strconv.Itoa(number)

	fmt.Println("float to int: ", floatInt)
	fmt.Println("int to float: ", intFloat)
	fmt.Println("int to string: ", strInt)
}
