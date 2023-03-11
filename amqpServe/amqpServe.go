package amqpserve

import (
	"GO_MSA/config"

	"github.com/streadway/amqp"
)

func GetAmqpConnection(envConfig config.Config) (*amqp.Connection, error) {
	connection, err := amqp.Dial(envConfig.AmqpUri)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func GetAmquChannel(a *amqp.Connection) (*amqp.Channel, error) {
	channel, err := a.Channel()
	if err != nil {
		return nil, err
	}

	return channel, nil
}
