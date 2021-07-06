package db

type Rating struct{
	Id 			int 	`json:"id"`
	Phone		string	`json:"phone"`
	Rate		float32	`json:"rate"`	
}