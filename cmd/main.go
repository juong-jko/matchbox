package main

import (
	"context"
	"juong/matchbox/matchbox"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx := context.Background()
	signalCtx, cancel := signal.NotifyContext(ctx, os.Kill, os.Interrupt)
	defer cancel()

	box := matchbox.NewMatchbox()
	box.RunSimulation(signalCtx, time.Duration(10))
}
