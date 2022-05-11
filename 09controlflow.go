package main

import "fmt"

func main() {

	age := 30
	healthScore := 7
	flexibility := 3
	fmt.Println(age >= 30)

	//the block in if condition will be executed only if condition
	//returns true (condition is true)
	if age >= 30 {
		fmt.Println("You are too old to climb the mountain!")
	}

	if healthScore >= 8 {
		fmt.Println("Your fitness level is excellent.")
	} else {
		fmt.Println("You need to improve your fitness level.")
	}

	if flexibility >= 8 {
		fmt.Println("You are superbly flexible!")
	} else if flexibility >= 5 && flexibility < 8 {
		fmt.Println("You are flexible but can improve on it!")
	} else if flexibility < 5 && flexibility >= 3 {
		fmt.Println("You are less flexible. Do some yoga! ")
	} else {
		fmt.Println("You are rigid stick! Do not break please!")
	}

}
