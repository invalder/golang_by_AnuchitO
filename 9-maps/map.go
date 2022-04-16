package main

import "fmt"

func main() {
	status := map[int]string{
		200: "OK",
		308: "Permanent Redirect",
		400: "Bad Request",
		500: "Internal Server Error",
	}

	fmt.Printf("% #v\n", status)

	l := len(status)
	fmt.Printf("length: %d\n", l)

	status[200] = "OKie"
	status[285] = "Are you Drunk?"
	fmt.Printf("% #v\n", status)
	fmt.Printf("length: %d\n", len(status))

	value := status[200]
	fmt.Printf("value: %#v\n", value)

	delete(status, 285)
	fmt.Printf("% #v\n", status)
	fmt.Printf("length: %d\n", len(status))

	v, ok := status[666]
	if ok {
		fmt.Printf("value: %#v\n", v)
	} else {
		fmt.Println("key not found")
	}

	if v, ok := status[66]; ok {
		fmt.Printf("value: %#v\n", v)
	} else {
		fmt.Println("key not found")
	}

	// var m map[string]string = map[string]string{}
	m := make(map[string]string)
	fmt.Printf("%#v\n", m)

	if m == nil {
		fmt.Printf("m is nil. value : %#v\n", m)
	} else {
		fmt.Printf("m is not nil. value : %#v\n", m)
	}
}
