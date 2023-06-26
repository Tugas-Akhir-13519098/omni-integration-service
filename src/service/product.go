package service

import (
	"context"
	"encoding/json"
	"fmt"
	"omni-integration-service/config"
	"omni-integration-service/src/model"
	"omni-integration-service/src/util"

	"github.com/segmentio/kafka-go"
)

type OrderService interface {
	ConsumeOrderMessages()
}

type orderService struct{}

func NewOrderService() OrderService {
	return &orderService{}
}

func (ps *orderService) ConsumeOrderMessages() {
	// Set up the Kafka reader for order topic
	cfg := config.Get()
	config := kafka.ReaderConfig{
		Brokers: []string{fmt.Sprintf("%s:%s", cfg.KafkaHost, cfg.KafkaPort)},
		Topic:   cfg.KafkaOrderTopic,
		GroupID: cfg.KafkaOrderConsumerGroup,
	}
	reader := kafka.NewReader(config)

	// Continuously read messages from Kafka
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error reading message from Kafka:", err.Error())
			continue
		}

		// Change kafka message from byte to struct
		var kafkaOrderMessage model.KafkaOrderMessage
		err = json.Unmarshal(msg.Value, &kafkaOrderMessage)
		if err != nil {
			fmt.Println("Can't unmarshal the kafka message")
			continue
		}

		if kafkaOrderMessage.Method == model.CREATE {
			createOrderBody := util.ConvertKafkaOrderMessageToCreateOrderRequest(kafkaOrderMessage)
			url := cfg.OmnichannelURL
			resp, _ := util.SendPostRequest(createOrderBody, url)
			util.AfterHTTPRequestHandler(createOrderBody.String(), resp, "CREATE", "POST", kafkaOrderMessage.TokopediaOrderID, kafkaOrderMessage.ShopeeOrderID)

		} else { // orderMessage.Method == model.UPDATE
			updateOrderBody := util.ConvertKafkaOrderMessageToUpdateOrderStatusRequest(kafkaOrderMessage)
			url := cfg.OmnichannelURL
			resp, _ := util.SendPutRequest(updateOrderBody, url)
			util.AfterHTTPRequestHandler(updateOrderBody.String(), resp, "UPDATE", "PUT", kafkaOrderMessage.TokopediaOrderID, kafkaOrderMessage.ShopeeOrderID)
		}
	}
}
