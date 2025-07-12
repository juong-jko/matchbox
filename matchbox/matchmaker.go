package matchbox

import (
	"context"
	"log"
	"time"
)

type Matchbox struct {
	Datacenters           map[DatacenterID]*Datacenter
	ActivePlayers         map[PlayerID]*Player
	InGamePlayers         map[PlayerID]*Player
	BetweenMatchesPlayers map[PlayerID]*Player

	MatchQueue
}

func NewMatchbox() *Matchbox {
	return &Matchbox{}
}

func (m *Matchbox) RunSimulation(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ctx.Done():
			{
				log.Println("Termination signal received")
				return
			}
		case <-ticker.C:
			{
				log.Println("Running matchmaker tick")
			}
		}
	}
}
