package kafkaconsumer

import (
	"context"
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/kanhaiya15/gopf/lib/logging/gopflogrus"
	"github.com/kanhaiya15/gopf/utils"
	"github.com/segmentio/kafka-go"
)

var logger = gopflogrus.NewLogger()

// InitConsumer InitConsumer
func InitConsumer(brokers []string) {
	logger.Debug(" Kafka Consumer -starts")
	topic, err := utils.GetConfValue("KAFKA_TOPIC")
	if err != nil {
		panic(err.Error())
	}
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	//New Reader will read from last position
	r.SetOffset(kafka.LastOffset)
	//var raw  map[string]interface{}

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			logger.Errorf("Error in Reading Kafka Message : %+v", err)
			continue
		}
		jsonMap := make(map[string]interface{})
		err = json.Unmarshal(m.Value, &jsonMap)
		if err != nil {
			logger.Errorf("Error in Unmarshalling data: %+v", err)
			continue
		}

		Client := jsonMap["client"].(string)
		TestID := jsonMap["testId"].(string)
		Status := jsonMap["status"].(string)
		UserID := "0"

		logger.Infof("Kafka Client : %s, TestID : %s, Status: %s, UserID : %s", Client, TestID, Status, UserID)

		switch reflect.TypeOf(jsonMap["userId"]).Name() {
		case "string":
			UserID = jsonMap["userId"].(string)
		case "float64":
			UserID = strconv.Itoa(int(jsonMap["userId"].(float64)))
		case "int64":
			UserID = strconv.FormatInt(jsonMap["userId"].(int64), 10)
		case "int":
			UserID = strconv.Itoa(jsonMap["userId"].(int))
		}

		logger.Infof("Kafka Message is: %+v", jsonMap)
	}
}
