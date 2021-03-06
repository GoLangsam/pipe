// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// ===========================================================================
// Beg of Plug - graceful terminator

// Plug returns a channel to receive every `inp` before close and a channel to signal this closing.
// Upon receipt of a stop signal,
// output is immediately closed,
// and for graceful termination
// any remaining input is drained before done is signalled.
func (inp anyThingFrom) Plug(stop <-chan struct{}) (out anyThingFrom, done <-chan struct{}) {
	cha := make(chan anyThing)
	doit := make(chan struct{})
	go inp.plug(cha, doit, stop)
	return cha, doit
}

func (inp anyThingFrom) plug(out anyThingInto, done chan<- struct{}, stop <-chan struct{}) {
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

// End of Plug - graceful terminator
// ===========================================================================
