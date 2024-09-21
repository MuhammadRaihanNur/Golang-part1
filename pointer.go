package main



type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1

	// address2.City = "Bandung"

	// fmt.Println(address1)
	// fmt.Println(address2)

	var address1 Address = Address{"Subang", "Jawa barat", "Indonesia"}
	var address2 *Address = &address1 //pointer

	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)
}
