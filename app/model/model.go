package model

import (
	"time"

	"github.com/sirupsen/logrus"
)

// LogMessgae mysql入库model
type LogMessgae struct {
	LogTime time.Time `gorm:"log_time"`
	Topic   string
	Mes     string
	Host    string
	Level   logrus.Level
}

type ArgLogCondition struct {
	LogSys       string    `form:"log_sys" json:"log_sys"`
	LogStartTime time.Time `form:"log_start_time" json:"log_start_time"`
	LogEndTime   time.Time `form:"log_end_time" json:"log_end_time"`
	LogLevel     string    `form:"log_level" json:"log_level"` //逗号隔开
	Topic        string    `form:"topic" json:"topic"`
	Content      string    `form:"content" json:"content"`
	Page         int64     `form:"page" json:"page"`
	PageSize     int64     `form:"page_size" json:"page_size"`
	Order        ArgOrder  `form:"order" json:"order"`
}

type ArgOrder struct {
	OrderField string `form:"order_field" json:"order_field"`
	OrderVal   int    `form:"order_val" json:"order_val"`
}
