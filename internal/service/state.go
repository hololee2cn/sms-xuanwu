package service

import (
	"fmt"

	"github.com/hololee2cn/sms-xuanwu/internal/dao"
	"github.com/hololee2cn/sms-xuanwu/internal/model"

	log "github.com/sirupsen/logrus"
)

func InsertMailsState(ms []model.MailState) (err error) {
	return dao.InsertMailsState(ms)
}
func ListStateByTime(tb, te int64) (ms []model.MailState, err error) {
	log.Info(tb, te)
	return dao.FindMailsState("submit_time >= ? and submit_time <= ?", tb, te)
}
func ListStateByID(msgID string) (ms []model.MailState, err error) {
	return dao.FindMailsState("batch_id = ? or send_msg_id like ? or recv_msg_id = ?", msgID, fmt.Sprintf("%v%%", msgID), msgID)
}
