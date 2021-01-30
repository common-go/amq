package amq

import (
	"context"
	"github.com/go-stomp/stomp"
	"github.com/go-stomp/stomp/frame"
)

type Producer struct {
	Conn        *stomp.Conn
	Destination string
	ContentType string
}

func NewProducer(client *stomp.Conn, destinationName string, subscriptionName string, contentType string) *Producer {
	des := destinationName + "::" + subscriptionName
	if len(contentType) == 0 {
		contentType = "text/plain"
	}
	return &Producer{Conn: client, Destination: des, ContentType: contentType}
}

func NewProducerByConfig(c Config, contentType string) (*Producer, error) {
	conn, err := NewConn(c.UserName, c.Password, c.Addr)
	if err != nil {
		return nil, err
	}
	return NewProducer(conn, c.DestinationName, c.SubscriptionName, contentType), nil
}

func (p *Producer) Produce(ctx context.Context, data []byte, messageAttributes map[string]string) (string, error) {
	opts := MapToFrame(messageAttributes)
	err := p.Conn.Send(p.Destination, p.ContentType, data, opts...)
	return "", err
}

func MapToFrame(messageAttributes map[string]string) []func(*frame.Frame) error {
	opts := make([]func(*frame.Frame) error, 0)
	if messageAttributes != nil {
		for k, v := range messageAttributes {
			opts = append(opts, stomp.SendOpt.Header(k, v))
		}
	}
	return opts
}
