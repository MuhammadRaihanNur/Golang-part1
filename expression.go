package main

import "fmt"

func main() {
	name := "Raihan"

	if name == "Raihan" {
		fmt.Println("Hello Raihan")
	} else if name == "Budi" {
		fmt.Println("woyy gantengg!")
	} else if name == "Haji" {
		fmt.Println("woyy laah ganteng!")
	} else {
		fmt.Println("jiakhhh lucu u")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Sudah benar namamu")
	}
}
