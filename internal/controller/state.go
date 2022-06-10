package controller

import (
	"context"
	"time"

	"github.com/hololee2cn/sms-xuanwu/internal/service"
	"github.com/hololee2cn/sms-xuanwu/pkg/grpcIFace"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type State struct {
	grpcIFace.UnimplementedStateServer
}

func (State) ListStateByID(ctx context.Context, req *grpcIFace.ListByIDRequest) (resp *grpcIFace.ListStateResponse, err error) {
	log.Info(ctx, req)
	ms, e := service.ListStateByID(req.ID)
	if e != nil {
		err = status.Error(codes.Internal, e.Error())
		return
	}
	resp = &grpcIFace.ListStateResponse{
		Total: int32(len(ms)),
	}
	for _, v := range ms {
		resp.State = append(resp.State, &grpcIFace.SmsState{
			To:            v.To,
			BatchID:       v.BatchID,
			SendMessageID: v.SendMessageID,
			RecvMessageID: v.RecvMessageID,
			State:         grpcIFace.SmsState_State(v.State),
			SubmitTime:    v.SubmitTime,
			DownTime:      v.DoneTime,
			OriginResult:  v.OriginResult,
		})
	}
	log.Info(resp, err)
	return
}
func (State) ListStateByTime(ctx context.Context, req *grpcIFace.ListByTimeRequest) (resp *grpcIFace.ListStateResponse, err error) {
	log.Info(ctx, req)
	if req.TimeEnd <= 0 {
		req.TimeEnd = time.Now().Unix()
	}
	if req.TimeBegin <= 0 {
		req.TimeBegin = req.TimeEnd - 7200
	}

	ms, e := service.ListStateByTime(req.TimeBegin, req.TimeEnd)
	if e != nil {
		err = status.Error(codes.Internal, e.Error())
		return
	}
	resp = &grpcIFace.ListStateResponse{
		Total: int32(len(ms)),
	}
	for _, v := range ms {
		resp.State = append(resp.State, &grpcIFace.SmsState{
			To:            v.To,
			BatchID:       v.BatchID,
			SendMessageID: v.SendMessageID,
			RecvMessageID: v.RecvMessageID,
			State:         grpcIFace.SmsState_State(v.State),
			SubmitTime:    v.SubmitTime,
			DownTime:      v.DoneTime,
			OriginResult:  v.OriginResult,
		})
	}
	log.Info(resp, err)
	return
}
