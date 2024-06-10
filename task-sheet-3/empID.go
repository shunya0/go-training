package main

import "fmt"

type employ struct {
	name     string
	position string
	salary   float64
	id       string
}

func createEmploy(name string, position string, salary float64, id string) *employ {

	return &employ{name: name, position: position, salary: float64(salary), id: id}

}

// func raise(name string, position string, salary int, id string) *employ {

// }
func (person employ) displayEmploy() {

	fmt.Println("Name: ", person.name)
	fmt.Println("Position: ", person.position)
	fmt.Println("Salary: ", person.salary)
	fmt.Println("Id: ", person.id)

}
func (person *employ) raiseEmploy(raiseAmnt float64) {
	person.salary += raiseAmnt
}

func main() {
	person1 := createEmploy("Suryansh", "Intern", 100.84, "sr738")
	person1.displayEmploy()

	fmt.Println("\n")
	fmt.Println("After a raise")
	fmt.Println("\n")

	person1.raiseEmploy(1129.34)
	person1.displayEmploy()

}
