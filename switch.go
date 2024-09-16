package main

import (
	"fmt"
)

func main() {

	name := "Muhammad Raihan Nur"

	switch name {
	case "Raihan":
		fmt.Println("Hello Raihan")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hi, ganteng banget si")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Cocok namamu")
	}

	length := len(name)
	switch {
	case length > 11:
		fmt.Println("panjang banget")
	case length > 5:
		fmt.Println("mayan panjang bro")
	default:
		fmt.Println("nah, sudah bener namau=mu")

	}
}
