package main

import "fmt"

func endDefer() {
	fmt.Println("Selesai memanggil Function")
}

func runApplication(value int) {
	defer endDefer()
	result := 10 / value
	fmt.Println(result)
	fmt.Println("Run Application")
}

func main() {
	runApplication(10)
	runApplication(0)
}
