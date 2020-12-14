package api

import (
	"apiServer/auth"
	"apiServer/data"
	"apiServer/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var credential struct {
	username string
	password string
}

func getUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("getUsers")

	username, password, ok := request.BasicAuth()
	if authenticated := auth.CheckBasicAuthentication(credential.username, credential.password, username, password, ok); !authenticated {
		response.Write([]byte(`{"message":"Authentication Failed!"}`))
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(model.User{})
		return
	}
	fmt.Println("Authentication successful!")

	json.NewEncoder(response).Encode(data.Users)
}

func getUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Println("getUser")

	username, password, ok := request.BasicAuth()
	if authenticated := auth.CheckBasicAuthentication(credential.username, credential.password, username, password, ok); !authenticated {
		response.Write([]byte(`{"message":"Authentication Failed!"}`))
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(model.User{})
		return
	}
	fmt.Println("Authentication successful!")

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

	username, password, ok := request.BasicAuth()
	if authenticated := auth.CheckBasicAuthentication(credential.username, credential.password, username, password, ok); !authenticated {
		response.Write([]byte(`{"message":"Authentication Failed!"}`))
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(model.User{})
		return
	}
	fmt.Println("Authentication successful!")

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

	username, password, ok := request.BasicAuth()
	if authenticated := auth.CheckBasicAuthentication(credential.username, credential.password, username, password, ok); !authenticated {
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

	username, password, ok := request.BasicAuth()
	if authenticated := auth.CheckBasicAuthentication(credential.username, credential.password, username, password, ok); !authenticated {
		response.Write([]byte(`{"message":"Authentication Failed!"}`))
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(model.User{})
		return
	}
	fmt.Println("Authentication successful!")
	params := mux.Vars(request)
	for index, user := range data.Users {
		if user.ID == params["id"] {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			break
		}
	}
	json.NewEncoder(response).Encode(data.Users)
}

func HandleRoutes(username, password string) {
	fmt.Println("in HandleRoutes!")
	credential.username = username
	credential.password = password
	fmt.Println("cred. username: ", credential.username, "cred. password: ", credential.password)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/users", getUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/api/user/{id}", addUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}