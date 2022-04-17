package main

import "fmt"

func main() {
	me := "Gopher"
	fmt.Printf("You are %s\n", me)

	var addr *string = &me
	fmt.Printf("% #v\n", addr)
	fmt.Printf("% #v\n", *addr)
	*addr = "Penguin"
	fmt.Printf("% #v\n", &me)
	fmt.Printf("% #v\n", me)
}
