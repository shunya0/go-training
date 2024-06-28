package main

import "fmt"

func main() {
	
	for num := 2; num <= 100; num++ {
		count := 0
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				count++
				break
			}
		}
		if count == 0 && num != 1 {
			fmt.Println(num)
		}
	}
}
