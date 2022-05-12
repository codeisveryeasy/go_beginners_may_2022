package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//initialize default input device with help of bufio and os standard package
	reader := bufio.NewReader(os.Stdin)
	//prompt the message on the screen depending on what you want the user to input
	fmt.Print("Enter your number: ")
	//use reader to read the string. Hit enter i.e. "\n" (new line) to specify the end of string user wants to enter
	text, _ := reader.ReadString('\n')
	//there is extra \n in the end of string. Need to trim extra \n and then convert the remaining string to int
	input := strings.TrimSuffix(text, "\n")
	fmt.Printf("input hello %v hello", input)
	fmt.Printf("Dataype of %v (input variable) is %T \n", input, input)
	//use strconv package to convert the int present in input (of datatype string) to int datatype
	number, _ := strconv.Atoi(input)
	myNumber := square(number)
	fmt.Printf("Square of %v is %v \n", number, myNumber)

}

func square(n int) int {
	return n * n
}
