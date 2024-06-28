package main

import "fmt"

type rectangle struct {
	length    int
	width     int
	AREA      int
	PREIMETER int
}

func (rect *rectangle) area() {

	rect.AREA = (rect.length * rect.width)
	fmt.Println("Area: ", rect.AREA)

}
func (rect *rectangle) premtr() {

	rect.PREIMETER = 2 * (rect.length + rect.width)
	fmt.Println("Perimeter: ", rect.PREIMETER)

}

func createRectangle(length int, width int) *rectangle {
	return &rectangle{length: length, width: width}
}

func (rect rectangle) displayRectangle() {
	fmt.Println("Length: ", rect.length)
	fmt.Println("Width: ", rect.width)
}

func main() {
	rectangle := createRectangle(15, 10)
	rectangle.displayRectangle()
	fmt.Println(" ")

	rectangle.area()
	rectangle.premtr()

}
