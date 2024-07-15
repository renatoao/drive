package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type RabbitMQConfig struct {
	URL       string
	TopicName string
	Timeout   time.Time
}

type RabbitMQConnection struct {
	cfg  RabbitMQConfig
	conn *amqp.Connection
}

func (rc *RabbitMQConnection) Publish(msg []byte) error {}

func (rc *RabbitMQConnection) Consume() error {}
