package model

type RegistrationLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type ResetPasswordRequest struct {
	OldPassword string `json:"oldpassword" form:"oldpassword"`
	NewPassword string `json:"newpassword" form:"newpassword"`
}

type UserRequest struct {
	Firstname *string `json:"firstname" form:"firstname"`
	Lastname  *string `json:"lastname" form:"lastname"`
	Bio       *string `json:"bio" form:"bio"`
}
