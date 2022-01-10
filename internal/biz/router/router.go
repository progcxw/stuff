// Package router git.nova.net.cn/nova/misc/cucc-sdwan/velocloud
//
//	联通SD-WAN适配层
//
//   Schemes: http
//   Consumes:
//   - application/json
//   Produces:
//   - application/json
//   Version: 0.1
//   swagger:meta
package router

import (
	"TimeLine/internal/consts"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() error {
	router := gin.Default()

	return router.Run(consts.ListenAddr)
}
