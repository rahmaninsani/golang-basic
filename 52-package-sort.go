package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type UserSlice []Student

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := []Student{
		{"Alice", 30},
		{"Bob", 20},
		{"Charlie", 25},
		{"Dennis", 40},
		{"Edith", 10},
	}
	fmt.Println(users)

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
