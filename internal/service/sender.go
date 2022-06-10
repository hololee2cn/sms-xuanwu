package service

import (
	"encoding/json"
	"net/http"

	"github.com/hololee2cn/sms-xuanwu/internal/consts"
	"github.com/hololee2cn/sms-xuanwu/internal/model"

	"github.com/hololee2cn/sms-xuanwu/internal/pkg/httpUtil"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SendSms(m model.SendMessage) (recv model.SendMessageRecv, err error) {
	b, err := json.Marshal(m)
	if err != nil {
		log.Error(err)
		err = status.Error(codes.InvalidArgument, err.Error())
		return
	}
	h := map[string][]string{
		"Content-Type":  {"application/json;charset=utf-8"},
		"Accept":        {"application/json"},
		"Authorization": {consts.SendAuthToken},
	}
	code, body, err := httpUtil.Request("POST", consts.SendUrl, b, h)
	if err != nil {
		log.Error(err)
		err = status.Error(codes.Internal, err.Error())
		return
	}
	log.Debug(string(body))
	if code != http.StatusOK {
		err = status.Error(codes.Unknown, string(body))
		log.Error(err)
		return
	}
	err = json.Unmarshal(body, &recv)
	if err != nil {
		log.Error(err)
		err = status.Error(codes.Unknown, string(body))
	}
	return
}
