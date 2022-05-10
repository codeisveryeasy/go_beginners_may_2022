package main

import "fmt"

func main() {
	var name string
	var score int
	var temperature float32

	name = "Air Asia Academy"
	score = 88
	temperature = 37.4

	fmt.Println(name)
	fmt.Println(score)

	fmt.Println("I can learn Go programming from " + name)
	fmt.Println("I can learn Go programming from", name)

	//single line comment
	/*
		this is
		multi line
		comment
	*/

	//using format specifiers -> %v
	fmt.Printf("I am learning Go progrmaming from %v. \n", name)
	fmt.Printf("My score is  %v \n", score)

	fmt.Printf("I am beginner in %v. I did good practice. My score in assessment was %v. \n", name, score)

	//format specifier -> %T
	fmt.Printf("Datatype of %v (temparature variable) is %T \n", temperature, temperature)
	fmt.Printf("Datatype of %v (name variable) is %T \n", name, name)
	fmt.Printf("Datatype of %v (score variable) is %T \n", score, score)

	//auto type inference (declare and initilize the variable and let go lang detect the datatype)
	academy := "AAA"
	fmt.Println("Shortform of Air Asia Academy is", academy)
	fmt.Printf("Datatype of %v (academy varaible) is %T \n", academy, academy)

	fever := 101.1
	check := true
	lifeLeft := 8

	fmt.Printf("Datatype of %v (fever varaible) is %T \n", fever, fever)
	fmt.Printf("Datatype of %v (check varaible) is %T \n", check, check)
	fmt.Printf("Datatype of %v (lifeLeft varaible) is %T \n", lifeLeft, lifeLeft)

}
