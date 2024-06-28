package main

import "fmt"

func main() {
	var word1 string
	fmt.Println("Enter 1st words: ")
	fmt.Scanln(&word1)
	var word2 string
	fmt.Println("Enter 1st words: ")
	fmt.Scanln(&word2)

	if len(word1) != len(word2) {
		fmt.Println("Not an anagrams")
	}

	var alphabetCounter [26]int
	for i := 0; i < len(word1); i++ {
		alphabetCounter[word1[i]-'a']++
		alphabetCounter[word2[i]-'a']--
	}

	for _, count := range alphabetCounter {
		if count != 0 {
			fmt.Println("The strings are not anagrams.")
			return
		}
	}

	fmt.Println("The strings are anagrams.")

}
