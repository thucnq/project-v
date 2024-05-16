package basemessage

import (
	"encoding/json"
	"time"
)

// Message ...
type Message struct {
	ExchangeName   string           `json:"exchangeName,omitempty"`
	RoutingKey     string           `json:"routingKey,omitempty"`
	Topic          string           `json:"topic,omitempty"`
	IssueAt        int64            `json:"issueAt,omitempty"`
	Issuer         string           `json:"issuer,omitempty"`
	MessageVersion string           `json:"messageVersion,omitempty"`
	RawMessage     *json.RawMessage `json:"message,omitempty"`
}

// Marshal ...
func (o Message) Marshal() ([]byte, error) {
	return json.Marshal(o)
}

func (o Message) GetExchangeName() string {
	return o.ExchangeName
}

func (o Message) GetTopic() string {
	return o.Topic
}
func (o Message) GetRoutingKey() string {
	return o.RoutingKey
}

// SetIssueAtNow ...
func (o Message) SetIssueAtNow() Message {
	o.IssueAt = time.Now().Unix()
	return o
}

func (o Message) WithExchangeName(exchangeName string) Message {
	o.ExchangeName = exchangeName
	return o
}

func (o Message) WithRoutingKey(routingKey string) Message {
	o.RoutingKey = routingKey
	return o
}

func (o Message) WithTopic(topic string) Message {
	o.Topic = topic
	return o
}
func (o Message) WithRawMsg(rawMessage json.RawMessage) Message {
	o.RawMessage = &rawMessage
	return o
}
