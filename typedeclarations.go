package main

import "fmt"

func main() {

	type NoKTP string

	var ktpRaihan NoKTP = "111111111"
	fmt.Println(ktpRaihan)
	fmt.Println(NoKTP("222222222"))
}
