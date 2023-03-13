package amqpserve

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"GO_MSA/config"
	"GO_MSA/middleware"

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
	for {
		select {
		case newChannel := <-Ea.newChannelEvent:
			fmt.Println("새로운 채널이 들어옴")

			newChannelName := newChannel.channelName
			newChannelQueue := newChannel.channelQueueName

			SetAmquChannel(Ea, newChannelName, newChannelQueue)

			msgs, err := Ea.channel[newChannelName].Consume(newChannelQueue, "", false, false, false, false, nil)
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
		}
	}
}

type RequestEventAmqp struct {
	Name      string `uri:"name" binding:"required"`
	ExChanger string `uri:"changer" binding:"required"`
}

func (Ea *EventAmqp) ServeHTTP(ctx *gin.Context) {
	var req RequestEventAmqp

	bodyCheckError := middleware.CheckBodyBinding(&req, ctx)

	if bodyCheckError != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bodyCheckError, "status": -1})
		return
	}

	fmt.Println(req.Name)
	fmt.Println(req.ExChanger)

	// eventValue := &amqpEvent{
	// 	ID:         "1",
	// 	Name:       "hojin",
	// 	LocationId: "3",
	// 	Start:      time.Now(),
	// }

	// jsonDoc, err := json.Marshal(eventValue)
	// if err != nil {
	// 	log.Fatal("jsonDoc error : ", err)
	// }

	// mseesage := amqp.Publishing{
	// 	Headers:     amqp.Table{"x-event-name": eventValue.Name},
	// 	Body:        jsonDoc,
	// 	ContentType: "application/json",
	// }

	// err = Ea.channel.Publish("events", "sample-key", false, false, mseesage)
	// if err != nil {
	// 	log.Fatal("Error Exchange Declare", err)
	// }
}
