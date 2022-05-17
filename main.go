package main

import (
	"stuff/internal/consts"
	"stuff/internal/router"
	"stuff/pkg/config"
	"stuff/pkg/gorm"
)

func main() {
	config.Default(consts.LogLevel)
	gorm.Init()
	consts.Init()

	router.ServeHTTP()
}
