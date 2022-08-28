package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// address2 := address1 // address2 is a copy of address1
	var address2 *Address = &address1 // address2 is a pointer to address1

	fmt.Println("\nReference Pertama")

	address2.City = "Bandung"
	address2.Province = "Jawa Barat"

	fmt.Println("address1:", address1)
	fmt.Println("address2:", *address2)

	fmt.Println("\nReference Setelah address2 di assign ulang dengan *")

	//address2 = &Address{"Surabaya", "Jawa Timur", "Indonesia"} // hanya mengubah dirinya sendiri (&)
	address3 := &address1
	*address2 = Address{"Surabaya", "Jawa Timur", "Indonesia"} //mengubah seluruh variabel yang reference ke address1 (*)

	//Semua variable yg reference ke address1 yakni address2 dan address3 isinya sama
	fmt.Println("address1:", address1)
	fmt.Println("address2:", *address2)
	fmt.Println("address3:", *address3)

	fmt.Println("\nPointer Baru")
	var address4 *Address = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println("address4:", address4)

	fmt.Println("\nPointer dengan Keyword new sehingga default value-nya kosong")
	var address5 *Address = new(Address)
	address5.City = "Jakarta"
	fmt.Println("address5:", address5)
}
