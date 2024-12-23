package main

import (
	"context"

	"github.com/jegj/linktly/internal/api"
	"github.com/jegj/linktly/internal/config"
)

var Version = "development"

func main() {
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	config := config.NewConfig()

	server := api.NewServer(config, serverCtx, Version)
	server.Start(serverCtx, serverStopCtx)
}
