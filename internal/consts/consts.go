package consts

import config "github.com/hololee2cn/sms-xuanwu/internal/pkg/config2"

var mosBaseUrl = config.DefaultString("mos_base_url", "https://mos.wxchina.com:9100")
var sendSmsPath = config.DefaultString("mox_send_path", "/api/v1.0.0/message/mass/send")
var SendUrl = mosBaseUrl + sendSmsPath
var SendAuthToken = config.DefaultString("mox_send_auth", "ZGV2b3BzQG5sOjk3NmZkMTg5MWU4YTBkNDNhZGZiNmVmMTg5MTZlOTMy")
var GrpcListenAddr = config.DefaultString("grpc_listen_addr", ":80")
var HttpListenAddr = config.DefaultString("http_listen_addr", ":8080")
