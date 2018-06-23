// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// anyOwner is the generic who shall own the methods.
//  Note: Need to use `generic.Number` here as `generic.Type` is an interface and cannot have any method.
type anyOwner generic.Number

// ===========================================================================
// Beg of anyThingPlug - graceful terminator

// anyThingPlug returns a channel to receive every `inp` before close and a channel to signal this closing.
// Upon receipt of a stop signal,
// output is immediately closed,
// and for graceful termination
// any remaining input is drained before done is signalled.
func (my anyOwner) anyThingPlug(inp <-chan anyThing, stop <-chan struct{}) (out <-chan anyThing, done <-chan struct{}) {
	cha := make(chan anyThing)
	doit := make(chan struct{})
	go my.pluganyThing(cha, doit, inp, stop)
	return cha, doit
}

func (my anyOwner) pluganyThing(out chan<- anyThing, done chan<- struct{}, inp <-chan anyThing, stop <-chan struct{}) {
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

// End of anyThingPlug - graceful terminator
// ===========================================================================
