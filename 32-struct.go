package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

type CustomerPremium struct {
	//CS Customer
	Customer
	Gift string
}

func main() {
	var customer1 Customer
	//var customer1 = Customer{}
	customer1.Name = "Bob"
	customer1.Address = "123 Main St."
	customer1.Age = 36

	fmt.Println(customer1.Name)
	fmt.Println(customer1.Address)
	fmt.Println(customer1.Age)

	var customer2 Customer = Customer{
		Name:    "Alice",
		Address: "321 North St.",
		Age:     40,
	}
	fmt.Println(customer2)

	customer3 := Customer{"Carol", "503 South St.", 50}
	fmt.Println(customer3)

	customerPremium := CustomerPremium{
		Customer: Customer{
			Name:    "Dave",
			Address: "123 Main St.",
			Age:     36,
		},
		Gift: "A million dollars",
	}
	fmt.Println(customerPremium)

	var customerPremium2 = CustomerPremium{}
	customerPremium2.Name = "Eve"
	customerPremium2.Address = "321 North St."
	customerPremium2.Age = 40
	customerPremium2.Gift = "A million dollars"
	fmt.Println(customerPremium2.Customer.Name, "- Gift:", customerPremium2.Gift)

	var userPremium = CustomerPremium{Customer: customer3, Gift: "MacBook Pro M1"}
	fmt.Println(userPremium.Name, "- Gift:", userPremium.Gift)

	//	Anonymous Struct
	// var basicUser struct {}
	// basicUser := struct {}
	var basicUser = struct {
		Customer
		Point int
	}{Customer: customer3, Point: 42}
	fmt.Println(basicUser)

}
