package main

import (
	"log"
	"os/signal"
	"syscall"

	"github.com/volkov-d-a/adm-requests-tracker/internal/app"
	"golang.org/x/exp/slog"
	"golang.org/x/net/context"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	a.ServiceProvider.Logger.Info("Application created")

	err = a.Run()
	if err != nil {
		a.ServiceProvider.Logger.Error("Error running application", slog.String("err", err.Error()))
	}

	<-ctx.Done()
	a.ServiceProvider.Logger.Info("Starting graceful shutdown")
	err = a.GracefulStop()
	if err != nil {
		a.ServiceProvider.Logger.Error("Error graceful shutdown", slog.String("err", err.Error()))
	}

}
