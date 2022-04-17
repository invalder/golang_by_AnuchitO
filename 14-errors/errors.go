package main

import (
	"errors"
	"fmt"
	"log"
)

func ReadFile(filename string) (string, error) {
	// read file
	// var err error = errors.New("file not found")
	// err := errors.New("file not found")

	if filename != "" {
		return "data...", nil
	} else {
		return "", errors.New("file not found")
	}
}

func main() {
	data, error := ReadFile("")
	if error != nil {
		// handle error
		log.Println(error)
		return
	}

	fmt.Println("read file successfully : ", data)
}
