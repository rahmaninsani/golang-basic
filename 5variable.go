package main

import "fmt"

func main() {
	var firstName string
	firstName = "Rahman"
	fmt.Println(firstName)

	var lastName = "Insani"
	fmt.Println(lastName)

	var age int8 = 22
	fmt.Println(age)

	fullName := "Rahman Insani"
	fmt.Println(fullName)

	var (
		firstNumber               int8 = 1
		secondNumber                   = 2
		thirdNumber, fourthNumber int8 = 3, 4
	)

	fmt.Println(firstNumber)
	fmt.Println(secondNumber)
	fmt.Println(thirdNumber)
	fmt.Println(fourthNumber)

	a, b := 1, 2
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)

}
