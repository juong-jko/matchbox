package main

import (
	"context"
	"juong/matchbox/matchbox"
	"time"
)

const (
	SimulationDuration = 8 * time.Second
)

func main() {
	m := matchbox.NewMatchbox()

	for i := 0; i < 10; i++ {
		player := &matchbox.Player{
			ID: matchbox.PlayerID(i),
		}
		m.AddPlayerToQueue(context.Background(), player)
	}

	m.RunSimulation(context.Background(), SimulationDuration)
}
