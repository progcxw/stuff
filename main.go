package main

import (
	"gamestuff/backend/internal/consts"
	"gamestuff/backend/internal/pkg/config"
	"gamestuff/backend/internal/pkg/gorm"
	"gamestuff/backend/internal/router"
)

func main() {
	config.Default(consts.LogLevel)
	gorm.Init()
	consts.Init()

	router.ServeHTTP()
}
