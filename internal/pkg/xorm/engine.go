package xorm

// 涉及外部依赖获取
// 1. go get github.com/go-sql-driver/mysql
// 2. go get github.com/go-xorm/xorm

import (
	"time"

	c "github.com/hololee2cn/sms-xuanwu/internal/pkg/config2"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
)

type Session = xorm.Session

var (
	Engine *xorm.Engine

	// mysql 相关配置
	dataSource  string
	maxIdleConn = c.DefaultInt("max_idle_conn", 1000)
	maxOpenConn = c.DefaultInt("max_open_conn", 1000)
)

func Init() {
	var err error

	dataSource = c.MustString("data_source")
	Engine, err = xorm.NewEngine("mysql", dataSource)
	if err != nil {
		log.Panic("Failed to xorm.NewEngine, err:", err)
	}
	err = Engine.Ping()
	if err != nil {
		log.Panic("Failed to connect MYSQL, err:", err)
	}

	Engine.SetMaxIdleConns(maxIdleConn)
	Engine.SetMaxOpenConns(maxOpenConn)
	Engine.SetConnMaxLifetime(time.Second * 10)
	log.Info("Success to connect mysql...")
}
