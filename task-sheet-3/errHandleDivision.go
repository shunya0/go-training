package main

import (
	"errors"
	"fmt"
)

func div(dividend int, divisor int) (int, error) {

	if divisor == 0 {
		return 0, errors.New("divisor can't be zero")
	}
	return (dividend / divisor), nil

}

func main() {
	divisor, dividend := 0, 4

	result, err := div(dividend, divisor)

	if err != nil {
		fmt.Println("Error ", err)
	} else {
		fmt.Printf("Division of %d/%d = %d", dividend, divisor, result)

	}

}
