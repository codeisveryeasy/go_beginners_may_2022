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

	//decalare the array of size 8 but no values
	var emptyArray = [8]int{}
	fmt.Println(emptyArray)

	//declare the array with only specific values at specific positions
	var specificArray = [12]int{0: 8, 4: 7, 11: 1}
	fmt.Println(specificArray)

	//array of strings
	var friends = [4]string{"obb", "oki", "mai", "che"}
	fmt.Println(friends)
	fmt.Println(friends[0])
	fmt.Println(friends[1])
	fmt.Println(friends[2])
	fmt.Println(friends[3])

	friends[3] = "wow"
	fmt.Println(friends)

}
