package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	bob := Man{"Bob"}
	fmt.Println(bob.Name) // Bob

	bob.Married()
	fmt.Println(bob.Name) // Mr. Bob
}
