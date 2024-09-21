package main

import (
	"fmt"
)

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("wah panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Maaf sedang ada gangguan")
	}
}

func main() {
	runApp(true)
	fmt.Println("okeee bisa lanjut")
}
