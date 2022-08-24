package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Alice",
		"city": "Bandung",
	}

	person["email"] = "alice@example.com"

	fmt.Println(person)
	fmt.Println(len(person))
	fmt.Println(person["name"])
	fmt.Println(person["city"])
	fmt.Println(person["email"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Learn Golang"
	book["wrong"] = "Wrong"
	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "wrong")
	fmt.Println(book)
	fmt.Println(len(book))

}
