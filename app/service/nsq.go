package service

// 直接将nsq挂载在server上，只要server实现nsq接口即可

import (
	"encoding/json"
	conf "logger/app/config"
	"logger/lib/snsq"

	"github.com/nsqio/go-nsq"
)

// 实现HandleMessage方法,为了实现nsq方法
// 消费者逻辑
func (s *Service) HandleMessage(message *nsq.Message) error {
	m := &snsq.NsqMes{}
	if err := json.Unmarshal(message.Body, m); err != nil {
		panic(err)
	}
	return s.AntsPool.Submit(func() {
		s.LogInto(m)
	})
}
func (s *Service) NsqInit(c *conf.Config) {
	n := snsq.NsqConfig{
		Address: c.Nsq.Address,
		Topic:   c.Nsq.Topic,
		Channel: c.Nsq.Channel,
	}
	logSeverNsq, err := snsq.NsqLogServerLoad(s, n)
	if err != nil {
		panic(err)
	}
	s.LogSeverNsq = logSeverNsq
}
