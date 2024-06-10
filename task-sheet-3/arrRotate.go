package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}
	k := 2
	arrRev(arr, k)

}

func arrRev(arr []int, k int) {

	fmt.Println("Current array: ")
	fmt.Println("[")
	for i, _ := range arr {
		fmt.Println(arr[i])
	}
	fmt.Println("]")

	fmt.Printf("Array after %d rotation", k)

	n := len(arr)
	// for i := k; i <= n-1; i++ {
	// 	//k =2 , arr[2] = 3
	// 	if i != n-1 {
	// 		fmt.Println(arr[i]) // 4 ,
	// 	} else {
	// 		fmt.Println(arr[i+1]) //5
	// 		i = 0                 //arr[0] = 1 , k =2
	// 		if i != k {
	// 			fmt.Println(arr[i]) // 1 ,2
	// 			i++
	// 		} else {
	// 			fmt.Println(arr[i])
	// 			break
	// 		}
	// 	}

	// }

	for i := k; i <= n-1; i++ {
		fmt.Println(arr[i])

	}

	for i := 0; i <= k; i++ {
		fmt.Println(arr[i])
	}



}
