// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package balance

// ===========================================================================
// Beg of Request

// Request is a function to be applied and channel on which to return the result.
type Request struct {
	fn func() Any // operation to perform
	c  chan Any   // channel on which to return result
}

// Beg of Fake
func workFn() (a Any) { return }

func requester(work chan<- Request) {
	cha := make(chan Any)
	for {
		// time.Sleep ....
		work <- Request{workFn, cha} // send a work request
		result := <-cha              // wait for answer
		_ = result                   // furtherProcess(result)
	}
}

func process() {
	requester(New(10))
}

// End of Fake

// End of Request
// ===========================================================================
