// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// Any is the generic type flowing thru the pipe network.
type Any generic.Type

// ===========================================================================
// Beg of PlugAny - graceful terminator

// PlugAny returns a channel to receive every `inp` before close and a channel to signal this closing.
// Upon receipt of a stop signal,
// output is immediately closed,
// and for graceful termination
// any remaining input is drained before done is signalled.
func PlugAny(inp <-chan Any, stop <-chan struct{}) (out <-chan Any, done <-chan struct{}) {
	cha := make(chan Any)
	doit := make(chan struct{})
	go plugAny(cha, doit, inp, stop)
	return cha, doit
}

func plugAny(out chan<- Any, done chan<- struct{}, inp <-chan Any, stop <-chan struct{}) {
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
		case <-stop:
			end = true
		}
	}

	close(out)

	for range inp {
		// drain inp
	}

	done <- struct{}{}
}

// End of PlugAny - graceful terminator
// ===========================================================================
