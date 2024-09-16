package main

import "fmt"

func main() {
	// Variable dengan kata kunci Var
	// var name string

	// name = "Muhammad Raihan"
	// fmt.Println(name)

	// name = "Raihan Nur"
	// fmt.Println(name)

	// Saat kita membuat variable, maka kita wajib menyebutkan tipe data variable tersebut
	//Namun jika kita langsung menginisialisasikan data pada variable nya, maka kita tidak wajib menyebutkan tipe data variable nya

	// var name = "Muhammad Raihan"
	// fmt.Println(name)

	// name = "Raihan Nur"
	// fmt.Println(name)

	//kata kunci var :=
	// name := "Muhammad Raihan"
	// fmt.Println(name)

	// name = "Raihan Nur"
	// fmt.Println(name)

	// Deklarasi Multiple Variable
	var (
		fristName = "Muhammad Raihan"
		lastName  = "Nur Rizqi Amin"
	)
	fmt.Println(fristName)
	fmt.Println(lastName)
}
