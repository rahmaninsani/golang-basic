package main

import "fmt"

func main() {
	const firstName string = "Rahman"
	fmt.Println(firstName)

	const lastName = "Insani"
	fmt.Println(lastName)

	const (
		a    = 1
		b    = 2
		c, d = 3, 4
	)

}
