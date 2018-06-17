// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of FanIn2Any simple binary Fan-In

// FanIn2Any returns a channel to receive all to receive all from both `inp1` and `inp2` before close.
func FanIn2Any(inp1, inp2 <-chan Any) (out <-chan Any) {
	cha := make(chan Any)
	go fanIn2Any(cha, inp1, inp2)
	return cha
}

/* not used any more - kept for reference only.
// fanin2Any as seen in Go Concurrency Patterns
func fanin2Any(out chan<- Any, inp1, inp2 <-chan Any) {
	for {
		select {
		case e := <-inp1:
			out <- e
		case e := <-inp2:
			out <- e
		}
	}
} */

func fanIn2Any(out chan<- Any, inp1, inp2 <-chan Any) {
	defer close(out)

	var (
		closed bool // we found a chan closed
		ok     bool // did we read sucessfully?
		e      Any  // what we've read
	)

	for !closed {
		select {
		case e, ok = <-inp1:
			if ok {
				out <- e
			} else {
				inp1 = inp2   // swap inp2 into inp1
				closed = true // break out of the loop
			}
		case e, ok = <-inp2:
			if ok {
				out <- e
			} else {
				closed = true // break out of the loop				}
			}
		}
	}

	// inp1 might not be closed yet. Drain it.
	for e = range inp1 {
		out <- e
	}
}

// End of FanIn2Any simple binary Fan-In
// ===========================================================================
