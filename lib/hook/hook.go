package hook

// 资料参考 https://github.com/sohlich/elogrus
import (
	"github.com/sirupsen/logrus"
)

// log逻辑函数结构
type HookFireFunc func(entry *logrus.Entry) error

// 基础钩子结构，实现了logrus需要的fire和Levels
// host服务地址
// 等级设置，设置最低记录的等级，比他高的等级都记录，包括自己
// fireFunc逻辑函数
type BashHook struct {
	// host     string
	level    logrus.Level
	fireFunc HookFireFunc
}

func (m BashHook) Fire(enter *logrus.Entry) error {
	return m.fireFunc(enter)
}

// 这里传入等级,只有这里定义了等级,对应等级的日志才能触发fire
func (m BashHook) Levels() []logrus.Level {
	return setMysqlHookLevels(m.level)
}

// NewBashHook,反馈一个普通hook
func NewBashHook(level logrus.Level, f HookFireFunc) *BashHook {
	return newHook(level, f)
}

// NewAsyncMysqlHook反馈一个异步记录的hook,待定
func NewAsyncMysqlHook(level logrus.Level, f HookFireFunc) *BashHook {
	return newHook(level, f)
}

func newHook(level logrus.Level, f HookFireFunc) *BashHook {
	return &BashHook{
		level:    level,
		fireFunc: f,
	}
}

func setMysqlHookLevels(level logrus.Level) []logrus.Level {
	var levels []logrus.Level
	for _, l := range []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	} {
		if l <= level {
			levels = append(levels, l)
		}
	}
	return levels
}

// 异步执行
// func asyncFireFunc(host string, entry *logrus.Entry) error {
// 	go syncFireFunc(host, entry)
// 	return nil
// }

// type messgae struct {
// 	LogTime time.Time `gorm:"log_time"`
// 	Topic   string
// 	Mes     string
// 	Host    string
// 	Level   string
// }

// 实际逻辑操作,入库定义
// func syncFireFunc(host string, entry *logrus.Entry) error {
// 	// da, _ := json.Marshal(entry.Data)
// 	cf(entry)
// 	// m := messgae{LogTime: entry.Time, Topic: entry.Message, Mes: string(da), Host: host, Level: entry.Level.String()}
// 	// if err := client.Table("log").Select("log_time", "topic", "mes", "host", "level").Create(&m).Error; err != nil {
// 	// 	log.Println("into sql:", err)
// 	// 	return err
// 	// }
// 	return nil
// }
