package main

import (
	"demo/functions"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := functions.GetPort()
	router := mux.NewRouter()
	router.HandleFunc("/addUser", functions.AddUser).Methods("POST")
	router.HandleFunc("/getUser", functions.GetUser).Methods("GET")
	router.HandleFunc("/deleteUser", functions.DelteUser).Methods("DELETE")
	router.HandleFunc("/updateUser", functions.UpdateUSer).Methods("PUT")
	router.HandleFunc("/getAllUsers", functions.GetUsers).Methods("GET")
	http.ListenAndServe(":"+port, router)

}
