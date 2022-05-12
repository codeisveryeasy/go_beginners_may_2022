package main

import (
	"fmt"
	"time"
)

func main() {

	current := time.Now()
	fmt.Println(current)
	fmt.Print(current.Hour(), ":", current.Minute(), ":", current.Second(), " on ", current.Weekday(), "\n")
	fmt.Println(current.Location())

}
