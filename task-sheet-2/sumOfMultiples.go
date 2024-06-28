package main

import "fmt"

func main() {
	n := 1000
	result := 0
	for i := 0; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}
	fmt.Println(result)
}
