package main

import (
	"context"
	"os"

	"github.com/isauran/gokit-microservice-profile/internal/app"
	"github.com/isauran/logger"
)

func main() {
	ctx := context.Background()

	log := logger.NewLogger(os.Stdout, logger.WithJSON(true))

	a, err := app.NewApp(ctx, log)
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
