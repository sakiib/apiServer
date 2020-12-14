package auth

import "fmt"

func CheckBasicAuthentication(username, password, curUsername, curPassword string, ok bool) bool {
	if !ok {
		fmt.Println("Authentication credentials, Username or Password not provided!")
		return false
	}

	if username != curUsername || password != curPassword {
		fmt.Println("Wrong Username or Password!")
		return false
	}

	return true
}