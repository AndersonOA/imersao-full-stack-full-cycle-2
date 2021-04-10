package main

import (
	"fmt"
	"github.com/AndersonOA/imersao-full-stack-full-cycle-2-simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main()  {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	//producer := kafka.NewKafkaProducer()
	//kafka.Publish("ola", "readtest", producer)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
	}

	//r := route.Route{
	//	ID: "1",
	//	ClientID: "1",
	//}
	//
	//r.LoadPositions()
	//stringJson, _ := r.ExportJsonPositions()
	//fmt.Println(stringJson[0])
}

