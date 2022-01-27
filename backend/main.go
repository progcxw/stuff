package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"TimeLine/backend/internal/biz"
	"TimeLine/backend/internal/consts"
	"TimeLine/backend/internal/pkg/gorm"

	log "github.com/sirupsen/logrus"
)

func main() {
	gorm.Init()
	consts.Init()

	ctx, cancel := context.WithCancel(context.Background())
	biz.Start(ctx)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Info("rev TERM signal!")
	biz.Close(cancel)
	log.Info("process exit!")
}
