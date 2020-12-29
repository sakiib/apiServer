# API server

### RESTful API using [go](https://github.com/golang), [cobra CLI](https://github.com/spf13/cobra), [gorilla mux](https://github.com/gorilla/mux), Basic Auth, [JWT Auth](https://github.com/dgrijalva/jwt-go) [![Go Report Card](https://goreportcard.com/badge/github.com/sakiib/apiServer)](https://goreportcard.com/report/github.com/sakiib/apiServer)

--- 
API Endpoints
| Endpoint | Function | Method | StatusCode | Auth |
| -------- | -------- | ------ | ---------- | ---- |
| `/api/login` | LogIn | POST | Success - StatusOK, Failure - StatusUnauthorized | Basic |
| `/api/users` | GetUsers | GET | Success - StatusOK | JWT |
| `/api/user/{id}` | GetUser | GET | Success - StatusOK, Failure - StatusNoContent | JWT |
| `/api/user/{id}` | AddUser | POST | Success - StatusCreated, Failure - StatusConflict | JWT |
| `/api/user/{id}` | UpdateUser | PUT | Success - StatusCreated, Failure - StatusNoContent | JWT |
| `/api/user/{id}` | DeleteUser | DELETE | Success - StatusOK, Failure - StatusNoContent | JWT |

---
Installation
* `go install github.com/sakiib/apiServer`

---
CLI Commands:
* help with the start commands `apiServer start -h` or `apiServer start --help`
* start the API server on the given port (def: 8080) `apiServer start --port=8080`
* start the API server with no auth required flag (def: auth required): `apiServer start --auth=false`

--- 
Set Environment variables for Basic Authentication
`export username=sakib`
`export password=12345`

---
Data Model
```
package model

type Movie struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Genre  string `json:"genre"`
	Rating int    `json:"rating"`
}

```
```
package model

type User struct {
	ID              string  `json:"id"`
	FirstName       string  `json:"firstname"`
	LastName        string  `json:"lastname"`
	FavouriteMovies []Movie `json:"favouriteMovies"`
}

```

---
Authentication Method
* Basic Authentication
* JWT Authentication (ToDo)

---
Testing the API Endpoints
* Primary API endpoints testing using [Postman](https://github.com/postmanlabs) 
* E2E Testing. 
	* modlues used: `net/http/httptest`, `testing`, `bytes`, `encoding/json`, `net/http`. 
	* Checks for the response Status Code against the Expected Status Code.

---
Resources:
* [sysdevbd learn GO](https://sysdevbd.com/go/)
* [Encoding & Decoding JSON](https://kevin.burke.dev/kevin/golang-json-http/)
* [A Beginner’s Guide to HTTP and REST](https://code.tutsplus.com/tutorials/a-beginners-guide-to-http-and-rest--net-16340)
* [HTTP Response Status Codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)
* [TutorialEdge Creating A Simple Web Server with Golang](https://tutorialedge.net/golang/creating-simple-web-server-with-golang/)


