package queue

import (
	"fmt"
	"log"
	"reflect"
)

type QueueType int

const (
	RabbitMQ QueueType = iota
)

func New(qt QueueType, cfg any) (q *Queue, err error) {
	q := new(Queue)
	rt := reflect.TypeOf(cfg)
	switch qt {
	case RabbitMQ:
		if rt.Name() != "RabbitMQConfig" {
			return nil, fmt.Errorf("Config need's to be type of RabbitMQConfig")
		}
		fmt.Println("nao implementado")
	default:
		log.Fatal("tipo nao implementado")
	}

	return q
}

type QueueConnection interface {
	Publish([]byte) error
	Consume() error
}

type Queue struct {
	cfg any
	qc  QueueConnection
}

func (q *Queue) Publish(msg []byte) error {
	return q.qc.Publish(msg)
}

func (q *Queue) Consume() error {
	return q.qc.Consume()
}
