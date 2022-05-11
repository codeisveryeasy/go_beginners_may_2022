package main

import "fmt"

func main() {

	number := 15
	myNumber := square(number)
	fmt.Printf("Square of %v is %v \n", number, myNumber)

}

func square(n int) int {
	return n * n
}
