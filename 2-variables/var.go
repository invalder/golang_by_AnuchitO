package main

import "fmt"

/*
	The zero value of a variable is 0
	The zero value of a string is ""
	The zero value of a boolean is false
*/

/*
	basic types

	bool
	string

	int	int8	int16	int32	int64
	uint	uint8	uint16	uint32	uint64 uinptr
	byte // alias for uint8
	rune // alias for int32
	float32	float64

	complex64	complex128
*/

func main() {
	var name string = "Gopher นะจ๊ะ!!!"

	var name2 = "Gopher นะจ๊ะ!!!"

	//inside function body only
	name3 := ""

	fmt.Printf("name: %v\n", name)
	fmt.Printf("name2: %q\n", name2)
	fmt.Printf("name: %T\n", name)

	fmt.Printf("name2: %v\n", name2)
	fmt.Printf("name2: %q\n", name2)
	fmt.Printf("name2: %T\n", name2)

	fmt.Printf("name3: %v\n", name3)
	fmt.Printf("name3: %q\n", name3)
	fmt.Printf("name3: %T\n", name3)
}

// func funcnum() {
// 	fmt.Println(name)
// }
