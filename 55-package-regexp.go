package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("b([a-z])b")

	fmt.Println(regex.MatchString("bob"))
	fmt.Println(regex.MatchString("bab"))
	fmt.Println(regex.MatchString("Bob"))

	search := regex.FindAllString("bob bab bib bub Bob", 1)
	fmt.Println(search)

	search = regex.FindAllString("bob bab bib bub Bob", -1)
	fmt.Println(search)
}
