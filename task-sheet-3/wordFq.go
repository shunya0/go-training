package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wrdMap(text string) map[string]int {

	wordMap := make(map[string]int) // hold frequency of each word

	isLetter := func(r rune) bool {
		// fmt.Println(r)
		return unicode.IsLetter(r) // checking if it is a letter or a digit

	}

	words := strings.FieldsFunc(text, func(r rune) bool { //splits text into words to add to map
		// fmt.Println(r)
		return !isLetter(r) // ignoring punctuation
	})

	for _, word := range words { // converting words to lowercase and update count in map
		normalizeWord := strings.ToLower(word)
		wordMap[normalizeWord]++
	}

	return wordMap
}

func main() {
	//  (var name) := map[(type)](value){
	//  "(key)": (value),
	//  "(key)": (value),
	// }

	//fmt.Println(var name)
	//fmt.Println(var name [type]) // for a certain value

	str := "This is ! ? , 1 "
	fq := wrdMap(str)

	for word, count := range fq {
		fmt.Println(word, count)
	}

}
