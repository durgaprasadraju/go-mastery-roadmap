package solutions

import "time"

// Message represents an event on a bus.
type Message struct {
	Topic     string
	Key       string
	Payload   []byte
	Timestamp time.Time
}

// NewMessage creates a message with current timestamp.
func NewMessage(topic, key string, payload []byte) Message {
	return Message{Topic: topic, Key: key, Payload: payload, Timestamp: time.Now()}
}