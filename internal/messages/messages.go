package messages

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math/rand"
	"time"

	logger "github.com/Black-tag/kafka-sampler/internal/logging"
)

type Message struct {
	Event   string `json:"event" xml:"event"`
	UserID  int    `json:"user_id" xml:"user_id"`
	Payload string `json:"payload" xml:"payload"`
}

var events = []string{
	"user_signup",
	"user_login",
	"user_deletion",
	"order_created",
	"order_dispatched",
	"order_canceled",
	"order_confirmed",
}

var payload = []string{
	"{\"data\": \"user created successfully\"}",
	"{\"data\": \"payment failed due to card decline\"}",
	"{\"data\": \"item added to cart\"}",
	"{\"data\": \"user logged out\"}",
	"{\"data\": \"order shipped\"}",
}

func randomMessage() Message {
	rand.Seed(time.Now().UnixNano())
	return Message{
		Event:   events[rand.Intn(len(events))],
		UserID:  rand.Intn(1000),
		Payload: payload[rand.Intn(len(payload))],
	}
}

func GenerateJSONMesssage() ([]byte, error) {
	logger.Log.Info("enetered json message generator function")

	msg := randomMessage()
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("could not marshal json: %v", err)
	}
	return data, nil

}

func GenerateXMLMessage() ([]byte, error) {
	logger.Log.Info("enetered xml message generator function")
	msg := randomMessage()
	data, err := xml.MarshalIndent(msg, "", " ")
	if err != nil {
		return nil, fmt.Errorf("could not marshal xml msg")

	}
	return data, nil
}


