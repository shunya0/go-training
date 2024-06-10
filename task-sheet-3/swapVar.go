package main

import "fmt"

func swap(a *int, b *int) (int, int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
	return *a, *b
}

func main() {
	a, b := 5, 10

	fmt.Printf("Before swap \n a: %d \n b: %d \n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap \n a: %d \n b: %d ", a, b)

}
