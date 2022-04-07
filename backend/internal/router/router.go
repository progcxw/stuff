package router

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gamestuff/backend/internal/consts"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ServeHTTP() {
	engine := gin.Default()
	registerRouter(engine)

	srv := &http.Server{
		Addr:    consts.ListenAddr,
		Handler: engine,
	}
	startServer(srv)
	gracefulShutdown(srv)
}


func registerRouter(router *gin.Engine) {

}

func startServer(srv *http.Server) {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("http server listen err: %+v", err)
		}
	}()
}

func gracefulShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Infoln("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Infoln("Server exiting")
}