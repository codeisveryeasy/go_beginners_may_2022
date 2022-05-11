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
	fmt.Print("Enter your number for table: ")
	//use reader to read the string. Hit enter i.e. "\n" (new line) to specify the end of string user wants to enter
	text, _ := reader.ReadString('\n')
	//there is extra \n in the end of string. Need to trim extra \n and then convert the remaining string to int
	input := strings.TrimSuffix(text, "\n")
	fmt.Printf("Dataype of %v (input variable) is %T \n", input, input)
	//use strconv package to convert the int present in input (of datatype string) to int datatype
	number, _ := strconv.Atoi(input)

	//number := 15
	printMultiplicationTable(number)

}

func printMultiplicationTable(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v x %v = %v \n", n, i, n*i)
	}
}
