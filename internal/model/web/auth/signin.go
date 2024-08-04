package auth

import "pawpawchat/internal/model/domain/user"

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	User  user.User `json:"user"`
	Token string    `json:"token_string"`
}
