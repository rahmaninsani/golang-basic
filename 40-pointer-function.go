package main

import "fmt"

type Address2 struct {
	Country string
}

func changeAddress(address *Address2, value string) {
	address.Country = value
}

func changeNumber(number *int, value int) {
	*number = value
}

func main() {
	var address1 = Address2{}
	changeAddress(&address1, "Indonesia")

	var address2 = &Address2{}
	changeAddress(address2, "China")

	fmt.Println(address1.Country)
	fmt.Println(address2.Country)

	numberA := 10
	fmt.Println("\nbefore:", numberA)

	changeNumber(&numberA, 20)
	fmt.Println("after:", numberA)

}
