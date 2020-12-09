package kafka

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

var producer sarama.SyncProducer

var conn struct {
	KafkaBroker string
	KafkaTopic  string
}

func init() {
	if temp := os.Getenv("kafkahost"); temp != "" {
		conn.KafkaBroker = temp
	} else {
		log.Fatal("System:", "kafka broker not defined")
	}
	if temp := os.Getenv("kafkatopic"); temp != "" {
		conn.KafkaTopic = temp
	} else {
		log.Fatal("System:", "kafka topic not defined")
	}
	var err error
	producer, err = initProducer()
	if err != nil {
		log.Fatal("System: Kafka Producer Error:", err)
	}
}

type Message struct {
	Notify string `json:"Notification"`
}

func initProducer() (sarama.SyncProducer, error) {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	prd, err := sarama.NewSyncProducer([]string{conn.KafkaBroker}, config)
	if err != nil {
		log.Fatal(err)
	}
	return prd, err
}

func Publish(key string) {
	// publish sync
	message := Message{"key:" + key + " was reloaded from database"}
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		log.Println("Kafka: json error:", err)
		return
	}
	msg := &sarama.ProducerMessage{
		Topic:     conn.KafkaTopic,
		Value:     sarama.StringEncoder(jsonBytes),
		Partition: 0,
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		log.Println("Kafka: Error publish:", err)
		return
	}
	log.Println("Kafka: Partition:", p, "Offset:", o, "key", key)
}
