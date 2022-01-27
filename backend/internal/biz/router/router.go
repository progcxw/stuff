package router

import (
	"TimeLine/backend/internal/consts"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() error {
	router := gin.Default()

	return router.Run(consts.ListenAddr)
}
