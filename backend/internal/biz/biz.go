package biz

import (
	"context"

	"TimeLine/backend/internal/biz/router"
	"TimeLine/backend/internal/config"
	"TimeLine/backend/internal/consts"
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
