package functions

import (
	"demo/db"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//done
func AddUser(w http.ResponseWriter, r *http.Request) {
	var creds db.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		token := CreateToken(w, r, creds.Phone)
		user := db.User{
			Phone: creds.Phone,
			Name:  creds.Name,
			Otp:   creds.Otp,
			Token: token,
		}
		InsertUserData(user)
		w.WriteHeader(http.StatusAccepted)
		return
	}
}

//done
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//userPhone := user.Phone
	user, founded := GetUserDb(params["phone"])
	if founded {
		json.NewEncoder(w).Encode(&user)
		w.WriteHeader(http.StatusAccepted)
		return
	} else {
		json.NewEncoder(w).Encode(&user)
		w.WriteHeader(http.StatusAccepted)
		return
	}
}

//done
func GetUsers(w http.ResponseWriter, r *http.Request) {
	//var user db.User
	user, founded := GetAllUsersDb()
	if founded {
		json.NewEncoder(w).Encode(&user)
		w.WriteHeader(http.StatusAccepted)
		return
	} else {
		json.NewEncoder(w).Encode(&user)
		w.WriteHeader(http.StatusAccepted)
		return
	}
}

//done
func UpdateUSer(w http.ResponseWriter, r *http.Request) {
	var userCred db.User

	err := json.NewDecoder(r.Body).Decode(&userCred)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		user := db.User{
			Name: userCred.Name,
		}
		updated := UpdateUSerDb(userCred.Phone, user)
		if updated {
			w.WriteHeader(http.StatusAccepted)
		}
	}
}

//done
func DelteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	deleted := DeleteUserDb(params["phone"])
	if deleted {
		w.WriteHeader(http.StatusAccepted)
	}

}

func DisplayWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}
