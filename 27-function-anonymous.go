package main

import (
	"fmt"
	"strings"
)

type UserRoleFilter func(string) bool

func authorization(name string, userRoleFilter UserRoleFilter) {
	if userRoleFilter(name) {
		fmt.Println("Welcome,", strings.ToUpper(name))
	} else {
		fmt.Println("Access denied")
	}
}

func main() {
	isAdmin := func(name string) bool {
		return strings.ToLower(name) == "admin"
	}

	authorization("Admin", isAdmin)
	authorization("User", isAdmin)

	authorization("User", func(name string) bool {
		return strings.ToLower(name) == "user"
	})

}
