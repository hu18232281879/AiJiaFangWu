package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	getImagecd "microProject/getImagecd/proto/getImagecd"
)

type GetImagecd struct{}

func (e *GetImagecd) Handle(ctx context.Context, msg *getImagecd.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *getImagecd.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
