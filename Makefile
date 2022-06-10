bin="mybin"

.PHONY: build
build: proto format
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags='-s -w' -o $(bin)

.PHONY: format
format:
	gofmt -w .

.PHONY: proto
proto:
	protoc --go_out=. --go-grpc_out=.  --grpc-gateway_out=logtostderr=true:. --openapiv2_out=. --openapiv2_opt=logtostderr=true api/*.proto
