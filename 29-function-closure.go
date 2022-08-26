package main

import "fmt"

func main() {
	counter1 := 0
	counter2 := 0

	increment := func() {
		fmt.Println("Increment")

		counter1++ // global counter1

		counter2 := 0 // redeclare local counter2
		counter2++    // local counter2
	}

	increment()

	fmt.Println(counter1)
	fmt.Println(counter2)

}
