// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// anyThingFrom is a receive-only anyThing channel
type anyThingFrom <-chan anyThing

// anyThingInto is a send-only anyThing channel
type anyThingInto chan<- anyThing

// ===========================================================================
// Beg of anyThingPipeDone

// anyThingPipeDone returns a channel to receive every `inp` before close and a channel to signal this closing.
func (inp anyThingFrom) anyThingPipeDone() (out anyThingFrom, done <-chan struct{}) {
	cha := make(chan anyThing)
	doit := make(chan struct{})
	go inp.pipeanyThingDone(cha, doit)
	return cha, doit
}

func (inp anyThingFrom) pipeanyThingDone(out anyThingInto, done chan<- struct{}) {
	defer close(out)
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// End of anyThingPipeDone
// ===========================================================================
