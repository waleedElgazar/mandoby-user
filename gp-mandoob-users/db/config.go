package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DBConn() (db *sql.DB) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	db_driver := os.Getenv("DB_DRIVER")
	db_User := os.Getenv("DB_ROOT")
	db_Password := os.Getenv("DB_PASSWORD")
	db_Port := os.Getenv("DB_PORT")
	db_Name := os.Getenv("DB_NAME")

	//path:=fmt.Sprintf(db_User,db_Password,"@tcp",db_Port,"/",db_Name)
	path := db_User + ":" + db_Password + "@tcp" + db_Port + "/" + db_Name
	fmt.Println(path)
	db, err = sql.Open(db_driver, path)
	if err != nil {
		panic(err.Error())
	}
	return db
}
