package types

import "time"

// LoginDTO defined the /login payload
type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"password"`
}

// SignupDTO defined the /login payload
type SignupDTO struct {
	LoginDTO
	Name string `json:"name" validate:"required,min=3"`
}

// User table
type User struct {
	ID       	uint32 `json:"id"`
	Username	string `json:"username"`
	Email    	string `json:"email"`
	Password 	string `json:"-"`
}

// AccessResponse todo
type AccessResponse struct {
	Token 	string 		`json:"token"`
	Expire 	time.Time	`json:"expire"`
}

// AuthResponse todo
type AuthResponse struct {
	User *User			 `json:"user"`
	Auth *AccessResponse `json:"auth"`
}
