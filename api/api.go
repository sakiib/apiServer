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
	fmt.Println("Authentication successful!")

	response.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(response).Encode(data.Users); err != nil {
		log.Fatal(err)
	}
}

//@route GET /api/user/id
//@desc Gets a single user with the given id

func GetUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("getUser")
	fmt.Println("Authentication successful!")

	ID := parseID(request)
	for _, user := range data.Users {
		if user.ID == ID {
			response.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(response).Encode(user); err != nil {
				log.Fatal(err)
			}
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
	if err := json.NewEncoder(response).Encode(data.Users); err != nil {
		log.Fatal(err)
	}
}

//@route PUT /api/user/id
//@desc Update a user details with given id

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("updateUser")
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
			if err := json.NewEncoder(response).Encode(data.Users); err != nil {
				log.Fatal(err)
			}
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
	fmt.Println("Authentication successful!")

	ID := parseID(request)
	for index, user := range data.Users {
		if user.ID == ID {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			response.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(response).Encode(data.Users); err != nil {
				log.Fatal(err)
			}
			return
		}
	}

	response.WriteHeader(http.StatusNoContent)
}

func HandleRoutes(port string) {
	fmt.Println("in HandleRoutes!")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/users", auth.BasicAuthentication(GetUsers)).Methods("GET")
	router.HandleFunc("/api/user/{id}", auth.BasicAuthentication(GetUser)).Methods("GET")
	router.HandleFunc("/api/user/{id}", auth.BasicAuthentication(AddUser)).Methods("POST")
	router.HandleFunc("/api/user/{id}", auth.BasicAuthentication(UpdateUser)).Methods("PUT")
	router.HandleFunc("/api/user/{id}", auth.BasicAuthentication(DeleteUser)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
