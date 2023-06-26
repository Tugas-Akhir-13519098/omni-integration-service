package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	KafkaHost               string `envconfig:"KAFKA_HOST" default:"localhost"`
	KafkaPort               string `envconfig:"KAFKA_PORT" default:"9092"`
	KafkaOrderTopic         string `envconfig:"KAFKA_ORDER_TOPIC" default:"order"`
	KafkaErrorTopic         string `envconfig:"KAFKA_ERROR_TOPIC" default:"error"`
	KafkaOrderConsumerGroup string `envconfig:"KAFKA_ORDER_CONSUMER_GROUP" default:"omni-order-consumer-group"`

	OmnichannelURL string `envconfig:"OMNICHANNEL_URL" default:"http://localhost:8080/api/v1/order"`
	AdminToken     string `envconfig:"ADMIN_TOKEN" default:"eyJhbGciOiJSUzI1NiIsImtpZCI6IjY3YmFiYWFiYTEwNWFkZDZiM2ZiYjlmZjNmZjVmZTNkY2E0Y2VkYTEiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vb21uaS00MGY5MSIsImF1ZCI6Im9tbmktNDBmOTEiLCJhdXRoX3RpbWUiOjE2ODcwMTgwNDQsInVzZXJfaWQiOiJTajZiT0JQSFVXV3BDZkNab3E0TEpnUHRmY0QzIiwic3ViIjoiU2o2Yk9CUEhVV1dwQ2ZDWm9xNExKZ1B0ZmNEMyIsImlhdCI6MTY4NzAxODA0NCwiZXhwIjoxNjg3MDIxNjQ0LCJlbWFpbCI6InNob3AxQGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJlbWFpbCI6WyJzaG9wMUBnbWFpbC5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCJ9fQ.Vo2zr_9loYrgikwbkh3LT6WLYo2cmkAAvitCic6qcRXR5VI_C5LEwKVs5HXM8wa9WwQ3IofuyClKbx_ff9W_eeeTnEwqZR34PPzHcECYBgJcsCBLQs7Gekr8dSU6-Vy4ivqqKgef9T6ub-0J_foXSubjP0iK6zXh0aHwiowgRZVvr-D1RvkUmcJTYhubG9l3zxgJyqmZIfPH7NCh9cn3GgSzf-3Kc113lizWQT_t_gtHiAG986v9NsU-uU16MN-BoZYpKV7qQ8mH1OohuPk8PUzZ3pxw9fRsSlOz10rwzgwwY5b9piERz2FypwYXoHepjPwGYnSg-5R9ImpwSlC1Yg"`
}

// Get to get defined configuration
func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)

	return cfg
}
