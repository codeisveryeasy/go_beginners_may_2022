/*
Arrays allow developers to store multiple values of
same datatype in same variable


*/

package main

import "fmt"

func main() {
	var scores = [4]int{44, 71, 88, 31}
	fmt.Println(scores)
	//first value will always be at index position of 0 in an array
	fmt.Println(scores[0])
	fmt.Println(scores[1])
	fmt.Println(scores[2])
	fmt.Println(scores[3])

	scores[3] = 13
	fmt.Println(scores[3])
	fmt.Println(scores)

	fmt.Println("Length of array:", len(scores))
	fmt.Println("Last value in scores array: ", scores[len(scores)-1])
}
