package service

import (
	conf "logger/app/config"
	"logger/lib/hook"
	"logger/lib/snsq"

	"github.com/sirupsen/logrus"
)

const (
	PanicLevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel

	logMes = "mes"
)

// log hook的启动关系到对应db的启动，通过日志开启,注入logurs
func (s *Service) HookInit(config *conf.Config) {
	hooks := []logrus.Hook{}
	if config.MongoDB.IsLogOpen {
		hooks = append(hooks, s.MongoHooKInit(config.MongoDB.Driver, config.MongoDB.Database))
	}
	if config.MysqlDB.IsLogOpen {
		lh := s.MysqlHookInit(config.MysqlDB.User, config.MysqlDB.Password, config.MysqlDB.Host, config.MysqlDB.Database, config.MysqlDB.Port)
		hooks = append(hooks, lh)
	}
	hook.HookLoad(s.Log, hooks...)
}

// log入库函数
// 不应该直接调用，应该调用LogPulish，通过队列提交
func (s *Service) LogInto(n *snsq.NsqMes) {
	etr := s.Log.WithFields(logrus.Fields{logMes: n.Msg})
	switch n.Level {
	case PanicLevel:
		panicExtr(n.Sys, etr)
	case FatalLevel:
		etr.Debug(n.Sys)
	case ErrorLevel:
		etr.Error(n.Sys)
	case WarnLevel:
		etr.Warn(n.Sys)
	case InfoLevel:
		etr.Info(n.Sys)
	case DebugLevel:
		etr.Debug(n.Sys)
	default:
		etr.Trace(n.Sys) //等级不够这句不会输出
	}
}

// 先处理掉这个panic，为了不直接在内部出现panic，同时可以进行http反馈
// 只是为了拦截panic，不用处理该err
func panicExtr(title string, etr *logrus.Entry) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	etr.Panic(title)
}
