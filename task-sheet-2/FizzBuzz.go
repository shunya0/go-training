package main

import "fmt"

func main() {

	fmt.Println(0)
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			if i%5 == 0 {
				fmt.Println("Buzz")
			} else {
				if i%3 == 0 {
					fmt.Println("Fizz")
				} else {
					fmt.Println(i)
				}
			}
		}

	}
}
