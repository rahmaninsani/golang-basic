package main

import (
	"43-access-modifier/helper"
	"fmt"
)

func main() {
	helper.SayHello("Alice")
	helper.sayHi("Bob") // error - private function

	fmt.Println(helper.Application)
	fmt.Println(helper.version) // error - private variable
}
