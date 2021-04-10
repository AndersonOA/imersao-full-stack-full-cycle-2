package kafka

import (
	"encoding/json"
	"github.com/AndersonOA/imersao-full-stack-full-cycle-2-simulator/application/route"
	"github.com/AndersonOA/imersao-full-stack-full-cycle-2-simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"time"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	router := route.NewRoute()
	json.Unmarshal(msg.Value, &router)
	router.LoadPositions()
	positions, err := router.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
