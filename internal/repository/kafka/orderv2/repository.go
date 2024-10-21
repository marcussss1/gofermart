package orderv2

import "gofermart/pkg/kafka"

type Repository struct {
	client *kafka.ProducerClient
}

func NewRepository(client *kafka.ProducerClient) *Repository {
	return &Repository{
		client: client,
	}
}
