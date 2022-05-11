package main

import "fmt"

func main() {

	//message := "I will be passed to the sayMessage function as argument!"
	message := "Hey, whats up with Go?"
	sayMessage(message)
	//using a function which expects the multiple parameters
	lifeLine(88.88, 8, "Yuko")
	myName := "OBBO"
	days := 4
	myScore := 77.4
	lifeLine(myScore, days, myName)

}

func sayMessage(msg string) {
	fmt.Println("Message received....")
	fmt.Println(msg)
}

func lifeLine(score float64, duration int, name string) {
	fmt.Printf("%v should score %v in %v days \n", name, score, duration)
}
