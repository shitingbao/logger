package service

import (
	"context"
	conf "logger/app/config"
	"logger/lib/pool"
	"logger/lib/snsq"

	"github.com/panjf2000/ants"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Service struct {
	Config      *conf.Config
	MongoDB     *mongo.Database
	MysqlDB     *gorm.DB
	Log         *logrus.Logger
	LogSeverNsq *snsq.LogSeverNsq
	AntsPool    *ants.Pool
}

// NewService新建服务并
func NewService(lg *logrus.Logger, config *conf.Config) *Service {
	s := &Service{
		Config:   config,
		Log:      lg,
		AntsPool: pool.NewAntsPool(),
	}
	return s
}

//
func (s *Service) Shutdown() {
	s.mongoClose()
	s.nsqClose()
	// mysqlClose gorm默认
	s.antsClose()
}
func (s *Service) mongoClose() {
	if s.MongoDB != nil {
		s.MongoDB.Client().Disconnect(context.Background())
	}
}
func (s *Service) nsqClose() {
	if s.LogSeverNsq != nil {
		s.LogSeverNsq.Stop()
	}
}
func (s *Service) antsClose() {
	if s.AntsPool != nil {
		s.AntsPool.Release()
	}

}
