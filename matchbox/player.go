package matchbox

type PlayerState int

const (
	New PlayerState = iota
	Ideal
	Expand
	WarmBody
)

type PlayerID uint64

// Represents a player in the matchmaking system
type Player struct {
	ID          PlayerID
	PlayerState PlayerState
	Counter     int
	Latency     float64
}

type Party struct {
	Players []*Player
}
