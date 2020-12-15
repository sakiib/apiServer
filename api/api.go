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

//@route GET /api/users
//@desc Gets all the available users

func getUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("getUsers")

	username, password, ok := request.BasicAuth()
	if authenticated := auth.CheckBasicAuthentication(credential.username, credential.password, username, password, ok, isAuthRequired); !authenticated {
		response.Write([]byte(`{"message":"Authentication Failed!"}`))
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(model.User{})
		return
	}
	fmt.Println("Authentication successful!")

	response.Write([]byte(`{"message":"Users List!"}`))
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(data.Users)
}

//@route GET /api/user/id
//@desc Gets a single user with the given id

func getUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("getUser")

	username, password, ok := request.BasicAuth()
	if authenticated := auth.CheckBasicAuthentication(credential.username, credential.password, username, password, ok, isAuthRequired); !authenticated {
		response.Write([]byte(`{"message":"Authentication Failed!"}`))
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(model.User{})
		return
	}
	fmt.Println("Authentication successful!")

	params := mux.Vars(request)
	for _, user := range data.Users {
		if user.ID == params["id"] {
			response.Write([]byte(`{"message":"User Found!"}`))
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode(user)
			return
		}
	}
	response.WriteHeader(http.StatusNoContent)
	json.NewEncoder(response).Encode(model.User{})
}

//@route POST /api/user/id
//@desc Create a new user with given info

func addUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("addUser")

	username, password, ok := request.BasicAuth()
	if authenticated := auth.CheckBasicAuthentication(credential.username, credential.password, username, password, ok, isAuthRequired); !authenticated {
		response.Write([]byte(`{"message":"Authentication Failed!"}`))
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(model.User{})
		return
	}
	fmt.Println("Authentication successful!")

	newUser := model.User{}
	if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
		log.Fatal(err)
	}

	for _, user := range data.Users {
		if newUser.ID == user.ID {
			response.WriteHeader(http.StatusConflict)
			response.Write([]byte(`{"message":"User with the given ID already exists!"}`))
			return
		}
	}

	response.Write([]byte(`{"message":"User Created!"}`))
	response.WriteHeader(http.StatusCreated)
	data.Users = append(data.Users, newUser)
	json.NewEncoder(response).Encode(data.Users)
}

//@route PUT /api/user/id
//@desc Update a user details with given id

func updateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("updateUser")

	username, password, ok := request.BasicAuth()
	if authenticated := auth.CheckBasicAuthentication(credential.username, credential.password, username, password, ok, isAuthRequired); !authenticated {
		response.Write([]byte(`{"message":"Authentication Failed!"}`))
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(model.User{})
		return
	}
	fmt.Println("Authentication successful!")

	newUser := model.User{}
	if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
		log.Fatal(err)
	}

	params := mux.Vars(request)
	userUpdated := false
	for index, user := range data.Users {
		if user.ID == params["id"] {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			data.Users = append(data.Users, newUser)
			userUpdated = true
			break
		}
	}
	if !userUpdated {
		response.WriteHeader(http.StatusNoContent)
		response.Write([]byte(`{"message":"User with the given ID not Found!"}`))
		return
	}

	response.Write([]byte(`{"message":"User Updated!"}`))
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(data.Users)
}

//@route DELETE /api/user/id
//@desc Delete an user with the given ID

func deleteUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("deleteUser")

	username, password, ok := request.BasicAuth()
	if authenticated := auth.CheckBasicAuthentication(credential.username, credential.password, username, password, ok, isAuthRequired); !authenticated {
		response.Write([]byte(`{"message":"Authentication Failed!"}`))
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(model.User{})
		return
	}
	fmt.Println("Authentication successful!")
	params := mux.Vars(request)
	userDeleted := false
	for index, user := range data.Users {
		if user.ID == params["id"] {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			userDeleted = true
			break
		}
	}

	if !userDeleted {
		response.WriteHeader(http.StatusNoContent)
		response.Write([]byte(`{"message":"User with the given ID not Found!"}`))
		return
	}

	response.Write([]byte(`{"message":"User Deleted!"}`))
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(data.Users)
}

func HandleRoutes(username, password, port string, authNeeded bool) {
	fmt.Println("in HandleRoutes!")
	credential.username = username
	credential.password = password
	isAuthRequired = authNeeded
	fmt.Println("cred. username: ", credential.username, "cred. password: ", credential.password)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/users", getUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/api/user/{id}", addUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
