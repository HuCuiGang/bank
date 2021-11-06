package models

//type User struct {
//	Name     string  `json:"name"`
//	Id       string  `json:"id"`
//	Password string  `json:"password"`
//	Money    float64 `json:"money"`
//}

type User struct {
	Name     string  `db:"name"`
	Id       string  `db:"id"`
	Password string  `db:"password"`
	Balance    float64 `db:"balance"`
}