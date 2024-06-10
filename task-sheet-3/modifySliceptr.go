package main

import "fmt"

func modify[T any](arr *[]T, newValue T) {

	for i, _ := range *arr {
		(*arr)[i] = newValue
	}
	fmt.Println("\nModified")
	for i, _ := range *arr {
		fmt.Print((*arr)[i], " ")
	}
}

func main() {

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original")
	for i, _ := range arr {

		fmt.Print(arr[i], " ")

	}
	modify(&arr, 1)

}
