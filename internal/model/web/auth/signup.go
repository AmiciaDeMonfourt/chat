package auth

import "pawpawchat/internal/model/domain"

type SignUpRequest struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type SignUpResponse struct {
	User  domain.User `json:"user"`
	Token string      `json:"token_string"`
}
