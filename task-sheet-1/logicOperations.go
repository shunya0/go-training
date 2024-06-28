package main

import "fmt"

func main() {
	shouldGo := true
	fmt.Println("Please enter your height (in cm): ")
	height := 0.0
	fmt.Scanln(&height)

	weight := 0.0
	isStaff := true

	fmt.Println("Are you are part of staff? (yes or no)")
	var staff string
	fmt.Scanln(&staff)
	if staff == "yes" {
		isStaff = true
	} else {
		isStaff = false
	}

	fmt.Println("Please enter your weight (in kgs): ")
	fmt.Scanln(&weight)
	if (height > 152.4 && weight < 80) || (isStaff) {
		shouldGo = true

	} else {
		shouldGo = false
	}

	if shouldGo {
		fmt.Println("Can go")
	} else {
		fmt.Println("Go back")
	}
}
