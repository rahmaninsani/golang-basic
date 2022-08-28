package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("\nBob Marley", "Marley"))
	fmt.Println(strings.Contains("Bob Marley", "Alice"))

	fmt.Println(strings.Split("\nBob Marley", " "))
	fmt.Println(strings.Split("Bob Marley", " ")[0])

	fmt.Println(strings.ToLower("\nBob Marley"))
	fmt.Println(strings.ToUpper("Bob Marley"))
	fmt.Println(strings.ToTitle("Bob Marley"))

	fmt.Println(strings.Trim("      Bob     ", " "))
	fmt.Println(strings.Trim("a      Bob     a", " "))

	fmt.Println(strings.ReplaceAll("Bob Marley Bob", "Bob", "Alice"))
	fmt.Println(strings.Replace("Bob Marley Bob", "Bob", "Alice", 1))

}
