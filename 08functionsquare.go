package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	//initialize default input device with help of bufio and os standard package
	reader := bufio.NewReader(os.Stdin)
	//prompt the message on the screen depending on what you want the user to input
	fmt.Print("Enter your number: ")
	//use reader to read the string. Hit enter i.e. "\n" (new line) to specify the end of string user wants to enter
	text, _ := reader.ReadString('\n')

	//use strconv package to convert the int present in text (of type string) to int type
	number, _ := strconv.ParseInt(text, 0, 64)
	myNumber := square(number)
	fmt.Printf("Square of %v is %v \n", number, myNumber)

}

func square(n int64) int64 {
	return n * n
}
