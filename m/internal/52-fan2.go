// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingFanIn2 simple binary Fan-In

// anyThingFanIn2 returns a channel to receive
// all from both `inp` and `inp2`
// before close.
func (inp anyThingFrom) anyThingFanIn2(inp2 anyThingFrom) (out anyThingFrom) {
	cha := make(chan anyThing)
	go inp.fanIn2anyThing(cha, inp2)
	return cha
}

/* not used - kept for reference only.
// (inp anyThingFrom) fanin2anyThing as seen in Go Concurrency Patterns
func fanin2anyThing(out anyThingInto, inp, inp2 anyThingFrom) {
	for {
		select {
		case e := <-inp:
			out <- e
		case e := <-inp2:
			out <- e
		}
	}
} */

func (inp anyThingFrom) fanIn2anyThing(out anyThingInto, inp2 anyThingFrom) {
	defer close(out)

	var (
		closed bool     // we found a chan closed
		ok     bool     // did we read successfully?
		e      anyThing // what we've read
	)

	for !closed {
		select {
		case e, ok = <-inp:
			if ok {
				out <- e
			} else {
				inp = inp2    // swap inp2 into inp
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

	// inp might not be closed yet. Drain it.
	for e = range inp {
		out <- e
	}
}

// End of anyThingFanIn2 simple binary Fan-In
// ===========================================================================
