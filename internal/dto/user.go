package dto

import (
	"encoding/json"
	"fmt"
	"io"
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/internal/model/domain"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// encodeUser encoding user to dst
func EncodeUser(user *domain.User, dst any) error {
	if err := validate.Struct(user); err != nil {
		return err
	}

	switch src := dst.(type) {
	case *authpb.User:
		*src = authpb.User{
			Id: 1,
			Biography: &authpb.Biography{
				Firstname:  user.Biography.FirstName,
				Secondname: user.Biography.SecondName,
			},
			Credentials: &authpb.Credentials{
				Email:    user.Credentials.Email,
				Password: user.Credentials.Password,
			},
		}
		return nil
	case *authpb.NewUser:
		*src = authpb.NewUser{
			Biography: &authpb.Biography{
				Firstname:  user.Biography.FirstName,
				Secondname: user.Biography.SecondName,
			},
			Credentials: &authpb.Credentials{
				Email:    user.Credentials.Email,
				Password: user.Credentials.Password,
			},
		}
		return nil
	default:
		return fmt.Errorf("incompatible user destination: %v", reflect.TypeOf(dst))
	}
}

func ExtractUser(from any) (*domain.User, error) {
	switch src := from.(type) {
	case io.ReadCloser:
		user := &domain.User{}
		if err := json.NewDecoder(src).Decode(user); err != nil {
			return nil, err
		}
		if err := validate.Struct(user); err != nil {
			return nil, err
		}
		return user, nil

	case *authpb.SignUpRequest:
		userRaw := src.GetUser()
		if userRaw == nil {
			return nil, fmt.Errorf("user is missing")
		}
		credentials, err := extractCredentials(userRaw)
		if err != nil {
			return nil, err
		}
		biography, err := extractBiogrphy(userRaw)
		if err != nil {
			return nil, err
		}
		return &domain.User{Credentials: *credentials, Biography: *biography}, nil

	default:
		return nil, fmt.Errorf("unknown request: %v", reflect.TypeOf(from))
	}
}

func extractCredentials(in any) (*domain.UserCredentials, error) {
	switch src := in.(type) {
	case *authpb.NewUser:
		credentialsRaw := src.GetCredentials()
		if credentialsRaw == nil {
			return nil, fmt.Errorf("credentials is missing")
		}

		credentials := &domain.UserCredentials{
			Email:    credentialsRaw.Email,
			Password: credentialsRaw.Password,
		}

		if credentials.Email == "" {
			return nil, fmt.Errorf("email is missing")

		} else if credentials.Password == "" {
			return nil, fmt.Errorf("password is missing")
		}
		return credentials, nil

	default:
		return nil, fmt.Errorf("incompatible user credentials source: %v", reflect.TypeOf(in))
	}
}

func extractBiogrphy(in any) (*domain.UserBiography, error) {
	switch src := in.(type) {
	case *authpb.NewUser:
		biographyRaw := src.GetBiography()
		if biographyRaw == nil {
			return nil, fmt.Errorf("biography is missing")
		}

		biography := &domain.UserBiography{
			FirstName:  biographyRaw.Firstname,
			SecondName: biographyRaw.Secondname,
		}

		if biography.FirstName == "" {
			return nil, fmt.Errorf("first name is missing")

		} else if biography.SecondName == "" {
			return nil, fmt.Errorf("second name is missing")
		}

		return biography, nil
	default:
		return nil, fmt.Errorf("incompatible biography source: %v", reflect.TypeOf(in))
	}
}
