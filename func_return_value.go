package main

import "fmt"

func sayHello(name string) string {
	hello := "hello " + name
	return hello
}

func main() {
	result := sayHello("Raihan")
	fmt.Println(result)

	fmt.Println(sayHello("Bud"))
	fmt.Println(sayHello("Han"))
}
