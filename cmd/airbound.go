package cmd

import (
	"airbound/api"
	"airbound/internal"
	"airbound/internal/config"
	"airbound/internal/log"
	"context"
	"errors"
	"fmt"
	"net/http"
)

// var glbClient *ent.Client

func Execute() {
	context := context.Background()
	cfg := config.LoadConfig()

	client := config.NewEntClient(cfg)
	defer client.Close()

	awsSESSession := config.NewSESSession(cfg)

	awsSQSSession := config.NewSQSSession(cfg)

	config.SchemaMigrateUp(context, client)

	router := api.Run(client, cfg, awsSESSession, awsSQSSession)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Error("server error: %s", err)
		}
	}()

	internal.GracefulShutdown(srv)
}
