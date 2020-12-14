package api

import (
	"apiServer/data"
	"apiServer/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
)

func getUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("getUsers")

	json.NewEncoder(response).Encode(data.Users)
}

func getUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("getUser")

	params := mux.Vars(request)
	for _, user := range data.Users {
		if user.ID == params["id"] {
			json.NewEncoder(response).Encode(user)
			return
		}
	}
	response.WriteHeader(http.StatusNoContent)
	json.NewEncoder(response).Encode(model.User{})
}

func addUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("addUser")

	user := model.User{}
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		log.Fatal(err)
	}

	data.Users = append(data.Users, user)
	json.NewEncoder(response).Encode(data.Users)
}

func updateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("updateUser")

	newUser := model.User{}
	if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
		log.Fatal(err)
	}

	params := mux.Vars(request)
	for index, user := range data.Users {
		if user.ID == params["id"] {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			data.Users = append(data.Users, newUser)
			break
		}
	}
	json.NewEncoder(response).Encode(data.Users)
}

func deleteUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("deleteUser")

	params := mux.Vars(request)
	for index, user := range data.Users {
		if user.ID == params["id"] {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			break
		}
	}
	json.NewEncoder(response).Encode(data.Users)
}

func HandleRoutes() {
	fmt.Println("in HandleRoutes!")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/users", getUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/api/user/{id}", addUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}