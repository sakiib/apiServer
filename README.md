# API server

### RESTful API using [go](https://github.com/golang), [cobra CLI](https://github.com/spf13/cobra) & [gorilla mux](https://github.com/gorilla/mux)

--- 
API Endpoints
| Endpoint | Function | Method | StatusCode |
| -------- | -------- | ------ | ---------- |
| `/api/users` | GetUsers - Get all the users | GET | Success - StatusOK |
| `/api/user/{id}` | GetUser - Get user with the `id` | GET | Success - StatusOK, Failure - StatusNoContent |
| `/api/user/{id}` | AddUser - Add user with the `id` | POST | Success - StatusCreated, Failure - StatusConflict |
| `/api/user/{id}` | UpdateUser - Update user with the `id` | PUT | Success - StatusCreated, Failure - StatusNoContent |
| `/api/user/{id}` | DeleteUser - Delete user with the `id` | DELETE | Success - StatusOK, Failure - StatusNoContent |

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
User Data Model
```
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
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


