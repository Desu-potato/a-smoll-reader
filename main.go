package main

import (
	"fmt"
	"flag"
	"os"
	"errors"
)

var FILENAME string

func init() {
	flag.Parse()
	FILENAME = (flag.Args())[0]
}


func main() {
	_, err := os.Stat(FILENAME)
    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("file does not exist")
		os.Exit(0)
	}
	file, err := os.Open(FILENAME)
	if err != nil {
		fmt.Println("error occure:",err.Error())
		os.Exit(0)
	}
	readbuffer := make([]byte, 900)
	_, err = file.Read(readbuffer)
	if err != nil {
		fmt.Println("error occure:",err.Error())
		os.Exit(0)
	}
	fmt.Print(string(readbuffer))
}