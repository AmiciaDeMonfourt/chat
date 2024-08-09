package orm

import (
	"context"
	"fmt"
	"pawpawchat/internal/model/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_GormUserRepository_Create(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}))
	assert.NoError(t, err)

	userRepository := NewGormUserRepository(gormDB)

	users := []domain.User{
		{
			Credentials: domain.UserCredentials{
				Email:    "email1@example.com",
				HashPass: "hashpass",
			},
		},
		{
			Credentials: domain.UserCredentials{
				Email:    "email2@example.com",
				HashPass: "hashpass",
			},
		},
	}

	for idx := range users {
		mock.ExpectBegin()

		mock.ExpectQuery(`INSERT INTO "user_credentials" \("email","hash_pass"\) VALUES \(\$1,\$2\) RETURNING \"user_credentials"."user_id"`).
			WithArgs(users[idx].Credentials.Email, users[idx].Credentials.HashPass).
			WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(idx))

		mock.ExpectCommit()
	}

	assert.NoError(t, userRepository.Create(context.TODO(), &users[0].Credentials))
	assert.NoError(t, userRepository.Create(context.TODO(), &users[1].Credentials))

	assert.NoError(t, mock.ExpectationsWereMet())

	fmt.Println(users[0].Credentials)
	fmt.Println(users[1].Credentials)
}
