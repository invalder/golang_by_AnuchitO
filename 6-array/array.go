package main

import "fmt"

func showAll(ns [4]string) {
	fmt.Printf("%#v\n", ns)
}

func main() {
	var langs = [4]string{"golang", "python", "java"}

	fmt.Printf("langs: %v\n", langs)
	fmt.Printf("langs: %#v\n", langs)

	n := langs[1]
	fmt.Printf("%#v\n", n)

	langs[1] = "pythonistas"
	fmt.Printf("langs: %#v\n", langs)

	showAll(langs)

	cord := [4]string{"A", "B", "C"}

	showAll(cord)
}
