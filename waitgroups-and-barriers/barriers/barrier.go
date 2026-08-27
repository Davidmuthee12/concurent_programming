package barriers

import "sync"

type Barrier struct {
	size      int // Total number of participants in the barrier
	waitCount int // Counter variable representing the number of concurrently suspended executions
	cond      *sync.Cond // Condition variable used in the barrier
}

func NewBarrier(size int) *Barrier {
	// Creates new condition variable
	condvar := sync.NewCond(&sync.Mutex{})
	// Creates and returns reference to new barrier
	return &Barrier{size,0,condvar}
}

func (b *Barrier) Wait() {
	// Protects access to the waitCount variable by using a mutex
	b.cond.L.Lock()
	// Increments the count variable by 1
	b.waitCount += 1

	// if waitCount has reached the barrier size, resets waitCount and broadcasts on the condition variable
	if b.waitCount == b.size {
		b.waitCount = 0
		b.cond.Broadcast()
	} else {
		// if waitCount hasnt reached the barrier size, waits on condition variable
		b.cond.Wait()
	}
	// Protects access to the waitCount variable by using a mutex
	b.cond.L.Unlock()
}