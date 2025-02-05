package model

type User struct {
	Firstname *string `json:"firstname" form:"firstname"`
	Lastname  *string `json:"lastname" form:"lastname"`
}
