// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package balance

import (
	"container/heap"
	"fmt"
)

// ===========================================================================
// Beg of Balancer

// Balancer has a Pool of Workers and a channel for Workers having finished
type Balancer struct {
	pool Pool
	done chan *Worker
}

// New returns a receive-only request channel
// processed by `cap` balanced workers
func New(cap int) chan<- Request {
	b := &Balancer{
		pool: make([]*Worker, 0, cap),
		done: make(chan *Worker),
	}

	for i := 0; i < cap; i++ { // populate the worker pool
		work := make(chan Request) // work to receive
		w := Worker{work, 0, i}    // by worker with index `i`
		b.pool[i] = &w             // as pool[i]
		go w.work(b.done)          // launch worker to work
	}

	heap.Init(&b.pool)

	work := make(chan Request)
	go b.balance(work)

	return work
}

// balance the work
func (b *Balancer) balance(work <-chan Request) {
	for {
		select {
		case req := <-work: // received a Request...
			b.dispatch(req) // ...so send it to a Worker
		case w := <-b.done: // a worker has finished ...
			b.completed(w) // ...so update its info
		}
		b.print()
	}
}

// dispatch sends Request to Worker
func (b *Balancer) dispatch(req Request) {
	w := heap.Pop(&b.pool).(*Worker) // grab least loaded worker ...
	w.requests <- req                // ... assign it the task.
	w.pending++                      // One more in it's queue.
	heap.Push(&b.pool, w)            // Push it back into its place on the heap.
}

// completed is the Job: update the Heap
func (b *Balancer) completed(w *Worker) {
	w.pending--                // one fewer in it's queue
	heap.Fix(&b.pool, w.index) // calling Fix is equivalent to, but less expensive than, calling Remove(h, i) followed by a Push of the new value.
	// heap.Remove(&b.pool, w.index) // remove it from Heap
	// heap.Push(&b.pool, w)         // put it back where it belongs
}

func (b *Balancer) print() {
	totalPending := 0
	sumsqPending := 0
	for _, w := range b.pool { // worker
		fmt.Printf("%d  ", w.pending)
		totalPending += w.pending
		sumsqPending += w.pending * w.pending
	}
	fmt.Printf("| %d  ", totalPending)
	avg := float64(totalPending) / float64(b.pool.Len())
	variance := float64(sumsqPending)/float64(len(b.pool)) - avg*avg
	fmt.Printf("| %.2f %.2f\n", avg, variance)

}

// End of Balancer
// ===========================================================================
