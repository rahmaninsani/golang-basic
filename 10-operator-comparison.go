package main

import "fmt"

func main() {
	var name1 string = "Alice"
	var name2 = "Bob"

	var isEqual bool = name1 == name2
	fmt.Println(isEqual)

	var firstNumber = 100
	var secondNumber = 200

	fmt.Println(firstNumber > secondNumber)
	fmt.Println(firstNumber < secondNumber)
	fmt.Println(firstNumber == secondNumber)
	fmt.Println(firstNumber != secondNumber)

}
