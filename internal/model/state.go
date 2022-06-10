package model

type MailState struct {
	To            string `json:"to" xorm:"to"`
	BatchID       string `json:"batch_id" xorm:"batch_id"`
	SendMessageID string `json:"send_msg_id" xorm:"send_msg_id"`
	RecvMessageID string `json:"recv_msg_id" xorm:"recv_msg_id"`
	State         int    `json:"state" xorm:"state"`
	SubmitTime    int64  `json:"submit_time" xorm:"submit_time"`
	DoneTime      int64  `json:"done_time" xorm:"done_time"`
	OriginResult  string `json:"origin_result" xorm:"origin_result"`
}
