package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	for i, arg := range args {
		fmt.Println(i+1, ":", arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname:", hostname)
	} else {
		fmt.Println("Error:", err)
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println("Username:", username)
	fmt.Println("Password:", password)

}
