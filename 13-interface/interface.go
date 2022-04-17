package main

import "fmt"

type Phone interface {
	Call(number string)
	Sms(number string, message string)
}

type Samsung struct {
	Name string
}

func (s Samsung) Call(number string) {
	fmt.Printf("%s is calling %s\n", s.Name, number)
}

func (s Samsung) Sms(number string, message string) {
	fmt.Printf("%s is sending sms to %s: %s\n", s.Name, number, message)
}

func (s Samsung) Answer(number string) {
	fmt.Printf("%s is answering %s\n", s.Name, number)
}

func Dial(p Phone, number string) {
	p.Call(number)
}

func main() {
	s := Samsung{
		Name: "S10",
	}

	Dial(s, "+66666666666")
}
