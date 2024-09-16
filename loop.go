package main

import (
	"fmt"
)

func main() {

	for counter := 90; counter <= 100; counter++ {
		fmt.Println("perulangan ke ", counter)
	}

	curros := []string{"Muhammad", "Raihan", "Nur"}
	for i := 0; i < len(curros); i++ {
		fmt.Print(curros[i])
	}

	for apa, iya := range curros {
		fmt.Println("yaila", apa, "=", iya)
	}
}
