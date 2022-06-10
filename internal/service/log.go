package service

import (
	"github.com/hololee2cn/sms-xuanwu/internal/dao"
	"github.com/hololee2cn/sms-xuanwu/internal/model"
)

func ListLogsByID(id string) (logs []model.MailContent, err error) {
	return dao.FindLogs("batch_id = ? or send_msg_id = ?", id, id)
}
func ListLogsByTime(tb, te int64) (logs []model.MailContent, err error) {
	return dao.FindLogs("time >= ? and time <= ?", tb, te)
}
func InsertNewLogs(logs []model.MailContent) (err error) {
	return dao.InsertNewMailLogs(logs)
}
