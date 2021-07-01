package functions

import (
	"demo/db"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//done
func InsertUserData(user db.User) bool {
	db := db.DBConn()
	defer db.Close()
	db_name := os.Getenv("DB_NAME")
	in := "INSERT INTO " + db_name + ".user SET phone =?, name=?, otp=?, token=?"
	insert, err := db.Prepare(in)
	if err != nil {
		panic(err.Error())
	}
	_, err = insert.Exec(user.Phone, user.Name, user.Otp, user.Token)
	if err != nil {
		panic(err.Error())
	}
	return true
}

//done
func GetUserDb(phone string) (db.User, bool) {
	var users db.User
	dbb := db.DBConn()
	defer dbb.Close()
	db_name := os.Getenv("DB_NAME")
	query := "SELECT id,phone,name, otp,token FROM " + db_name + ".user WHERE phone = ?"
	err := dbb.QueryRow(query, phone).Scan(&users.Id, &users.Phone, &users.Name, &users.Otp, &users.Token)
	if err != nil {
		fmt.Println(err)
		return users, false
	}
	return users, true
}

//done
func GetAllUsersDb() ([]db.User, bool) {
	users := []db.User{}
	dbb := db.DBConn()
	defer dbb.Close()
	db_name := os.Getenv("DB_NAME")
	query := "SELECT id,phone,name, otp,token FROM " + db_name + ".user "
	result, err := dbb.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	var id int
	var name, phone, otp, token string

	for result.Next() {
		err = result.Scan(&id, &phone, &name, &otp, &token)
		if err != nil {
			fmt.Println("error ", err.Error())
			return nil, false
		}
		user := db.User{
			Id:    id,
			Name:  name,
			Phone: phone,
			Otp:   otp,
			Token: token,
		}
		users = append(users, user)
	}

	return users, true
}

//done
func UpdateUSerDb(phone string, user db.User) bool {

	db := db.DBConn()
	defer db.Close()
	db_name := os.Getenv("DB_NAME")
	query := "UPDATE " + db_name + ".user set name =? WHERE phone =?"
	update, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	_, err = update.Exec(user.Name, phone)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

//done
func DeleteUserDb(phone string) bool {
	db := db.DBConn()
	defer db.Close()
	db_name := os.Getenv("DB_NAME")
	query := "DELETE FROM " + db_name + ".user WHERE phone=?"
	delete, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	_, err = delete.Exec(phone)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true

}

//done
func CreateToken(w http.ResponseWriter, r *http.Request, phone string) string {

	var jwtKey = []byte("mandoby_project")
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := db.Claims{
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return ""
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	return tokenString

}

func GetPort() string{
	dbb := db.DBConn()
	defer dbb.Close()
	return os.Getenv("PORT")
}
