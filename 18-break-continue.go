package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 8 {
			fmt.Println("Break")
			break
		}

		if i%2 != 0 {
			continue
		}

		fmt.Println("Loop:", i)
	}
}
