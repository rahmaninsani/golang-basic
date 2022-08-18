package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Alice"
	names[1] = "Bob"
	names[2] = "Eve"

	//fmt.Println(names[0])
	//fmt.Println(names[1])
	//fmt.Println(names[2])

	for _, name := range names {
		fmt.Println(name)
	}

	values := [3]int{
		1,
		2,
		3,
	}

	fmt.Println("Values Length:", len(values))

	for _, value := range values {
		fmt.Println(value)
	}
}
