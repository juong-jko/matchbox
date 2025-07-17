package matchbox

type QueueEntry struct {
	Priority uint64
	Players  []*Player
	Index    int
}

type MatchQueue []*QueueEntry

func (queue *MatchQueue) Len() int {
	return len(*queue)
}

func (queue *MatchQueue) Less(i, j int) bool {
	return (*queue)[i].Priority < (*queue)[j].Priority
}

func (queue *MatchQueue) Push(x any) {
	n := len(*queue)
	item := x.(*QueueEntry)
	item.Index = n
	*queue = append(*queue, item)
}

func (queue *MatchQueue) Pop() any {
	n := len(*queue)
	entry := (*queue)[n-1]
	entry.Index = -1

	(*queue)[n-1] = nil
	*queue = (*queue)[0 : n-1]
	return entry
}
