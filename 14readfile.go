package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	filedata, error := ioutil.ReadFile("filesample1.txt")
	//error will be nil if file read was success. filedata will contain
	//the bytes of data

	//if error is not equal to nil i.e. it contains some message
	if error != nil {
		fmt.Println("Error: ")
		fmt.Println(error)
		return
	}
	fmt.Println("File read success....")
	fmt.Println(filedata)
	fmt.Println(string(filedata))

}
