package main

import "fmt"

type User struct {
	Username      string
	FullName      string
	ProfilePicURL string
}

//User Method Info
func (u User) Info() {
	fmt.Printf("User Name: %v\n", u.Username)
	fmt.Printf("Full Name: %v\n", u.FullName)
	fmt.Printf("Profile Pic URL: %v\n", u.ProfilePicURL)
}

//function Info
func info(u User) {
	fmt.Printf("User Name: %v\n", u.Username)
	fmt.Printf("Full Name: %v\n", u.FullName)
	fmt.Printf("Profile Pic URL: %v\n", u.ProfilePicURL)
}

func main() {
	u := User{
		Username:      "golang",
		FullName:      "Basic Golang",
		ProfilePicURL: "https://knowhere.fake/gopher.jpg",
	}

	info(u)
	u.Info()
}
