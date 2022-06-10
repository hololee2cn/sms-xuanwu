package controller

import (
	"context"
	"time"

	"github.com/hololee2cn/sms-xuanwu/internal/model"
	"github.com/hololee2cn/sms-xuanwu/internal/service"
	"github.com/hololee2cn/sms-xuanwu/pkg/grpcIFace"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var tLayout = "20060102150405"

type Callback struct {
	grpcIFace.UnimplementedCallbackServer
}

func (Callback) RecvCallback(ctx context.Context, req *grpcIFace.CallbackRequest) (*grpcIFace.CallbackResponse, error) {
	log.Info(ctx, req)
	if req.Total == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty items")
	}

	ms := make([]model.MailState, 0, req.Total)
	for _, v := range req.DataList {
		m := model.MailState{
			BatchID:       v.Uuid,
			To:            v.Phone,
			SendMessageID: v.CustomMsgID,
			RecvMessageID: v.MsgID,
			State:         int(v.State),
			OriginResult:  v.OriginResult,
		}

		t, err := time.ParseInLocation(tLayout, v.SubmitTime, time.Local)
		if err != nil {
			log.Error(err)
		} else {
			m.SubmitTime = t.Unix()
		}
		t, err = time.ParseInLocation(tLayout, v.DoneTime, time.Local)
		if err != nil {
			log.Error(err)
		} else {
			m.DoneTime = t.Unix()
		}

		ms = append(ms, m)
	}

	err := service.InsertMailsState(ms)
	if err != nil {
		log.Errorf("%v, %+v", err, ms)
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Info("callback finish")
	return &grpcIFace.CallbackResponse{}, nil
}
