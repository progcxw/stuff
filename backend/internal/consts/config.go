package consts

import (
	"TimeLine/backend/internal/config"

	log "github.com/sirupsen/logrus"
)

var (
	LogLevel   = log.Level(config.DefaultInt("log_level", 4))
	ListenAddr = config.DefaultString("listen_addr", ":5555")

	// Ch 控制协程调用频率
	Ch chan struct{}
)

func Init() {

}
