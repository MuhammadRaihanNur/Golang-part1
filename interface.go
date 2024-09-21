package main

import "fmt"

type Hasname interface {
	Getname() string
}

func sayHello(value Hasname) {
	fmt.Println("Hello", value.Getname())

}

type Person struct {
	Name string
}

func (person Person) Getname() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) Getname() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Raihan"}
	sayHello(person)

	animal := Animal{Name: "Kucing"}
	sayHello(animal)
}
