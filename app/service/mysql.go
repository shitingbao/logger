package service

import (
	"encoding/json"
	"log"
	"logger/app/logdriver"
	"logger/app/model"
	"logger/lib/hook"

	"github.com/sirupsen/logrus"
)

// MysqlLogHook，继承基础钩子
type MysqlLogHook struct {
	*hook.BashHook
}

// NewMysqlHook
func NewMysqlHook(level logrus.Level, h hook.HookFireFunc) *MysqlLogHook {
	return &MysqlLogHook{hook.NewBashHook(level, h)}
}

// mysqlSyncFireFunc 实际逻辑操作,入库定义,func(host string, entry *logrus.Entry) error
func (s *Service) mysqlSyncFireFunc(entry *logrus.Entry) error {
	da, _ := json.Marshal(entry.Data)
	m := model.LogMessgae{LogTime: entry.Time, Topic: entry.Message, Mes: string(da), Level: entry.Level}
	if err := s.MysqlDB.Table("log").Select("log_time", "topic", "mes", "host", "level").Create(&m).Error; err != nil {
		log.Println("into sql:", err) // 待定
		return err
	}
	return nil
}
func (s *Service) MysqlHookInit(user, pas, host, dataBase, port string) logrus.Hook {
	mysql, err := logdriver.OpenMysql(user, pas, host, dataBase, port)
	if err != nil {
		panic(err)
	}
	s.MysqlDB = mysql
	lh := NewMysqlHook(logrus.DebugLevel, s.mysqlSyncFireFunc)
	return lh
}
