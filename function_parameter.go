package main

import "fmt"

func sayHai(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Woyy", filteredName)
}

func spamFilter(name string) string {
	if name == "Babi" {
		return "..."
	} else {
		return name
	}
}
func main() {
	sayHai("Raihan", spamFilter)

	filter := spamFilter
	sayHai("Babi", filter)
}
