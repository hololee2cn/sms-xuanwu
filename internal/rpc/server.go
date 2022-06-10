package rpc

import (
	"net"

	"github.com/hololee2cn/sms-xuanwu/internal/consts"
	"github.com/hololee2cn/sms-xuanwu/internal/controller"
	"github.com/hololee2cn/sms-xuanwu/pkg/grpcIFace"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Start() {
	l, err := net.Listen("tcp", consts.GrpcListenAddr)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	grpcIFace.RegisterSenderServer(s, &controller.Sender{})
	grpcIFace.RegisterCallbackServer(s, &controller.Callback{})
	grpcIFace.RegisterLogServer(s, &controller.Log{})
	grpcIFace.RegisterStateServer(s, &controller.State{})
	defer s.GracefulStop()

	log.Info(s.Serve(l))
}
