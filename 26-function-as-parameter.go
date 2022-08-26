package main

import (
	"fmt"
	"strings"
)

//func userRole(name string, userRoleFilter func(string) string) {
//	role := userRoleFilter(name)
//	fmt.Printf("%s - %s\n", name, role)
//}

type Filter func(string) string

func userRole(name string, userRoleFilter Filter) {
	role := userRoleFilter(name)
	fmt.Printf("%s - %s\n", name, role)
}

func userRoleFilter(name string) string {
	switch strings.ToLower(name) {
	case "alice":
		return "Production Manager"
	case "bob":
		return "Tech Manager"
	default:
		return "Staff"
	}
}

func main() {
	userRole("Alice", userRoleFilter)
	userRole("Bob", userRoleFilter)
	userRole("Charlie", userRoleFilter)
}
