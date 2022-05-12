package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0

	//all cli will be of datatype string
	//if i have to do mathematical operation
	//then i must convert each cli to int
	for i := 1; i < len(os.Args); i++ {
		tempNumber, _ := strconv.Atoi(os.Args[i])
		sum = sum + tempNumber
	}
	fmt.Println("Sum of CLI is: ", sum)
}
