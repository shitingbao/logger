package loader

import (
	conf "logger/app/config"
	"logger/app/server/grpc"
	"logger/app/server/http"
	"logger/app/service"
	"logger/lib/hook"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
)

func AutoLoader() {
	s := servers()
	lend := make(chan bool)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(*service.Service) {
		for range c {
			logrus.Info("all back job finished,now shutdown http server...")
			s.Shutdown()
			logrus.Info("success shutdown")
			lend <- true
			break
		}
	}(s)
	<-lend
}

func servers() *service.Service {
	config := conf.Init()
	lg := hook.NewLogues(config.Log.Level)
	s := service.NewService(lg, config)
	s.NsqInit(config)
	s.HookInit(config)
	go grpc.GrpcInit(config.Grpc.Port, s)
	go http.HttpInit(s)
	return s
}
