package main

import "fmt"

func main() {
	var langs = []string{"golang", "python", "java"}
	fmt.Printf("langs: %#v\n", langs)

	a := langs[0:2]
	fmt.Printf("a: %#v\n", a)
	fmt.Printf("langs: %#v\n", langs)

	b := langs[1:3]
	fmt.Printf("b: %#v\n", b)
	fmt.Printf("langs: %#v\n", langs)

	n := langs[0:]
	fmt.Printf("n: %#v\n", n)
	fmt.Printf("langs: %#v\n", langs)

	r := langs[0:3]
	s := langs[:3]
	t := langs[0:]
	u := langs[:]

	fmt.Printf("r: %#v\n", r)
	fmt.Printf("s: %#v\n", s)
	fmt.Printf("t: %#v\n", t)
	fmt.Printf("u: %#v\n", u)

	printSlices(langs)
	cord := [4]string{"A", "B", "C", "D"}
	printSlices(cord[:])

}

func printSlices(s []string) {
	fmt.Printf("Print Slice: s: %#v\n", s)
}
