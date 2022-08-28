package userconsumer

import (
	"context"
	// "encoding/binary"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
)

const (
	addTopic = "add_users"
	updateTopic = "update_users"
	deleteTopic = "delete_users"
)

type HandlerFunc func([]byte) error

type UserConsumer struct {
	userApp userapp.App
	handler map[string]HandlerFunc
}

func New(userApp userapp.App) *UserConsumer {
	handler := make(map[string]HandlerFunc)
	uC := &UserConsumer{
		userApp: userApp,
		handler: handler,
	}
	handler[addTopic] = uC.addUser
	handler[updateTopic] = uC.updateUser
	handler[deleteTopic] = uC.deleteUser

	return uC
}

func (c *UserConsumer) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *UserConsumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *UserConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Debug("Consumer start")
	for {
		select {
		case <-session.Context().Done():
			log.Info("Done")
			return nil
		case msg, ok := <-claim.Messages():
			if !ok {
				log.Info("Data channel closed")
				return nil
			}
			log.Infof("Handle msg: partition: %v, topic: %v, data: %v", msg.Partition, msg.Topic, string(msg.Value))
			err := c.handler[msg.Topic](msg.Value)
			if err != nil {
				log.Error(err)
			}
			log.Info("Done")
			session.MarkMessage(msg, "")
		}
	}
}

func (c *UserConsumer) addUser(msgValue []byte) error {
	var u user.User
	if err := json.Unmarshal(msgValue, &u); err != nil {
		return err
	}

	return c.userApp.Add(context.Background(), u)
}

func (c *UserConsumer) updateUser(msgValue []byte) error {
	var u user.User
	if err := json.Unmarshal(msgValue, &u); err != nil {
		return err
	}

	return c.userApp.Update(context.Background(), u)
}

func (c *UserConsumer) deleteUser(msgValue []byte) error {
	var u user.User
	if err := json.Unmarshal(msgValue, &u); err != nil {
		return err
	}
	id := u.GetId()
	return c.userApp.Delete(context.Background(), id)
}