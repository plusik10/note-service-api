package main

import (
	"context"
	"log"
	"os"

	"github.com/plusik10/note-service-api/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx, os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatalf("Failed to create app: %v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("Failed to run app: %v", err)
	}

}
