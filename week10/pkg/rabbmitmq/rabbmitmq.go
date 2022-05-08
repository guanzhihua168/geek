package rabbmitmq

import (
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var RabbitMQDefaultConnction *amqp.Connection
var RabbitMQDefault *RabbitMQ

type RabbitMQExchnageType struct {
	exchange string
}

func RabbitMQConnCreate(host, user, password string) (*amqp.Connection, error) {
	c, e := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", user, password, host))
	if e != nil {
		logrus.WithFields(logrus.Fields{
			"host": host,
		}).Error("Create Connection Fail:" + e.Error())
		return nil, e
	}
	return c, nil
}

type RabbitMQ struct {
	conn *amqp.Connection
}

func NewRabbitMQ(conn *amqp.Connection) (*RabbitMQ, error) {
	return &RabbitMQ{conn: conn}, nil
}

func (c *RabbitMQ) Channel() (*amqp.Channel, error) {
	return c.conn.Channel()
}
