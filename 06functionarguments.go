package main

import "fmt"

func main() {

	//message := "I will be passed to the sayMessage function as argument!"
	message := "Hey, whats up with Go?"
	sayMessage(message)
}

func sayMessage(msg string) {
	fmt.Println("Message received....")
	fmt.Println(msg)
}
