package authroutes

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/internal/model/domain"
	"pawpawchat/internal/model/web/auth"
	"pawpawchat/utils/response"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *AuthRoutes) SignUp(w http.ResponseWriter, req *http.Request) {
	signUpRequest := &auth.SignUpRequest{}

	// decode http request body
	if err := json.NewDecoder(req.Body).Decode(signUpRequest); err != nil {
		response.BadReq(w, err.Error())
		return
	}

	// verify sign up request credentials
	if err := checkCredentials(*signUpRequest); err != nil {
		response.BadReq(w, err.Error())
		return
	}

	// create new record in database
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	authreq := &authpb.SignUpRequest{Email: signUpRequest.Email, Password: signUpRequest.Password}

	authresp, err := r.authClient.SignUp(ctx, authreq)

	if err != nil {
		st := status.Convert(err)

		switch st.Code() {
		case codes.InvalidArgument:
			response.BadReq(w, st.Err().Error())
			return
		default:
			response.InternalErr(w, st.Err().Error())
			return
		}
	}

	signUpResponse := &auth.SignUpResponse{
		User: domain.User{
			ID:         authresp.User.GetId(),
			FirstName:  authresp.User.GetFirstName(),
			SecondName: authresp.User.GetSecondName()},
		Token: authresp.GetTokenString(),
	}

	response.Created(w, signUpResponse)
}

func checkCredentials(req auth.SignUpRequest) error {
	if req.Email == "" {
		return errors.New("email is missing")
	}

	if req.FirstName == "" {
		return errors.New("first name is missing")
	}

	if req.SecondName == "" {
		return errors.New("second name is missing")
	}

	if req.Password == "" {
		return errors.New("password is missing")
	}

	return nil
}
