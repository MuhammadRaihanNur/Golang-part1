package main

import "fmt"

func getFullName() (string, string) {
	return "Muhamad", "Raihan"
}

func main() {
	// fristName, lastName := getFullName()
	// fmt.Println(fristName, lastName)
	fristName, _ := getFullName()
	fmt.Println(fristName)
}
