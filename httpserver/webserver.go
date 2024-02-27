package httpserver

import (
	"golang.org/x/sync/errgroup"
	"monopoly-server/settings"
	"net/http"
	"time"
	"context"
)

var server *http.Server

func InitializeWebServer() {
	server = &http.Server{
		Addr:   settings.GetWebServerAddress(),
		Handler: routerEngine(),
	}
}

func RunWebServer() error {
	var g errgroup.Group
	g.Go(func() error {
		timeout := settings.ShutdownTimeout() * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		return server.Shutdown(ctx)
	})

	g.Go(func() error {
		if err := server.ListenAndServe();err != nil&& err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	return g.Wait()
}