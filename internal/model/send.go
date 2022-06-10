package model

type MessageItem struct {
	To          string `json:"to"`
	CustomMsgID string `json:"customMsgID"`
}
type SendMessage struct {
	Items   []MessageItem `json:"items"`
	Content string        `json:"content"`
	MsgType string        `json:"msgType"`
}
type SendMessageRecv struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
	UUID    string `json:"uuid"`
}
