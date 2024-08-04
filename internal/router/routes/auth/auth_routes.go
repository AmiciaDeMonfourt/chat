package auth

import (
	"net/http"
	"pawpawchat/internal/producer"
)

type AuthRoutes interface {
	SignUp(http.ResponseWriter, *http.Request)
	SignIn(http.ResponseWriter, *http.Request)
}

type AuthRoutesImpl struct {
	producer *producer.Producer
}

func NewAuthRoutes() AuthRoutes {
	producer := producer.New("test-topic")
	go producer.StartProduce()

	return &AuthRoutesImpl{
		producer: producer,
	}
}
