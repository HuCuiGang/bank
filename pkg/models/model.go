package models

type User struct {
	Name     string  `json:"name"`
	Id       string  `json:"id"`
	Password string  `json:"password"`
	Money    float64 `json:"money"`
}
