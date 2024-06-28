package main

import "fmt"

func median(arr []int) {

	n := len(arr)

	if n%2 == 0 {

		fmt.Println(arr[(n/2)-1] + arr[(n/2)])
	} else {
		fmt.Println(arr[n/2])
	}

}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	median(arr)
}
