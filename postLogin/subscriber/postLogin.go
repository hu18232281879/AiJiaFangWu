package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	postLogin "microProject/postLogin/proto/postLogin"
)

type PostLogin struct{}

func (e *PostLogin) Handle(ctx context.Context, msg *postLogin.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *postLogin.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
