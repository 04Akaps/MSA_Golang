package services

import (
	"context"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type AmqpServiceImpl struct {
	channel *amqp.Channel
	ctx     context.Context
}

func NewAmqpService(ctx context.Context, channel *amqp.Channel) Amqp {
	return &AmqpServiceImpl{
		ctx:     ctx,
		channel: channel,
	}
}

func (Ai *AmqpServiceImpl) GetTest() {
	msgs, err := Ai.channel.Consume("my_queue", "", false, false, false, false, nil)
	if err != nil {
		log.Fatal("Wrong", err)
	}

	for msg := range msgs {
		fmt.Println("message receiver : ", string(msg.Body))
		msg.Ack(false)
	}
}
