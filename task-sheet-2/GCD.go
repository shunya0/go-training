package main

import "fmt"

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	a := 0
	b := 0
	fmt.Println("Enter two number u wanna find GCD of: ")
	fmt.Scanln(&a, &b)
	// fmt.Scanln(&b)

	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	if b > a {
		c := a
		a = b
		b = c
	}
	GCD := gcd(a, b)
	fmt.Println("GCD is: ", GCD)
}
