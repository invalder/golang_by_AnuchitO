package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//while loop
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// infiniteLoop()
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }

	langs := []string{"Go", "Python", "Ruby", "JavaScript"}

	// for loop basic
	for i := 0; i < len(langs); i++ {
		value := langs[i]
		fmt.Println(i, ":", value)
	}

	// for range slice
	for index, value := range langs {
		fmt.Println(index, ":", value)
	}

	// only value
	for _, value := range langs {
		fmt.Println(value)
	}

	// only index
	for index := range langs {
		fmt.Println(index)
	}
}
