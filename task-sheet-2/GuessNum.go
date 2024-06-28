package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(100)
	// fmt.Println(n)
	for i := 0; i <= 10; i++ {
		a := 0
		fmt.Println("Guess the number b/w 1 to 100")
		fmt.Scanln(&a)

		if a < n {
			fmt.Println("Too low")
		} else {
			if a > n {
				fmt.Println("Too far")

			} else {
				if a == n {
					fmt.Println("you guessed it correct!")
					break
				}
			}
		}
	}
}
