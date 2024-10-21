package orders

import (
	"context"
	"gofermart/pkg/kafka"
)

type Worker struct {
	client *kafka.KafkaConsumer
}

func NewWorker(client *kafka.KafkaConsumer) *Worker {
	return &Worker{
		client: client,
	}
}

func (w *Worker) Start() {
	err := w.client.Consume(context.TODO())
	if err != nil {
		panic(err)
	}
}
