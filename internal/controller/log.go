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

type Log struct {
	grpcIFace.UnimplementedLogServer
}

func (Log) ListLogsByID(ctx context.Context, req *grpcIFace.ListByIDRequest) (resp *grpcIFace.ListContentsResponse, err error) {
	log.Info(ctx, req)
	mails, e := service.ListLogsByID(req.ID)
	if e != nil {
		err = status.Error(codes.Internal, e.Error())
		return
	}
	resp = &grpcIFace.ListContentsResponse{
		Total: int32(len(mails)),
	}
	for _, v := range mails {
		resp.Mails = append(resp.Mails, &grpcIFace.SmsContent{
			To:            v.To,
			BatchID:       v.BatchID,
			Content:       v.Content,
			Sender:        v.Sender,
			SendMessageID: v.SendMessageID,
			Time:          v.Time,
		})
	}
	log.Info(resp, err)
	return
}
func (Log) ListLogsByTime(ctx context.Context, req *grpcIFace.ListByTimeRequest) (resp *grpcIFace.ListContentsResponse, err error) {
	log.Info(ctx, req)
	if req.TimeEnd <= 0 {
		req.TimeEnd = time.Now().Unix()
	}
	if req.TimeBegin <= 0 {
		req.TimeBegin = req.TimeEnd - 7200
	}

	mails, e := service.ListLogsByTime(req.TimeBegin, req.TimeEnd)
	if e != nil {
		err = status.Error(codes.Internal, e.Error())
		return
	}
	resp = &grpcIFace.ListContentsResponse{
		Total: int32(len(mails)),
	}
	for _, v := range mails {
		resp.Mails = append(resp.Mails, &grpcIFace.SmsContent{
			To:            v.To,
			BatchID:       v.BatchID,
			Content:       v.Content,
			Sender:        v.Sender,
			SendMessageID: v.SendMessageID,
			Time:          v.Time,
		})
	}
	log.Info(resp, err)
	return
}
