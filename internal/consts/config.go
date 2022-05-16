package consts

import (
	"gamestuff/backend/internal/pkg/config"

	log "github.com/sirupsen/logrus"
)

var (
	LogLevel   = log.Level(config.DefaultInt("log_level", 4))
	ListenAddr = config.DefaultString("listen_addr", ":80")
)

func Init() {

}
