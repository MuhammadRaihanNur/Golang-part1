package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Raihan"
	names[2] = "Nur"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		87,
		89,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)
}
