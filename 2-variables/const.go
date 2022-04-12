package main

import "fmt"

func main() {
	var day = "Monday"
	fmt.Println("day:", day)

	day = "Muanday"
	fmt.Println("day:", day)

	// const (
	// 	Sunday    = 0
	// 	Monday    = 1
	// 	Tuesday   = 2
	// 	Wednesday = 3
	// 	Thursday  = 4
	// 	Friday    = 5
	// 	Saturday  = 6
	// )

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)

}
