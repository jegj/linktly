package main

import (
	"context"

	bootstrap "github.com/jegj/linktly/internal/api/booststrap"
)

func main() {
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	server := bootstrap.NewServer(serverCtx)
	server.Start(serverCtx, serverStopCtx)
}
