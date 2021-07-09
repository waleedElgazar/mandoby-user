package functions

import (
	"demo/db"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
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

func AddRate(w http.ResponseWriter, r *http.Request) {
	var rate db.Rating
	err := json.NewDecoder(r.Body).Decode(&rate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		rate := db.Rating{
			Rate:  rate.Rate,
			Phone: rate.Phone,
		}
		InsertUserRate(rate)
		w.WriteHeader(http.StatusAccepted)
		return
	}
}

func GetRate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	rate := GetUserRate(params["phone"])
	json.NewEncoder(w).Encode(rate)
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

func GetUserToCall(w http.ResponseWriter, r *http.Request,phone string ) db.User{
	w.Header().Set("Content-Type", "application/json")
	user, _ := GetUserDb(phone)
	return user
}
var mySigningKey = []byte("mandopy-project")

func IsAuthorized(w http.ResponseWriter, r *http.Request)  {
	var user db.User
	if r.Header["Token"] != nil {

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return mySigningKey, nil
		})
		if err != nil {

			fmt.Fprintf(w, err.Error())
		}
		if token.Valid {
			param := mux.Vars(r)
			user =GetUserToCall(w,r,param["toCall"])
		}
	} else {
		param := mux.Vars(r)
		user=GetUserToCall(w,r,param["toCall"])
	}
	json.NewEncoder(w).Encode(user)
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
	fmt.Fprint(w, "approved")
}
