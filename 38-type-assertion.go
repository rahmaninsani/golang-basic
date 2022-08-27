package main

import "fmt"

func random() interface{} {
	return "Bob"
}

func main() {
	result := random()

	//resultString := result.(string)
	//fmt.Println(resultString)
	//
	//resultInt := result.(int)
	//fmt.Println(resultInt) // panic

	switch value := result.(type) {
	case string:
		fmt.Println(value, "is string")
	case int:
		fmt.Println(value, "is integer")
	default:
		fmt.Println(value, "is unknown")
	}

}
