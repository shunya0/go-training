package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	char  byte
	freq  int
	left  *Node
	right *Node
}
type byFreq []*Node

func (h byFreq) Less(i int, j int) bool {
	return h[i].freq > h[j].freq
}
func (h byFreq) Len() int {
	return len(h)
}
func (h byFreq) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func push(stack []*Node, element *Node) {
	stack = append(stack, element)
}
func Pop(stack []*Node) (*Node, bool) {
	if len(stack) == 0 {
		fmt.Println("stack is empty")
		return nil, false
	}

	prev := len(stack) - 1
	popped := stack[prev]
	stack = stack[:prev]
	return popped, true

}

func makeTree(charFq map[byte]int) *Node {
	pq := make(byFreq, 0, len(charFq))
	for char, fq := range charFq {
		pq = append(pq, &Node{char: char, freq: fq})
	}

	for len(pq) > 1 {

	}

}
func readFileAndCompress(fileName string) ([]byte, error) {

	inputFile, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error while importing file: ", err)
		return nil, err
	}
	defer inputFile.Close()

	reader := bufio.NewReader(inputFile)
	buf := make([]byte, 4096)
	data, err := reader.Read(buf)

	for {
		if err != nil {
			fmt.Println("Error while reading , ", err)
			return nil, err
		}
		if data == 0 {
			break
		}
	}
	charFq := make(map[byte]int)
	for _, char := range data {
		charFq[char] += 1
	}

	return nil, err
}

func main() {
	dirPath := "C:\\Users\\surya\\OneDrive\\Desktop\\Files\\Codes\\Go\\task-sheet-4\\SampleDir"

	dir, err := os.Open(dirPath)
	if err != nil {
		fmt.Println("Error while opening file:  ", err)
	}
	defer dir.Close()
	entries, err := dir.Readdir(0)
	if err != nil {
		fmt.Println("Error while reading dir ", err)

	}
	for _, v := range entries {

		go readFile(dirPath + "\\" + v.Name())
		fmt.Println(v.Name() + " Successfully compressed")

	}

}
