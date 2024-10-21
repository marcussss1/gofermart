package orderv2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"gofermart/internal/models"
	"time"
)

func (r Repository) UploadOrder(ctx context.Context, order *models.OrderV2) error {
	// Сериализуем заказ в JSON
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return err
	}

	// Создаем ключ сообщения (можно использовать ID заказа, если он есть)
	key := sarama.StringEncoder(order.ID) // Предполагается, что у OrderV2 есть поле ID

	// Создаем сообщение
	msg := &sarama.ProducerMessage{
		Topic:     "orders",
		Key:       key,
		Value:     sarama.ByteEncoder(orderJSON),
		Timestamp: time.Now(),
	}

	// Отправляем сообщение
	_, _, err = r.client.Producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Println("succesfully uploaded order")
	return nil
}
