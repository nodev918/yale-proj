package main

import (
	"log"
	"io"
	"os"
)

func WriteFile(){
	_, err := io.WriteString(os.Stdout, "Hello World!")
	if err != nil{
		log.Fatal(err)
	}
}

func Append(){
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _,err = f.WriteString(text); err != nil {
		panic(err)
	}
}