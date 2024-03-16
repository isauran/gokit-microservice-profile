package main

import (
	"context"
	"github.com/isauran/gokit-microservice-profile/internal/app"
	"log"
)

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app %s", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app %s", err)
	}
}
