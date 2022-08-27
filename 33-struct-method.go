package main

import "fmt"

type User struct {
	Name, Address string
	Age           int
}

// Struct Method
func (customer User) greeting() {
	fmt.Println("Hello", customer.Name)
}

func main() {
	user := User{Name: "John", Address: "123 Main St", Age: 30}
	user.greeting()
}
