package internal

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const TimeoutWindow = 5 * time.Second

func GracefulShutdown(srv *http.Server) {
	close := make(chan os.Signal, 1)

	signal.Notify(close, syscall.SIGINT, syscall.SIGTERM)

	<-close

	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), TimeoutWindow)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("forcing server shutdown: %s", err)
	}

	log.Println("exiting server...")
}
