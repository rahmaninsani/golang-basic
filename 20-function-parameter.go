package main

import "fmt"

func sayHelloParam(index int, fullName string) {
	fmt.Printf("User-%d %s\n", index, fullName)
}

func main() {
	users := []map[string]string{
		{"fullName": "Alice"},
		{"fullName": "Bob"},
		{"fullName": "Dylan"},
	}

	for index, user := range users {
		sayHelloParam(index+1, user["fullName"])
	}

}
