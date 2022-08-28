package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Package container/list adalah implementasi dari struktur data Double Linked List

	data := list.New()
	data.PushBack("Alice")
	data.PushBack("Bob")
	data.PushBack("Charlie")
	data.PushFront("David")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	fmt.Println(data.Front().Next().Next().Value)
	fmt.Println(data.Back().Prev().Prev().Value)

	// Depan -> Belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// Belakang -> Depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
