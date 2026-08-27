package waitgroupSemaphore

import (
	"github.com/Davidmuthee12/threads/semaphore"
)

type WaitGrp struct {
	// store semaphore reference
	sema *semaphore.Semaphore
}

func NewWaitGrp(size int) *WaitGrp{
	// initialize a new semaphore with 1 - size permits
	return &WaitGrp{sema: semaphore.NewSemaphore(1 - size)}
}

func (wg *WaitGrp) Wait() {
	// calls Acquire() on the semphore in the wait() function
	wg.sema.Acquire()
}

func (wg *WaitGrp) Done() {
	// when done, calls Release() on the semaphore
	wg.sema.Release()
}