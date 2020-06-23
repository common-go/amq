package amq

import (
	"context"
	"github.com/go-stomp/stomp"
)

type AMQHealthService struct {
	Conn *stomp.Conn
	name string
	des  string
}

func NewAMQHealthService(conn *stomp.Conn, name string, des string) *AMQHealthService {
	if len(name) == 0 {
		name = "amq"
	}
	if len(des) == 0 {
		des = "Test::Message"
	}
	return &AMQHealthService{conn, name, des}
}

func NewDefaultAMQHealthService(client *stomp.Conn, des string) *AMQHealthService {
	return &AMQHealthService{client, "amq", des}
}

func (s *AMQHealthService) Name() string {
	return s.name
}

func (s *AMQHealthService) Check(ctx context.Context) (map[string]interface{}, error) {
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

func (s *AMQHealthService) Build(ctx context.Context, data map[string]interface{}, err error) map[string]interface{} {
	if err == nil {
		return data
	}
	data["error"] = err.Error()
	return data
}
