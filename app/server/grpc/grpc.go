package grpc

import (
	"logger/app/api"
	"logger/app/service"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func GrpcInit(port string, serve *service.Service) {
	externalServer(port, serve)
}

func externalServer(port string, serve *service.Service) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	api.RegisterLogServerServer(s, serve)
	logrus.WithFields(logrus.Fields{"listen": port}).Info("Grpc")
	s.Serve(lis)
}
