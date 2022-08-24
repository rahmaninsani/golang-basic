package main

import "fmt"

func getCompleteName() (firstName, lastName string) {
	firstName = "John"
	lastName = "Doe"
	return
}

func main() {
	_, lastName := getCompleteName()
	fmt.Println(lastName)
}
