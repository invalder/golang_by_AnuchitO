package main

import "fmt"

func info(name, message string, age int) {
	fmt.Println("My name is:", name)
	fmt.Println("Message:", message)
	fmt.Printf("I'm %v yo:\n", age)

}

func today() string {
	return "มื้อนี้"
}

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	info("Gopher", "gopher เด้อ", 11)
	info("Vader", "ซำบายดีนะจ๊ะ", 50)

	day := today()
	fmt.Println(day)

	a, b := swap("ซำบายดีนะจ๊ะ", "Gopher")
	fmt.Println(a, b)
}
