package main

import "fmt"

func main() {

	counter := 1

	for counter <= 5 {
		fmt.Println("Loop:", counter)
		counter++
	}

	for counter := 1; counter <= 5; counter++ {
		fmt.Println("Loop:", counter)
	}

	users := []string{"John", "Paul", "George", "Ringo"}
	for _, user := range users {
		fmt.Println(user)
	}

	books := map[string]string{"The Hobbit": "J.R.R. Tolkien", "The Lord of the Rings": "J.R.R. Tolkien"}
	for title, author := range books {
		fmt.Println(title, "by", author)
	}

}
