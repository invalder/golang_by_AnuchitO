package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Monday

	switch today {
	case time.Monday:
		fmt.Println("Today is Monday")
		// fallthrough do below case
	case time.Saturday, time.Sunday:
		fmt.Println("Today is Weekend")
	default:
		fmt.Println("Today is not Monday")
	}
}
