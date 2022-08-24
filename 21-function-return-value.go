package main

import "fmt"

func addTwoNumber(firstNumber int, secondNumber int) int {
	if firstNumber == 0 {
		return secondNumber
	} else if secondNumber == 0 {
		return firstNumber
	} else {
		return firstNumber + secondNumber
	}
}

func main() {
	result := addTwoNumber(10, 90)
	fmt.Println(result)
}
