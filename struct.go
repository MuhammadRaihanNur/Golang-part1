package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("hello", name, "my name is", customer.Name)
}
func main() {
	var raihan Customer
	raihan.Name = "Muhammad Raihan"
	raihan.Address = "Jakarta"
	raihan.Age = 22
	fmt.Println(raihan)

	dinda := Customer{
		Name:    "Dinda",
		Address: "Jakarta",
		Age:     21,
	}
	fmt.Println(dinda)

	budi := Customer{"Budi", "Jakarta", 21}
	fmt.Println(budi)

	budi.sayHello("agus")
	dinda.sayHello("agus")
	raihan.sayHello("agus")
}
