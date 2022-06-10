package main

import (
	http "github.com/hololee2cn/sms-xuanwu/internal/http/server"
	rpc2 "github.com/hololee2cn/sms-xuanwu/internal/rpc"

	"github.com/hololee2cn/pkg/extra"
	"github.com/hololee2cn/sms-xuanwu/internal/pkg/xorm"
)

func main() {
	extra.Default()
	xorm.Init()
	go rpc2.Start()
	http.Start()
}
