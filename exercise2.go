package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	//var teams = [4]string{"Manchester", "Liverpool", "Arsenal", "Chelsea"}
	//var run int
	//run = len(teams)
	current := time.Now()
	currentTime := strconv.Itoa(current.Hour()) + ":" + strconv.Itoa(current.Minute()) + ":" + strconv.Itoa(current.Second())

	for i := 1; i < len(os.Args); i++ {
		f, err := os.Create(os.Args[i])
		if err != nil {
			log.Fatal(err)
		}
		f.WriteString(currentTime)
		defer f.Close()
		fmt.Println(f.Name())
		fmt.Println(currentTime)
	}

}
