package main

import "fmt"

func dadahYes(name string) string {
	return "Good Bye " + name
}

func main() {
	nyoba := dadahYes
	jiah := dadahYes

	fmt.Println(nyoba("Budi"))
	fmt.Println(jiah("Raihan"))
}
