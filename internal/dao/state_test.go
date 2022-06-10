package dao

import (
	"testing"

	goxorm "github.com/go-xorm/xorm"
	"github.com/hololee2cn/sms-xuanwu/internal/pkg/xorm"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hololee2cn/sms-xuanwu/internal/model"
)

// a successful case
func TestInsertMailsState(t *testing.T) {

	e, err := goxorm.NewEngine("mysql", "root:test@tcp(notify-mysql:3306)/sms_xuanwu?charset=utf8")
	if err != nil {
		t.Fatal(err)
	}
	defer e.Close()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	e.DB().DB = db
	xorm.Engine = e

	mock.ExpectExec("INSERT INTO `mail_state` \\(`to`, `batch_id`, `send_msg_id`, `recv_msg_id`, `state`, `submit_time`, `done_time`, `origin_result`\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)").
		WithArgs("to",
			"batch_id",
			"send_msg_id",
			"recv_msg_id",
			1,
			1,
			2,
			"res").WillReturnResult(sqlmock.NewResult(1, 1))

	ms := make([]model.MailState, 0)
	ms = append(ms, model.MailState{
		To:            "to",
		BatchID:       "batch_id",
		SendMessageID: "send_msg_id",
		RecvMessageID: "recv_msg_id",
		State:         1,
		SubmitTime:    1,
		DoneTime:      2,
		OriginResult:  "res",
	})
	if err = InsertMailsState(ms); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
