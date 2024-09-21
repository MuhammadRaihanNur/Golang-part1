package main

import "fmt"

func faktorloop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}
func faktormudah(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * faktormudah(value-1)
	}
}
func main() {
	//perulangan dikali dari 10 hingga 1 berurutan
	fmt.Println(faktorloop(10))
	//perulangan namun lebih sederhana
	fmt.Println(faktormudah(10))
}
