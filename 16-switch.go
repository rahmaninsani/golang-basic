package main

import "fmt"

func main() {
	var name string = ""

	switch name {
	case "Alice":
		fmt.Println("Hi, Alice!")
	case "Bob":
		fmt.Println("Hi, Bob!")
	default:
		fmt.Println("Hi, who are you?")
	}

	switch length := len(name); length > 0 {
	case true:
		fmt.Println("Your name has", length, "characters.")
	case false:
		fmt.Println("Your name has 0 characters.")
	}

	length := len(name)
	switch {
	case length < 0:
		fmt.Println("Your name has 0 characters.")
	case length > 10:
		fmt.Println("Your name is too long.")
	default:
		{
			fmt.Println("Welcome")
			fmt.Println("Your name has", length, "characters.")
		}
	}

}
