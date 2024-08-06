package authroutes

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/internal/model/domain"
	web "pawpawchat/internal/model/web/auth"
	"pawpawchat/utils/response"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *AuthRoutes) SignUp(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	signUpRequest := &web.SignUpRequest{}
	if err := json.NewDecoder(req.Body).Decode(signUpRequest); err != nil {
		response.BadReq(w, err.Error())
		return
	}

	if err := checkCredentials(*signUpRequest); err != nil { // verify sign up request credentials
		response.BadReq(w, err.Error())
		return
	}

	authreq := &authpb.SignUpRequest{
		Credentials: &authpb.Credentials{
			Email:    signUpRequest.Email,
			Password: signUpRequest.Password,
		},
		Userinfo: &authpb.PersonalInfo{
			Firstname:  signUpRequest.FirstName,
			Secondname: signUpRequest.SecondName,
		},
	}
	authresp, err := r.authClient.SignUp(ctx, authreq)

	// REFACTOR
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

	signUpResponse := &web.SignUpResponse{
		User: domain.User{
			UserID: authresp.GetUser().GetId(),
			PersonalInfo: domain.UserPersonalInfo{
				UserID:     authresp.GetUser().GetId(),
				FirstName:  authresp.GetUser().GetUserinfo().GetFirstname(),
				SecondName: authresp.GetUser().GetUserinfo().GetSecondname(),
			},
			Credentials: domain.UserCredentials{
				UserID:   authresp.GetUser().GetId(),
				Email:    authresp.GetUser().GetCredentials().GetEmail(),
				Password: authreq.GetCredentials().GetPassword(),
			},
		},
		Token: authresp.GetTokenString(),
	}

	response.Created(w, signUpResponse)
}

func checkCredentials(req web.SignUpRequest) error {
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
