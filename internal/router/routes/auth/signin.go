package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/internal/model/domain"
	web "pawpawchat/internal/model/web/auth"
	"pawpawchat/utils/response"
	"time"
)

func (r *AuthRoutes) SignIn(w http.ResponseWriter, req *http.Request) {
	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	signInRequest := &authpb.SignInRequest{}
	if err := json.NewDecoder(req.Body).Decode(signInRequest); err != nil {
		response.BadReq(w, err.Error())
		return
	}

	signInResponse, err := r.authClient.SignIn(context, signInRequest)
	if err != nil {
		// REFACTOR
		response.BadReq(w, err.Error())
		return
	}

	user := domain.User{}

	response.OK(w, &web.SignInResponse{User: user, Token: signInResponse.GetTokenString()})
}
