package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// Package container/ting adalah implementasi dari struktur data Circular List

	var data *ring.Ring = ring.New(5)

	//data.Value = "Alice"
	//var data2 = data.Next()
	//data2.Value = "Bob"

	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
