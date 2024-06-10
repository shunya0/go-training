package main

import "fmt"

type lib struct {
	book      string
	author    string
	ISBN      int
	available bool
}

func (book *lib) borrowBook() {

	book.available = false

}
func (book *lib) returnBook() {
	book.available = true
}
func (book lib) display() {
	fmt.Println("Book Name: ", book.book)
	fmt.Println("Book Author: ", book.author)
	fmt.Println("ISBN: ", book.ISBN)
	fmt.Println("Availability: ", book.available)
}

func newBook(book string, author string, ISBN int, available bool) *lib {
	return &lib{book: book, author: author, ISBN: ISBN, available: available}
}

func main() {
	bookForGo := newBook("Head First Go: A Brain-Friendly Guide ", "Jay McGavren", 9781491969557, true)
	bookForGo.display()
	fmt.Println(" ")

	bookForPython := newBook("Advanced Python Guide: Master concepts, build applications, and prepare for interviews", "Kriti Kumari Sinha", 9789355516756, true)
	bookForPython.display()
	fmt.Println(" ")

	fmt.Println("Book borrowed: Python & Go")
	fmt.Println(" ")

	bookForGo.borrowBook()
	bookForGo.display()

	fmt.Println(" ")

	bookForPython.borrowBook()
	bookForPython.display()
	fmt.Println(" ")

	fmt.Println("Book returned: go")
	bookForGo.returnBook()
	fmt.Println(" ")

	bookForGo.display()
	fmt.Println(" ")

	bookForPython.display()

}
