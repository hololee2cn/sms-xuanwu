package controller

import (
	"context"
	"time"

	"github.com/hololee2cn/sms-xuanwu/internal/model"
	"github.com/hololee2cn/sms-xuanwu/internal/service"
	"github.com/hololee2cn/sms-xuanwu/pkg/grpcIFace"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Sender struct {
	grpcIFace.UnimplementedSenderServer
}

/*
0	成功
-1	账号无效
-2	参数：无效
-3	连接不上服务器， 原因一般为:
	1.  HTTP请求方法不正确, 如POST请求误写为GET。
	2. 请求头中Content-Type、Accept、Authorization填写不正确。
-6	用户名密码错误
-7	旧密码不正确
-9	资金账户不存在
-11	包号码数量超过最大限制
-12	余额不足
-13	账号没有发送权限
-99	系统内部错误
-100/-200	其它错误/网络问题
*/
func fail(code string) codes.Code {
	if code == "0" {
		return codes.OK
	}
	switch code {
	case "-1", "-3", "-6", "-7", "-9", "-13":
		return codes.Internal
	case "-2", "-11":
		return codes.InvalidArgument
	case "-12":
		return codes.ResourceExhausted
	default:
		return codes.Unknown
	}
}

func (Sender) SendMessage(ctx context.Context, req *grpcIFace.SendMsgRequest) (resp *grpcIFace.SendMsgResponse, err error) {
	log.Info(ctx, req)
	var m model.SendMessage
	m.Content = req.Content
	m.MsgType = "sms"
	for _, v := range req.Items {
		item := model.MessageItem{
			To:          v.To,
			CustomMsgID: v.MessageID,
		}
		if len(item.CustomMsgID) == 0 {
			item.CustomMsgID = uuid.New().String()
		}
		m.Items = append(m.Items, item)
	}

	recv, err := service.SendSms(m)
	if err != nil {
		log.Error(err)
		return
	}
	if fail(recv.Code) != codes.OK {
		s := status.New(codes.Unknown, recv.Message)
		d, e := s.WithDetails(&grpcIFace.SendMsgError{
			Code:    recv.Code,
			Message: recv.Message,
		})
		if e != nil {
			return nil, s.Err()
		}
		return nil, d.Err()
	}
	resp = new(grpcIFace.SendMsgResponse)
	resp.BatchID = recv.UUID

	go func() {
		ls := make([]model.MailContent, 0, len(req.Items))
		ts := time.Now().Unix()
		for _, v := range m.Items {
			ls = append(ls, model.MailContent{
				BatchID:       recv.UUID,
				To:            v.To,
				Content:       m.Content,
				Time:          ts,
				SendMessageID: v.CustomMsgID,
				Sender:        req.Sender,
			})
		}
		err = service.InsertNewLogs(ls)
		if err != nil {
			log.Error(err)
			return
		}
	}()
	log.Info(resp, err)
	return
}
