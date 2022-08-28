package helper

import "fmt"

/*
SayHello is Exported Function (Public) - PascalCase
*/
func SayHello(name string) {
	fmt.Println("Hello", name)
}

/*
sayHi is Unexported Function (Private) - camelCase
*/
func sayHi(name string) {
	fmt.Println("Hi", name)
}

/*
Application is Exported Variable (Public) - PascalCase
*/
var Application string = "Golang Basic"

/*
version is Unexported Variable (Private) - camelCase
*/
var version string = "1"
