package amqpserve

import (
	"fmt"
	"log"

	"GO_MSA/config"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type EventAmqp struct {
	amqp    *amqp.Connection
	channel *amqp.Channel
}

func GetAmqpConnection(envConfig config.Config) (*EventAmqp, error) {
	connection, err := amqp.Dial(envConfig.AmqpUri)
	if err != nil {
		fmt.Println("0000000", err)
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		fmt.Println("1111", err)
		return nil, err
	}

	// defer connection.Close()
	// defer channel.Close()

	return &EventAmqp{amqp: connection, channel: channel}, nil
}

func (Ea *EventAmqp) GetChannel() *amqp.Channel {
	return Ea.channel
}

func (Ea *EventAmqp) SetAmquChannel(name, queue string) error {
	err := Ea.channel.ExchangeDeclare(name, "topic", true, false, false, false, nil)
	if err != nil {
		return err
	}

	_, err = Ea.channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = Ea.channel.QueueBind(queue, "#", "events", false, nil)
	if err != nil {
		return err
	}

	return nil
}

func (Ea *EventAmqp) Listening() {
	forever := make(chan bool)

	msgs, err := Ea.channel.Consume("my_queue", "", false, false, false, false, nil)
	if err != nil {
		log.Fatal("Wrong", err)
	}

	go func() {
		for msg := range msgs {
			fmt.Println("message receiver : ", string(msg.Body))
			msg.Ack(false)
		}
	}()

	<-forever
}

func (Ea *EventAmqp) ServeHTTP(c *gin.Context) {
	mseesage := amqp.Publishing{
		Body: []byte("Heelo World"),
	}

	err := Ea.channel.Publish("events", "sample-key", false, false, mseesage)
	if err != nil {
		log.Fatal("Error Exchange Declare", err)
	}
}
