package main

import "fmt"

func main() {
	var nilai32 int32 = 12
	fmt.Println(nilai32)

	var nilai64 int64 = int64(nilai32)
	fmt.Println(nilai64)

	var nilai16 int16 = int16(nilai32)
	fmt.Println(nilai16)

	var fullName = "Rahman Insani"
	//var charR uint8 = fullName[0]
	var charR byte = fullName[0]
	var charRString string = string(charR)
	fmt.Println(charRString)

}
