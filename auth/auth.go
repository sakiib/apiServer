package auth

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	user string
	pass string
)

func BasicAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		auth := os.Getenv("auth")
		if auth == "false" {
			next.ServeHTTP(response, request)
			return
		}
		user = os.Getenv("username")
		pass = os.Getenv("password")

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

func GetToken() (string, error) {
	signingKey := []byte("mysecretsigninkeysakibalaminappscode")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user,
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("mysecretsigninkeysakibalaminappscode")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

func JWTAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		log.Println("in jwt auth func")
		auth := os.Getenv("auth")
		log.Println("auth: ", auth)

		if auth == "false" {
			next.ServeHTTP(response, request)
			return
		}

		tokenString := request.Header.Get("Authorization")
		if len(tokenString) == 0 {
			response.WriteHeader(http.StatusUnauthorized)
			response.Write([]byte("Missing Authorization Header"))
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := verifyToken(tokenString)

		if err != nil {
			response.WriteHeader(http.StatusUnauthorized)
			response.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}

		name := claims.(jwt.MapClaims)["name"].(string)
		request.Header.Set("name", name)

		next.ServeHTTP(response, request)
	}
}
