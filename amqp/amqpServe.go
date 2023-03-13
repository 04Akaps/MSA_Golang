package amqpserve

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"GO_MSA/config"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type EventAmqp struct {
	amqp    *amqp.Connection
	channel *amqp.Channel
}

type amqpEvent struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	LocationId string    `json:"location"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
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

			rawEventName, ok := msg.Headers["x-event-name"]

			if !ok {
				err := msg.Nack(false, false)
				if err != nil {
					log.Fatal("Nack error : ", err)
				}
				// 만약 잘못되었다면 다른 구독자에게 전달도 안하지만 메시지는 표시하게
				continue
			}

			eventName, ok := rawEventName.(string)

			if !ok {
				err := msg.Nack(false, false)
				if err != nil {
					log.Fatal("Nack error : ", err)
				}
				continue
			}

			if eventName == "event.created" {
				// 후에 event를 새로 생성하는 요청도 들어갈 예정
				continue
			}

			var event amqpEvent

			err := json.Unmarshal(msg.Body, &event)
			if err != nil {
				log.Fatal("error unmarshalling : ", err)
			}

			fmt.Println(event)
			msg.Ack(false)

			err = msg.Ack(false)
			if err != nil {
				log.Fatal("Nack error : ", err)
			}
		}
	}()

	<-forever
}

func (Ea *EventAmqp) ServeHTTP(c *gin.Context) {
	// 보통은 이제 메시지를 파라메터로 받아서 처리하지만, 나는 어떤방식으로 동작하는지가 궁금했기때문에
	// value를 fix하여 테스트 진행

	eventValue := &amqpEvent{
		ID:         "1",
		Name:       "hojin",
		LocationId: "3",
		Start:      time.Now(),
	}

	jsonDoc, err := json.Marshal(eventValue)
	if err != nil {
		log.Fatal("jsonDoc error : ", err)
	}

	mseesage := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": eventValue.Name},
		Body:        jsonDoc,
		ContentType: "application/json",
	}

	err = Ea.channel.Publish("events", "sample-key", false, false, mseesage)
	if err != nil {
		log.Fatal("Error Exchange Declare", err)
	}
}
