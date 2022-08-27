package main

import "fmt"

type HashName interface {
	GetName() string
}

func sayHi(hasName HashName) {
	fmt.Println("Hi", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "John"}
	sayHi(person)

	animal := Animal{Name: "Cat"}
	sayHi(animal)
}
