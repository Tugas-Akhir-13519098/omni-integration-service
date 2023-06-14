package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	KafkaHost               string `envconfig:"KAFKA_HOST" default:"localhost"`
	KafkaPort               string `envconfig:"KAFKA_PORT" default:"9092"`
	KafkaOrderTopic         string `envconfig:"KAFKA_ORDER_TOPIC" default:"order"`
	KafkaOrderConsumerGroup string `envconfig:"KAFKA_ORDER_CONSUMER_GROUP" default:"omni-order-consumer-group"`

	OmnichannelURL string `envconfig:"OMNICHANNEL_URL" default:"http://localhost:8080/api/v1/order"`
}

// Get to get defined configuration
func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)

	return cfg
}
