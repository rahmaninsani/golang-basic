package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Pesan Error:", message)
	}

	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(false)
	runApp(true)
	fmt.Println("Hello")
}
