package main

import (
	"errors"
	"fmt"
)

func divide(value int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("divisor cannot be zero")
	}

	return value / divisor, nil
}

func main() {
	var errorExample error = errors.New("error example")
	fmt.Println(errorExample.Error())

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
