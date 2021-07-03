package db


type User struct{
	Id			int 		`json:id`
	Phone 		string		`json:"phone"`
	Name		string 		`json:"name"`
	Otp			string		`json:"otp"`
	Token		string		`json:"token"`
}