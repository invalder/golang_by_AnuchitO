package main

import "fmt"

func main() {
	var langs = []string{"golang", "python", "java"}
	fmt.Printf("langs: %#v\n", langs)

	n := langs[0]
	fmt.Printf("%#v\n", n)

	langs[1] = "pythonistas"
	fmt.Printf("langs: %#v\n", langs)

	l := len(langs)
	fmt.Printf("len of langs: %v\n", l)

	cord := []string{"A", "B", "C"}
	langs = append(langs, "F#", "Em", "Am")
	langs = append(cord, langs...)
	fmt.Printf("langs: %#v\n", langs)
	fmt.Printf("len of langs: %v\n", len(langs))
}
