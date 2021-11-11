package models

//type User struct {
//	Name     string  `json:"name"`
//	Id       string  `json:"id"`
//	Password string  `json:"password"`
//	Money    float64 `json:"money"`
//}

type User struct {
	UserName     string  `db:"name" json:"name"`
	Id       string  `db:"id" json:"id"`
	Password string  `db:"password" json:"password"`
	Balance    float64 `db:"balance" json:"balance"`
}
