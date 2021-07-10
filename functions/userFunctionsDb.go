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
	//db_name := os.Getenv("DB_NAME")
	in := "INSERT INTO "+"user SET phone =?, name=?, otp=?, token=?"
	insert, err := db.Prepare(in)
	if err != nil {
		fmt.Println("error start from here ", err.Error())
		fmt.Println("ends")
		panic(err.Error())
	}
	_, err = insert.Exec(user.Phone, user.Name, user.Otp, user.Token)
	if err != nil {
		panic(err.Error())
	}
	return true
}

func InsertUserRate(rate db.Rating) bool {
	db := db.DBConn()
	defer db.Close()
	//db_name := os.Getenv("DB_NAME")
	in := "INSERT INTO "+"rating SET phone =?, rate=?"
	insert, err := db.Prepare(in)
	if err != nil {
		fmt.Println("error start from here ", err.Error())
		fmt.Println("ends")
		panic(err.Error())
	}
	_, err = insert.Exec(rate.Phone,rate.Rate)
	if err != nil {
		panic(err.Error())
	}
	return true
}

func GetUserRate(phone string)float64{
	db := db.DBConn()
	defer db.Close()
	query:="SELECT AVG(rate) FROM rating WHERE phone ="+phone
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		return 0
	} else {
		var avg_price float64
		for rows.Next() {
			rows.Scan(&avg_price)
		}
		return avg_price
	}
}

//done
func GetUserDb(phone string) (db.User, bool) {
	var users db.User
	dbb := db.DBConn()
	defer dbb.Close()
//	db_name := os.Getenv("DB_NAME")
	query := "SELECT id,phone,name, otp,token,imageUrl FROM "  + "user WHERE phone = ?"
	err := dbb.QueryRow(query, phone).Scan(&users.Id, &users.Phone, &users.Name, &users.Otp, &users.Token,&users.ImageUrl)
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
	query := "SELECT id,phone,name, otp,token,imageUrl FROM " + db_name + ".user "
	result, err := dbb.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	var id int
	var name, phone, otp, token,imageUrl string

	for result.Next() {
		err = result.Scan(&id, &phone, &name, &otp, &token,&imageUrl)
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
			ImageUrl: imageUrl,
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
	query := "UPDATE " + db_name + ".user set name =?, imageUrl= ? WHERE phone =?"
	update, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	_, err = update.Exec(user.Name,user.ImageUrl, phone)
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
