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
	amqp            *amqp.Connection
	channel         map[string]*amqp.Channel
	newChannelEvent chan *NewChannelEvent
}

type NewChannelEvent struct {
	channelName      string
	channelQueueName string
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

	// defer connection.Close()
	// defer channel.Close()

	newEventAmqp := &EventAmqp{amqp: connection, channel: make(map[string]*amqp.Channel), newChannelEvent: make(chan *NewChannelEvent)}

	newEventAmqp.newChannelEvent <- &NewChannelEvent{
		channelName:      "events",
		channelQueueName: "my_queue",
	}

	return newEventAmqp, nil
}

func (Ea *EventAmqp) GetChannel(name string) *amqp.Channel {
	return Ea.channel[name]
}

func SetAmquChannel(Ea *EventAmqp, name, queue string) error {
	channel, err := Ea.amqp.Channel()
	if err != nil {
		return err
	}

	Ea.channel[name] = channel

	Ea.newChannelEvent <- &NewChannelEvent{
		channelName:      name,
		channelQueueName: queue,
	}

	err = Ea.channel[name].ExchangeDeclare(name, "topic", true, false, false, false, nil)
	if err != nil {
		return err
	}

	_, err = Ea.channel[name].QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = Ea.channel[name].QueueBind(queue, "#", "events", false, nil)
	if err != nil {
		return err
	}

	return nil
}

func (Ea *EventAmqp) Listening() {
	forever := make(chan bool)

	for {
		select {
		case newChannel := <-Ea.newChannelEvent:

			newChannelName := newChannel.channelName
			newChannelQueue := newChannel.channelQueueName

			SetAmquChannel(Ea, newChannelName, newChannelQueue)
		}
	}

	for i := 0; i < len(Ea.channelNameList); i++ {
		channelName := Ea.channelNameList[i]
	}

	msgs, err := Ea.channel[channelName].Consume("my_queue", "", false, false, false, false, nil)
	if err != nil {
		log.Fatal("Wrong", err)
	}

	go func() {
		for msg := range msgs {

			rawEventName, ok := msg.Headers["x-event-name"]

			if !ok {
				msg.Nack(false, false)
				// 만약 잘못되었다면 다른 구독자에게 전달도 안하지만 메시지는 표시하게
				continue
			}

			eventName, ok := rawEventName.(string)

			if !ok {
				msg.Nack(false, false)
				continue
			}

			fmt.Println(eventName)

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
