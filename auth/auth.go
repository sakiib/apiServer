package auth

import "fmt"

func CheckBasicAuthentication(username, password, curUsername, curPassword string, ok, isAuthRequired bool) bool {
	if !isAuthRequired {
		fmt.Println("Auth check is not required, set by flag!")
		return true
	}

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
