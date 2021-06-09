package hook

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func NewLogues(level int) *logrus.Logger {
	lg := logrus.New()
	lg.Level = logrus.Level(level)
	canleStandardOut(lg)
	return lg
}

// HookLoad 放入所有钩子，即对应钩子产生作用
func HookLoad(lg *logrus.Logger, hooks ...logrus.Hook) {
	for _, hook := range hooks {
		lg.AddHook(hook)
	}
}

// 取消标准输出
func canleStandardOut(lg *logrus.Logger) {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Open Src File err", err)
	}
	writer := bufio.NewWriter(src)
	lg.SetOutput(writer)
}
