package main

import "fmt"

func main() {
	welcome()
	welcome()
	welcome()
	welcome()

}

func welcome() {
	fmt.Println("--------------------------------------")
	fmt.Println("Welcome to functions in Go!!!!")

	sayMessage()
	fmt.Println("--------------------------------------")
}

func sayMessage() {
	fmt.Println("Functions are easy to understand!")
	fmt.Println("Functions help in code re-usability!")
}
