package snsq

import (
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

//NewNsqCustomer 新建一个Customer,handle必须是实现了HandleMessage方法,内部连接，handle中接收数据
func NewNsqCustomer(tcpNsqdAddree, topic, channel string, hd nsq.Handler) (*nsq.Consumer, error) {
	con, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if err != nil {
		return nil, err
	}
	// defer con.Stop()
	con.AddHandler(hd)
	if err = con.ConnectToNSQD(tcpNsqdAddree); err != nil {
		return nil, err
	}
	logrus.WithFields(logrus.Fields{"tcpNsqdAddrr": tcpNsqdAddree}).Info("nsq")
	return con, nil
}
