package main

import (
	"fmt"

	"github.com/Swarnavo2003/podcast/auth"
	"github.com/Swarnavo2003/podcast/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("kyojin","secret")

	session := auth.GetSession()
	fmt.Println("session :",session)

	user := user.User{
		Email: "kyojin@email.com",
		Name: "Kyojin",
	}

	// fmt.Println(user.Email, user.Name)
	color.Red(user.Email)
	color.Green(user.Name)
}