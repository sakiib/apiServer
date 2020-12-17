package auth

import (
	"net/http"
	"os"
)

func BasicAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		user := os.Getenv("username")
		pass := os.Getenv("password")

		username, password, authOK := request.BasicAuth()
		if authOK == false {
			http.Error(response, "Not authorized", http.StatusUnauthorized)
			return
		}

		if username != user || password != pass {
			http.Error(response, "Not authorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(response, request)
	}
}
