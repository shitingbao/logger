package snsq

// 可插拔式操作
// 实现一个基本HandleMessage(message *nsq.Message) error
// load之后，LogPulish发送即可
import (
	"encoding/json"
	"log"

	"github.com/nsqio/go-nsq"
)

// 默认使用这个topic
var (
	DefaultNsqTopic   = "logger"
	DefaultAddress    = "127.0.0.1:4150"
	DefaultChannel    = "log"
	DefaultLogVersion = "1.0.0"
)

type NsqMes struct {
	Sys   string
	Level int64
	Msg   NsqMesContent
}

// 注意这里的Topic指的是消息中的话题，不是nsq通道
type NsqMesContent struct {
	Host    string `json:"host,omitempty"`
	Topic   string `json:"topic"`
	Content string `json:"content"`
	Version string `json:"version"`
	LogTime string `json:"log_time"`
}

type NsqConfig struct {
	Address string `toml:"address"`
	Topic   string `toml:"topic"` // 为空就使用默认
	Channel string `toml:"channel"`
}

type LogSeverNsq struct {
	*nsq.Producer
}

type NsqLogServer interface {
	HandleMessage(message *nsq.Message) error
}

func NsqLogServerLoad(nServer NsqLogServer, c NsqConfig) (*LogSeverNsq, error) {
	if c.Address == "" {
		c.Address = DefaultAddress
	}
	if c.Topic == "" {
		c.Topic = DefaultNsqTopic
	}
	if c.Channel == "" {
		c.Channel = DefaultChannel
	}
	_, err := NewNsqCustomer(c.Address, c.Topic, c.Channel, nServer)
	if err != nil {
		return nil, err
	}
	client, err := NewNsqProducerClient(c.Address)
	if err != nil {
		return nil, err
	}
	return &LogSeverNsq{client}, nil
}

// topic参数代表了nsq通道，暂时没有使用，使用的是默认的，预留以后可以同一个客户端向多个通道发送
func (s *LogSeverNsq) LogPulish(body *NsqMes, topic ...string) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("LogPulish:", err)
		}
	}()
	switch {
	case body.Msg.Version == "":
		body.Msg.Version = DefaultLogVersion
	}
	topic = append(topic, DefaultNsqTopic)
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	return s.Publish(topic[0], b)
}
