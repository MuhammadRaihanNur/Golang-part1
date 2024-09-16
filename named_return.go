package main

import "fmt"

func namaLengkapYuk() (fristName, middleName, lastName string) {
	fristName = "Muhammad"
	middleName = "Raihan"
	lastName = "Nur"

	return fristName, middleName, lastName
}

func main() {
	a, b, c := namaLengkapYuk()
	fmt.Println(a, b, c)
}
