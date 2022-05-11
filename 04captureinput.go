package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//initialize default input device with help of bufio and os standard package
	reader := bufio.NewReader(os.Stdin)
	//prompt the message on the screen depending on what you want the user to input
	fmt.Print("Enter you name: ")
	//use reader to read the string. Hit enter i.e. "\n" (new line) to specify the end of string user wants to enter
	text, _ := reader.ReadString('\n')

	//print the string entered by user in the console/terminal
	fmt.Printf("Your name is: %v \n", text)
	fmt.Printf("Datatype of %v is %T \n", text, text)

}
