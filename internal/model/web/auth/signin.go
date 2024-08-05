package auth

import "pawpawchat/internal/model/domain"

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	User  domain.User `json:"user"`
	Token string      `json:"token_string"`
}
