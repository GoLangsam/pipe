// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"time"

	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// anyOwner is the generic who shall own the methods.
//  Note: Need to use `generic.Number` here as `generic.Type` is an interface and cannot have any method.
type anyOwner generic.Number

// ===========================================================================
// Beg of anyThingPlugAfter - graceful terminator

// anyThingPlugAfter returns a channel to receive every `inp` before close and a channel to signal this closing.
// Upon receipt of a time signal
// (e.g. from `time.After(...)`),
// output is immediately closed,
// and for graceful termination
// any remaining input is drained before done is signalled.
func (my anyOwner) anyThingPlugAfter(inp <-chan anyThing, after <-chan time.Time) (out <-chan anyThing, done <-chan struct{}) {
	cha := make(chan anyThing)
	doit := make(chan struct{})
	go my.pluganyThingAfter(cha, doit, inp, after)
	return cha, doit
}

func (my anyOwner) pluganyThingAfter(out chan<- anyThing, done chan<- struct{}, inp <-chan anyThing, after <-chan time.Time) {
	defer close(done)

	var end bool   // shall we end?
	var ok bool    // did we read successfully?
	var e anyThing // what we've read

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

// End of anyThingPlugAfter - graceful terminator
// ===========================================================================
