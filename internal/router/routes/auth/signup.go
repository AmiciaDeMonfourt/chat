package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"pawpawchat/internal/model/domain/user"
	"pawpawchat/internal/model/web/auth"
	"pawpawchat/internal/producer"
	"pawpawchat/utils/jwt"
	"pawpawchat/utils/response"
)

func (r *AuthRoutesImpl) SignUp(w http.ResponseWriter, req *http.Request) {
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

	// created user logic (refactor [add record at the database])
	newUser := user.New(1, signUpRequest.FirstName, signUpRequest.SecondName)

	// write new record to the topic
	if err := writeToTopic(r.producer, newUser); err != nil {
		response.InternalErr(w, err.Error())
		return
	}

	// generated jwt token
	token, err := jwt.GenerateToken(newUser.ID)
	if err != nil {
		response.InternalErr(w, err.Error())
		return
	}

	// http response
	signUpResponse := &auth.SignUpResponse{
		User:  newUser,
		Token: token,
	}

	response.Created(w, signUpResponse)
}

func writeToTopic(producer *producer.Producer, newUser user.User) error {
	value, err := json.Marshal(newUser)
	if err != nil {
		return err
	}

	key, err := json.Marshal(newUser.ID)
	if err != nil {
		return err
	}

	return producer.Write(value, key)
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
