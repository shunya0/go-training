package main

import (
	"fmt"
	"sort"
	"strings"
)

// splites string in chars , sort chars and rejoins them
func sortStrings(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func groupAnagram(words []string) map[string][]string {
	anagramGroups := make(map[string][]string) // sorted string as key , values are chars

	for _, word := range words {
		sortedWords := sortStrings(word)
		anagramGroups[sortedWords] = append(anagramGroups[sortedWords], word)

	} // sorts characters of word and append them with original to the map
	return anagramGroups
}

func main() {
	words := []string{"listen", "lies", "hello", "world", "go", "language"}
	groupedAnagram := groupAnagram(words)

	for sortedWords, anagrams := range groupedAnagram {
		fmt.Println(sortedWords, anagrams)
	}
}
