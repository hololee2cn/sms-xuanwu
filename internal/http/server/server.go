package server

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hololee2cn/sms-xuanwu/internal/consts"
	"github.com/hololee2cn/sms-xuanwu/pkg/grpcIFace"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Start() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// grpc服务地址
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// HTTP转grpc
	err := grpcIFace.RegisterCallbackHandlerFromEndpoint(ctx, mux, consts.GrpcListenAddr, opts)
	if err != nil {
		panic(err)
	}
	log.Info(http.ListenAndServe(consts.HttpListenAddr, mux))
}
