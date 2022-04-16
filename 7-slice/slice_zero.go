package main

import "fmt"

func main() {
	langs := []string{}
	fmt.Printf("langs: %#v\n", langs)

	if langs == nil {
		fmt.Println("langs is nil")
	} else {
		fmt.Printf("langs is not nil, value: %#v\n", langs)
	}
}
