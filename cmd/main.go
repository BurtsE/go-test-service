package main

import (
	"context"
	"log"
	"test-service/internal/app"
)

func main() {
	ctx := context.Background()

	a := app.NewApp(ctx)
	if a == nil {
		log.Fatalf("failed to create app")
	}
	err := a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
