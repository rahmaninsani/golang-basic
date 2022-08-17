package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 2
	c := a * b
	fmt.Println(c)

	var d = 10
	d += 10
	fmt.Println(d)

	var f = 19
	f++
	fmt.Println(f)

	e := 21
	e--
	fmt.Println(e)

	var positiveNumber = +20
	fmt.Println(positiveNumber)

	negativeNumber := -positiveNumber
	fmt.Println(negativeNumber)

	var isLoading = true
	fmt.Println(isLoading)

	varIsNotLoading := !isLoading
	fmt.Println(varIsNotLoading)
}
