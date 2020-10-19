package amq

import (
	"context"

	"github.com/go-stomp/stomp"
)

type AMQHealthCheck struct {
	Conn *stomp.Conn
	name string
	des  string
}

func NewAMQHealthCheck(conn *stomp.Conn, name string, des string) *AMQHealthCheck {
	if len(name) == 0 {
		name = "amq"
	}
	if len(des) == 0 {
		des = "Test::Message"
	}
	return &AMQHealthCheck{conn, name, des}
}

func NewDefaultAMQHealthCheck(client *stomp.Conn, des string) *AMQHealthCheck {
	return &AMQHealthCheck{client, "amq", des}
}

func (s *AMQHealthCheck) Name() string {
	return s.name
}

func (s *AMQHealthCheck) Check(ctx context.Context) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	subscription, err := s.Conn.Subscribe(s.des, stomp.AckAuto,
		stomp.SubscribeOpt.Header("subscription-type", "ANYCAST"),
	)
	if err != nil {
		return res, err
	}
	err = subscription.Unsubscribe()
	res["version"] = s.Conn.Version()
	return res, err
}

func (s *AMQHealthCheck) Build(ctx context.Context, data map[string]interface{}, err error) map[string]interface{} {
	if err == nil {
		return data
	}
	data["error"] = err.Error()
	return data
}
