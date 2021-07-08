package main

import (
	"demo/functions"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	os.Setenv("DB_HOST", "us-cdbr-east-04.cleardb.com")
	os.Setenv("DB_PORT", "(us-cdbr-east-04.cleardb.com)")
	os.Setenv("DB_DRIVER", "mysql")
	os.Setenv("DB_ROOT", "ba9602972fccac")
	os.Setenv("DB_PASSWORD", "15fb18e3")
	os.Setenv("DB_NAME", "heroku_3478a5eb4b093f3")
	//os.Setenv("PORT", "8083")
	port := os.Getenv("PORT")

	router := mux.NewRouter()
	router.HandleFunc("/", functions.DisplayWelcome)
	router.HandleFunc("/addUser", functions.AddUser).Methods("POST")
	router.HandleFunc("/addRate", functions.AddRate).Methods("POST")
	router.HandleFunc("/getUser/{phone}", functions.GetUser).Methods("GET")
	router.HandleFunc("/isAuthorized/{toCall}", functions.IsAuthorized).Methods("GET")
	router.HandleFunc("/getUserRate/{phone}", functions.GetRate).Methods("GET")
	router.HandleFunc("/deleteUser/{phone}", functions.DelteUser).Methods("DELETE")
	router.HandleFunc("/updateUser", functions.UpdateUSer).Methods("PUT")
	router.HandleFunc("/getAllUsers", functions.GetUsers).Methods("GET")
	http.ListenAndServe(":"+port, router)

}
