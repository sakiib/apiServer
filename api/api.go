package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sakiib/apiServer/auth"
	"github.com/sakiib/apiServer/data"
	"github.com/sakiib/apiServer/model"
	"log"
	"net/http"
)

var credential struct {
	username string
	password string
}

var isAuthRequired bool

func parseID(request *http.Request) string {
	params := mux.Vars(request)
	ID := params["id"]
	if len(ID) > 0 {
		return ID
	}

	values := request.URL.Query()
	if val, ok := values["id"]; ok && len(val) > 0 {
		return val[0]
	}
	return ""
}

//@route GET /api/users
//@desc Gets all the available users

func GetUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("getUsers")

	username, password, ok := request.BasicAuth()
	if authenticated := auth.CheckBasicAuthentication(credential.username, credential.password, username, password, ok, isAuthRequired); !authenticated {
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(model.User{})
		return
	}
	fmt.Println("Authentication successful!")

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(data.Users)
}

//@route GET /api/user/id
//@desc Gets a single user with the given id

func GetUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("getUser")

	username, password, ok := request.BasicAuth()
	if authenticated := auth.CheckBasicAuthentication(credential.username, credential.password, username, password, ok, isAuthRequired); !authenticated {
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(model.User{})
		return
	}
	fmt.Println("Authentication successful!")

	ID := parseID(request)
	for _, user := range data.Users {
		if user.ID == ID {
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode(user)
			return
		}
	}

	response.WriteHeader(http.StatusNoContent)
}

//@route POST /api/user/id
//@desc Create a new user with given info

func AddUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("addUser")

	username, password, ok := request.BasicAuth()
	if authenticated := auth.CheckBasicAuthentication(credential.username, credential.password, username, password, ok, isAuthRequired); !authenticated {
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(model.User{})
		return
	}
	fmt.Println("Authentication successful!")

	newUser := model.User{}
	if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
		log.Fatal(err)
	}

	ID := parseID(request)
	for _, user := range data.Users {
		if ID == user.ID {
			response.WriteHeader(http.StatusConflict)
			return
		}
	}

	data.Users = append(data.Users, newUser)
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(data.Users)
}

//@route PUT /api/user/id
//@desc Update a user details with given id

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("updateUser")

	username, password, ok := request.BasicAuth()
	if authenticated := auth.CheckBasicAuthentication(credential.username, credential.password, username, password, ok, isAuthRequired); !authenticated {
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(model.User{})
		return
	}
	fmt.Println("Authentication successful!")

	newUser := model.User{}
	if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
		log.Fatal(err)
	}

	ID := parseID(request)
	for index, user := range data.Users {
		if user.ID == ID {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			data.Users = append(data.Users, newUser)
			response.WriteHeader(http.StatusCreated)
			json.NewEncoder(response).Encode(data.Users)
			return
		}
	}

	response.WriteHeader(http.StatusNoContent)
}

//@route DELETE /api/user/id
//@desc Delete an user with the given ID

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("deleteUser")

	username, password, ok := request.BasicAuth()
	if authenticated := auth.CheckBasicAuthentication(credential.username, credential.password, username, password, ok, isAuthRequired); !authenticated {
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(model.User{})
		return
	}
	fmt.Println("Authentication successful!")

	ID := parseID(request)
	for index, user := range data.Users {
		if user.ID == ID {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode(data.Users)
			return
		}
	}

	response.WriteHeader(http.StatusNoContent)
}

func HandleRoutes(username, password, port string, authNeeded bool) {
	fmt.Println("in HandleRoutes!")
	credential.username = username
	credential.password = password
	isAuthRequired = authNeeded
	fmt.Println("cred. username: ", credential.username, "cred. password: ", credential.password)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/users", GetUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/api/user/{id}", AddUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
