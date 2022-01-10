package biz

import (
	"context"

	"TimeLine/internal/biz/router"
	"TimeLine/internal/config"
	"TimeLine/internal/consts"
)

func Start(ctx context.Context) {
	config.Default(consts.LogLevel)
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	panic(router.ServeHTTP())
}

func Close(cancel context.CancelFunc) {
	close(consts.Ch)
	cancel()
}
