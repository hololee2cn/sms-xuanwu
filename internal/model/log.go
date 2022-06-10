package model

type MailContent struct {
	Sender        string `json:"sender" xorm:"sender"`
	To            string `json:"to" xorm:"to"`
	BatchID       string `json:"batch_id" xorm:"batch_id"`
	SendMessageID string `json:"send_msg_id" xorm:"send_msg_id"`
	Content       string `json:"content" xorm:"content"`
	Time          int64  `json:"time" xorm:"time"`
}
