package main

import "fmt"

func main() {
	var word string
	fmt.Println("Word: ")
	fmt.Scanln(&word)
	arr := make([]byte, len(word))

	for i := 0; i <= len(word)-1; i++ {
		arr[i] = word[i]
	}
	// for a := 0; a < len(word); a++ {
	// 	fmt.Printf("%c", arr[a])
	// }
	flag := 0
	for s := 0 ; s <= len(word)/2 ; s++{
		if(arr[s] != arr[len(word) -1 - s]){
			flag++
		}
	}
	if flag != 0 {
		fmt.Println("Not a Planidrome")
	} else {
		fmt.Println("Planidrome")
	}
}
