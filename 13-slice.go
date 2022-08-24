package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println("Len:", len(slice1))
	fmt.Println("Cap:", cap(slice1))

	months[5] = "Changed-June"
	fmt.Println(slice1)

	slice1[0] = "Changed-May"
	fmt.Println(months)

	slice2 := months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "New Month")
	fmt.Println(slice3)
	slice3[1] = "Changed-December"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Alice"
	newSlice[1] = "Bob"

	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Print(copySlice)

}
