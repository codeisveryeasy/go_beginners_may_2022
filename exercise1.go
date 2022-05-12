package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var teams = [4]string{"Manchester", "Liverpool", "Arsenal", "Chelsea"}
	var run int
	run = len(teams)

	for i := 0; i < run; i++ {
		f, err := os.Create(teams[i] + ".txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		fmt.Println(f.Name())
	}

}
