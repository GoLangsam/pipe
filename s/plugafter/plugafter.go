// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"time"

	"github.com/cheekybits/genny/generic"
)

type Any generic.Type

// ===========================================================================
// Beg of PlugAnyAfter - graceful terminator

// PlugAnyAfter returns a channel to receive every `inp` before close and a channel to signal this closing.
// Upon receipt of a time signal
// (e.g. from `time.After(...)`),
// output is immediately closed,
// and for graceful termination
// any remaining input is drained before done is signalled.
func PlugAnyAfter(inp <-chan Any, after <-chan time.Time) (out <-chan Any, done <-chan struct{}) {
	cha := make(chan Any)
	doit := make(chan struct{})
	go plugAnyAfter(cha, doit, inp, after)
	return cha, doit
}

func plugAnyAfter(out chan<- Any, done chan<- struct{}, inp <-chan Any, after <-chan time.Time) {
	defer close(done)

	var end bool // shall we end?
	var ok bool  // did we read successfully?
	var e Any    // what we've read

	for !end {
		select {
		case e, ok = <-inp:
			if ok {
				out <- e
			} else {
				end = true
			}
		case <-after:
			end = true
		}
	}

	close(out)

	for range inp {
		// drain inp
	}

	done <- struct{}{}
}

// End of PlugAnyAfter - graceful terminator
// ===========================================================================
