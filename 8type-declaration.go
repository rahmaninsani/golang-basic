package main

import "fmt"

func main() {
	type NIK string
	var myNIK NIK = "12345"
	fmt.Println(myNIK)
	fmt.Println(NIK("12345"))

	type Married bool
	var marriedStatus Married = false
	fmt.Println(marriedStatus)

}
