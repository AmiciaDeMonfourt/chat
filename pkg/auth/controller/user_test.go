package controller

import (
	"pawpawchat/generated/proto/authpb"
	"pawpawchat/internal/dto"
	"pawpawchat/internal/model/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_decodeUser(t *testing.T) {
	user := &domain.User{
		Biography: domain.UserBiography{
			FirstName: "xui",
		},
		Credentials: domain.UserCredentials{},
	}

	created := &authpb.User{}
	assert.NoError(t, dto.EncodeUser(user, created))

	assert.NotNil(t, created)
}
