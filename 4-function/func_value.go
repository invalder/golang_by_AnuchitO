package main

import "fmt"

var name string = "Nong"

// var plus = func(a, b int) int { return a + b }

// func plus(a, b int) int {
// 	return a + b
// }

func say(n string) {
	fmt.Printf("My name is %s\n", n)
}

func cal(op func(int, int) int) {
	r := op(4, 5)
	fmt.Println("result = ", r)
}

func main() {
	fmt.Printf("value: %v\n", name)
	fmt.Printf("name: %T\n", name)
	say(name)

	plus := func(a, b int) int { return a + b }

	p := plus(1, 2)
	fmt.Println("1+2 = ", p)
	fmt.Printf("plus: %T\n", plus)

	cal(plus)

	minus := func(a, b int) int { return a - b }

	cal(minus)
}
