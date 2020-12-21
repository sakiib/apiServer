package api

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sakiib/apiServer/auth"
	"github.com/sakiib/apiServer/data"
	"github.com/sakiib/apiServer/model"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	log.Println("getUsers")
	log.Println("Authentication successful!")

	response.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(response).Encode(data.Users); err != nil {
		log.Println(err.Error())
		return
	}
}

//@route GET /api/user/id
//@desc Gets a single user with the given id

func GetUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	log.Println("getUser")
	log.Println("Authentication successful!")

	ID := parseID(request)
	for _, user := range data.Users {
		if user.ID == ID {
			response.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(response).Encode(user); err != nil {
				log.Println(err.Error())
				return
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
	log.Println("addUser")
	log.Println("Authentication successful!")

	newUser := model.User{}
	if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
		log.Println(err.Error())
		return
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
		log.Println(err.Error())
		return
	}
}

//@route PUT /api/user/id
//@desc Update a user details with given id

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	log.Println("updateUser")
	log.Println("Authentication successful!")

	newUser := model.User{}
	if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
		log.Println(err.Error())
		return
	}

	ID := parseID(request)
	for index, user := range data.Users {
		if user.ID == ID {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			data.Users = append(data.Users, newUser)
			response.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(response).Encode(data.Users); err != nil {
				log.Println(err.Error())
				return
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
	log.Println("deleteUser")
	log.Println("Authentication successful!")

	ID := parseID(request)
	for index, user := range data.Users {
		if user.ID == ID {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			response.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(response).Encode(data.Users); err != nil {
				log.Println(err.Error())
				return
			}
			return
		}
	}

	response.WriteHeader(http.StatusNoContent)
}

func LogIn(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	log.Println("LogIn")
	log.Println("Authentication successful!")
	log.Println("successfully logged in!")

	token, err := auth.GetToken()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_, _ = response.Write([]byte("Error generating JWT token: " + err.Error()))
	} else {
		response.Header().Set("Authorization", "Bearer "+token)
		response.WriteHeader(http.StatusOK)
		_, _ = response.Write([]byte("Token: " + token))
	}
}

func HandleRoutes(port string) {
	log.Println("in HandleRoutes!")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/login", auth.BasicAuthentication(LogIn)).Methods("POST")
	router.HandleFunc("/api/users", auth.JWTAuthentication(GetUsers)).Methods("GET")
	router.HandleFunc("/api/user/{id}", auth.JWTAuthentication(GetUser)).Methods("GET")
	router.HandleFunc("/api/user/{id}", auth.JWTAuthentication(AddUser)).Methods("POST")
	router.HandleFunc("/api/user/{id}", auth.JWTAuthentication(UpdateUser)).Methods("PUT")
	router.HandleFunc("/api/user/{id}", auth.JWTAuthentication(DeleteUser)).Methods("DELETE")

	//log.Fatal(http.ListenAndServe(":"+port, router))
	// gracefully shutdown the server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
