package main

import "fmt"

func main() {

	var rows int

	fmt.Print("Rows to Print the Half Star Pyramid = ")
	fmt.Scanln(&rows)
	fmt.Printf("Half Star Pyramid with %v rows \n", rows)

	for i := 1; i <= rows; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf(" * ")
		}
		fmt.Println()

	}

}
