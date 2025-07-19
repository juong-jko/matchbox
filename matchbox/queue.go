package matchbox

// An QueueEntry is something we manage in a priority queue.
type QueueEntry struct {
	Priority uint64 // The priority of the item in the queue.
	Players  []*Player
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int
}

// A MatchQueue implements heap.Interface and holds Items.
type MatchQueue []*QueueEntry

func (mq MatchQueue) Len() int { return len(mq) }

func (mq MatchQueue) Less(i, j int) bool {
	// Max heap
	return mq[i].Priority > mq[j].Priority
}

func (mq MatchQueue) Swap(i, j int) {
	mq[i], mq[j] = mq[j], mq[i]
	mq[i].index = i
	mq[j].index = j
}

func (mq *MatchQueue) Push(x any) {
	n := len(*mq)
	item := x.(*QueueEntry)
	item.index = n
	*mq = append(*mq, item)
}

func (mq *MatchQueue) Pop() any {
	old := *mq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*mq = old[0 : n-1]
	return item
}
