package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/isauran/gokit-microservice-profile/internal/app"
	"github.com/isauran/gokit-microservice-profile/pkg/logger"
)

func main() {
	ctx := context.Background()
	log := logger.SlogJSONLogger(os.Stderr, slog.LevelError)

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Error("failed to init app", err)
		os.Exit(1)
	}

	err = a.Run()
	if err != nil {
		log.Error("failed to run app", err)
		os.Exit(1)
	}
}
