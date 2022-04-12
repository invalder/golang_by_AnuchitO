package user

import "fmt"

func Info(name, message string, age int) {
	fmt.Println("My name is:", name)
	fmt.Println("Message:", message)
	fmt.Printf("I'm %v year old:\n", age)
}
