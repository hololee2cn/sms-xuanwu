module github.com/hololee2cn/sms-xuanwu

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.4.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/xorm v0.7.9
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.3.0
	github.com/hololee2cn/pkg v0.0.0-20220609064745-6bcbb6ef3e4b
	github.com/hololee2cn/sms-xuanwu/pkg/grpcIFace v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.36.1
	xorm.io/builder v0.3.7 // indirect
)

replace github.com/hololee2cn/sms-xuanwu/pkg/grpcIFace => ./pkg/grpcIFace
