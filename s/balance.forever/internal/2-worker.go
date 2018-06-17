// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package balance

// ===========================================================================
// Beg of Worker

// A Worker works on received requests
type Worker struct {
	requests chan Request // work to do (a buffered channel)
	pending  int          // count of pending tasks
	index    int          // index in the heap
}

// work keeps receiving requests, and for each does:
//  - reply on the requestor-provided channel the result into the request
//  - inform on the balancer-provided channel by sending itself when done
func (w *Worker) work(done chan<- *Worker) {
	for {
		req := <-w.requests // get requests from load balancer
		req.c <- req.fn()   // do the work and send the answer back to the requestor
		done <- w           // tell load balancer a task has been completed by worker w.
	}
}

// End of Worker
// ===========================================================================
