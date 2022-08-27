package main

import "fmt"

// func emptyInterfaceFunction(i int) any {}
func emptyInterfaceFunction(i int) interface{} {
	if i == 1 {
		return true
	} else if i == 2 {
		return 2
	} else {
		return "Ups"
	}
}

func main() {
	var result1 interface{} = emptyInterfaceFunction(1)
	fmt.Println(result1)

	var result2 = emptyInterfaceFunction(2)
	fmt.Println(result2)

	result3 := emptyInterfaceFunction(3)
	fmt.Println(result3)

	var result4 any = emptyInterfaceFunction(4)
	fmt.Println(result4)
}
