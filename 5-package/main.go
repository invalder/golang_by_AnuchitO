package main

//go mod init <project name>
import (
	"fmt"
	"igapp/time"
	"igapp/user"
)

func main() {
	user.Info("Nong", "gopher เด้อ", 15)

	t := time.Today()
	fmt.Println("Today is:", t)

}
