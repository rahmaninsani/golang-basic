package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumAll(40 + 30 + 20 + 10)
	fmt.Println(total)

	slice := []int{40, 30, 20, 10}
	total = sumAll(slice...)
	fmt.Println(total)
}
