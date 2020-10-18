package mail

import (
	"context"
	"encoding/json"
)

type Producer interface {
	Produce(ctx context.Context, data []byte, attributes *map[string]string) (string, error)
}

type MQMailSender struct {
	Producer   Producer
	Goroutines bool
}

func NewMQMailSender(producer Producer, goroutines bool) *MQMailSender {
	return &MQMailSender{Producer: producer, Goroutines: goroutines}
}
func (s *MQMailSender) Send(m SimpleMail) error {
	if s.Goroutines {
		go Produce(s.Producer, m)
		return nil
	} else {
		return Produce(s.Producer, m)
	}
}

func Produce(producer Producer, m SimpleMail) error {
	b, er1 := json.Marshal(m)
	if er1 == nil {
		_, er2 := producer.Produce(context.TODO(), b, nil)
		return er2
	} else {
		return er1
	}
}
