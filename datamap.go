package main

import "fmt"

func main() {

	person := map[string]string{
		"name":   "Muhammad",
		"addres": "Raihan",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["addres"])

	book := make(map[string]string)
	book["title"] = "Buku Chur-ros"
	book["author"] = "Muhammad Raihan"
	book["wrong"] = "Ups"

	delete(book, "wrong")

	fmt.Println(book)

}
