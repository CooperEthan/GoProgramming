package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(s)
	fmt.Println(bs)
	loginpassword := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginpassword))
	if err != nil {
		fmt.Println("You cant login")
		return
	}
	fmt.Println("you are logged in!!!")
}
