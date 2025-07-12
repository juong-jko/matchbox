package matchbox

type QueueEntry struct {
	Priority uint64
	Players  []*Player
}

type MatchQueue []*QueueEntry
