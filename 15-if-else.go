package main

import "fmt"

func main() {
	name := ""

	if name == "Bob" {
		fmt.Println("Hello, Bob!")
	} else if name == "Alice" {
		fmt.Println("Hello, Alice!")
	} else {
		fmt.Println("Hello, who are you?")
	}

	if length := len(name); length > 0 {
		fmt.Println("Your name has", length, "characters.")
	} else {
		fmt.Println("Your name has 0 characters.")
	}
}
