package dao

import (
	"github.com/hololee2cn/sms-xuanwu/internal/model"

	"github.com/hololee2cn/sms-xuanwu/internal/pkg/xorm"
	log "github.com/sirupsen/logrus"
)

func InsertNewMailLogs(ls []model.MailContent) (err error) {
	_, err = xorm.Engine.Insert(ls)
	if err != nil {
		log.Error(err)
	}
	return
}
func FindLogs(query string, args ...interface{}) (logs []model.MailContent, err error) {
	// "batch_id = ? or send_msg_id = ?", id, id
	logs = make([]model.MailContent, 0)
	err = xorm.Engine.Where(query, args...).Find(&logs)
	if err != nil {
		log.Error(err)
	}
	return
}
