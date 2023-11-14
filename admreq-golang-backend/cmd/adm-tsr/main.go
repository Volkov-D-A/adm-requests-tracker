package main

import (
	"log"
	"os/signal"
	"syscall"

	"github.com/volkov-d-a/adm-requests-tracker/internal/app"
	"golang.org/x/net/context"
)

func main() {
	se := make(chan error)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	a, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	a.ServiceProvider.Logger.Info("Application created")

	a.Run(se)

	select {
	case <-ctx.Done():
		a.ServiceProvider.Logger.Info("Shutting down application by signal")
	case err = <-se:
		a.ServiceProvider.Logger.Info("shutting down application by error", "err", err.Error())
	}

	a.ServiceProvider.Logger.Info("Starting graceful shutdown")
	err = a.GracefulStop()
	if err != nil {
		a.ServiceProvider.Logger.Error("Error graceful shutdown", "err", err.Error())
	}

}
