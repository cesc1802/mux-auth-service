package dto

type UserRegisterRequest struct {
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	LastName  *string `json:"last_name"`
	FirstName *string `json:"first_name"`
}
