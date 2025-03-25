package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	fmt.Println("Hello, World!")
	message := greetings.Hello("Gladys") // Ensure the greetings package is correctly implemented
	fmt.Println(message)
}
