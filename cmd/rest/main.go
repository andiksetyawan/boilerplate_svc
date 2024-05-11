package main

import (
	"context"

	"github.com/andiksetyawan/boilerplate_svc/internal/app"
)

func main() {
	if restServer, err := app.NewRestServer(context.Background()); err == nil {
		restServer.Run()
	}
}
