package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	fmt.Print("main\n")

	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	Write()
	Append()
}