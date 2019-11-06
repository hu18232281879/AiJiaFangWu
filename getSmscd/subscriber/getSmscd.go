package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	getSmscd "microProject/getSmscd/proto/getSmscd"
)

type GetSmscd struct{}

func (e *GetSmscd) Handle(ctx context.Context, msg *getSmscd.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *getSmscd.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
