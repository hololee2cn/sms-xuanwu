package dao

import (
	"github.com/hololee2cn/sms-xuanwu/internal/model"

	"github.com/hololee2cn/sms-xuanwu/internal/pkg/xorm"
	log "github.com/sirupsen/logrus"
)

func FindMailsState(query string, args ...interface{}) (res []model.MailState, err error) {
	res = make([]model.MailState, 0)
	err = xorm.Engine.Where(query, args...).Find(&res)
	if err != nil {
		log.Error(err)
	}
	return
}
func InsertMailsState(ms []model.MailState) (err error) {
	xorm.Engine.ShowSQL(true)
	_, err = xorm.Engine.Omit("id").Insert(ms)
	if err != nil {
		log.Error(err)
	}
	return
}
func UpdateMailState(record model.MailState) (err error) {
	_, err = xorm.Engine.Where("send_msg_id = ?", record.SendMessageID).Update(record)
	if err != nil {
		log.Error(err)
	}
	return
}
