package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr.  " + man.Name
}

func main() {
	raihan := Man{"Raihan"}
	raihan.Married()

	fmt.Println(raihan)
}
