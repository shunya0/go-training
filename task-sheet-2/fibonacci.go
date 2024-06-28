package main

import "fmt"

func main() {

	n := 0
	fmt.Println("Enter n: ")
	fmt.Scanln(&n)
	series := make([]int, n)

	series[0] = 0
	series[1] = 1

	for i := 2; i < n; i++ {
		series[i] = series[i-1] + series[i-2]
	}
	for i := 0; i <= n-1; i++ {
		fmt.Println(series[i])
	}
	//F(n) = F(n-1) + F(n-2)
}
