package auth

import (
	"context"
	"net/http"
	"pawpawchat/internal/dto"
	"pawpawchat/utils/response"
	"time"
)

func (r *AuthRoutes) SignUp(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), time.Second*10)
	defer cancel()

	user, err := dto.ExtractUser(req.Body)
	if err != nil {
		response.BadReq(w, err.Error())
		return
	}

	return

	user, err = r.authClient.SignUp(ctx, user)
	if err != nil {
		response.BadReq(w, err.Error())
		return
	}

	response.Created(w, user)
}
