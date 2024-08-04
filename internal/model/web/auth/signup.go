package auth

import "pawpawchat/internal/model/domain/user"

type SignUpRequest struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type SignUpResponse struct {
	User  user.User `json:"user"`
	Token string    `json:"token_string"`
}
