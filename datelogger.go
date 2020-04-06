package datelogger /*一个按日期存放的日志系统*/

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/Sirupsen/logrus"
)

//DateLogger log对象
type DateLogger struct {
	Path        string //存放的目录名
	Level       logrus.Level
	log         *logrus.Logger
	logFile     *os.File
	logFileName string //当前的日志文件名
}

//NewDateLog 反馈一个log对象
func NewDateLog(pathName string) *DateLogger {
	return &DateLogger{Path: pathName, Level: logrus.DebugLevel}
}

func (d *DateLogger) checkLogFile() error {
	strPath := filepath.Join(d.Path, time.Now().Format("2006-01-02")+".txt")
	//因为上面获取的是日期，每一天一个文件，会保存上次的目录地址，这里就判断是否需要创建新的文件
	if strPath != d.logFileName {
		if d.logFile != nil {
			if err := d.logFile.Close(); err != nil {
				return err
			}
		} else {
			d.log = &logrus.Logger{
				Formatter: &logrus.TextFormatter{
					TimestampFormat: "20060102T150405",
				},
				Hooks: make(logrus.LevelHooks),
				Level: d.Level,
			}
		}
		if d.log == nil {
			d.log = &logrus.Logger{
				Formatter: &logrus.TextFormatter{
					TimestampFormat: "20060102T150405",
				},
				Hooks: make(logrus.LevelHooks),
				Level: d.Level,
			}
		}
		d.logFileName = strPath
		//确保创建目录
		if err := os.MkdirAll(filepath.Dir(d.logFileName), os.ModePerm); err != nil {
			return err
		}
		flog, err := os.OpenFile(d.logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		d.logFile = flog
		d.log.Out = io.MultiWriter(flog, os.Stdout)
	}
	return nil
}

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func (d *DateLogger) WithError(err error) *logrus.Entry {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	return d.log.WithError(err)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func (d *DateLogger) WithField(key string, value interface{}) *logrus.Entry {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	return d.log.WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func (d *DateLogger) WithFields(fields logrus.Fields) *logrus.Entry {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	return d.log.WithFields(fields)
}

// Debug logs a message at level Debug on the standard logger.
func (d *DateLogger) Debug(args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func (d *DateLogger) Print(args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Print(args...)

}

// Info logs a message at level Info on the standard logger.
func (d *DateLogger) Info(args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func (d *DateLogger) Warn(args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func (d *DateLogger) Warning(args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Warning(args...)
}

// Error logs a message at level Error on the standard logger.
func (d *DateLogger) Error(args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func (d *DateLogger) Panic(args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func (d *DateLogger) Fatal(args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Fatal(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func (d *DateLogger) Debugf(format string, args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func (d *DateLogger) Printf(format string, args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func (d *DateLogger) Infof(format string, args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func (d *DateLogger) Warnf(format string, args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func (d *DateLogger) Warningf(format string, args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func (d *DateLogger) Errorf(format string, args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func (d *DateLogger) Panicf(format string, args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger.
func (d *DateLogger) Fatalf(format string, args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Fatalf(format, args...)
}

// Debugln logs a message at level Debug on the standard logger.
func (d *DateLogger) Debugln(args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Debugln(args...)
}

// Println logs a message at level Info on the standard logger.
func (d *DateLogger) Println(args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Println(args...)
}

// Infoln logs a message at level Info on the standard logger.
func (d *DateLogger) Infoln(args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func (d *DateLogger) Warnln(args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Warnln(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func (d *DateLogger) Warningln(args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Warningln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func (d *DateLogger) Errorln(args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Errorln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func (d *DateLogger) Panicln(args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Panicln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger.
func (d *DateLogger) Fatalln(args ...interface{}) {
	if err1 := d.checkLogFile(); err1 != nil {
		logrus.Panic(err1)
	}
	d.log.Fatalln(args...)
}
func (d *DateLogger) Close() error {
	if d.logFile != nil {
		err := d.logFile.Close()
		d.log = nil
		d.logFile = nil
		return err
	} else {
		return nil
	}
}
