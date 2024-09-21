package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	var result any = random()
	// resultStirng := result.(string)
	// fmt.Println(resultStirng)
	// resultInt := result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}
