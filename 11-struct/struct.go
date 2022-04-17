package main

import "fmt"

type User struct {
	Username      string
	FullName      string
	ProfilePicURL string
}

func main() {
	usernames := "golang"
	fullname := "Basic Golang"
	profilePicURL := "https://knowhere.fake/gopher.jpg"
	fmt.Println(usernames, fullname, profilePicURL)

	// u := User{}
	// u.Username = usernames
	// u.FullName = fullname
	// u.ProfilePicURL = profilePicURL

	u := User{
		Username:      usernames,
		FullName:      fullname,
		ProfilePicURL: profilePicURL,
	}

	fmt.Printf("%#v\n", u)

	fmt.Println(u.Username)
	fmt.Println(u.FullName)
	fmt.Println(u.ProfilePicURL)
}
