package matchbox

import (
	"container/heap"
	"context"
	"log"
	"time"
)

const groupSize = 4

type Matchbox struct {
	Datacenters           map[DatacenterID]*Datacenter
	ActivePlayers         map[PlayerID]*Player
	InGamePlayers         map[PlayerID]*Player
	BetweenMatchesPlayers map[PlayerID]*Player

	queue MatchQueue
}

func NewMatchbox() *Matchbox {
	m := &Matchbox{
		ActivePlayers:         make(map[PlayerID]*Player),
		InGamePlayers:         make(map[PlayerID]*Player),
		BetweenMatchesPlayers: make(map[PlayerID]*Player),
		queue:                 make(MatchQueue, 0),
	}

	heap.Init(&m.queue)

	return m
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

				if m.queue.Len() > groupSize {
					players := make([]*Player, 0, groupSize)
					for i := 0; i < groupSize; i++ {
						entry := heap.Pop(&m.queue).(*QueueEntry)
						players = append(players, entry.Players...)
					}

					for _, player := range players {
						log.Printf("player %d\n", player.ID)
					}
				}
			}
		}
	}
}

func (m *Matchbox) AddPlayerToQueue(ctx context.Context, player *Player) error {
	m.ActivePlayers[player.ID] = player

	entry := &QueueEntry{
		Players:  []*Player{player},
		Priority: uint64(player.ID),
	}

	heap.Push(&m.queue, entry)

	return nil
}
