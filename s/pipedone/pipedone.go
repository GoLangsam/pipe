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
// Beg of anyThingPipeDone

// anyThingPipeDone returns a channel to receive every `inp` before close and a channel to signal this closing.
func anyThingPipeDone(inp <-chan anyThing) (out <-chan anyThing, done <-chan struct{}) {
	cha := make(chan anyThing)
	doit := make(chan struct{})
	go pipeanyThingDone(cha, doit, inp)
	return cha, doit
}

func pipeanyThingDone(out chan<- anyThing, done chan<- struct{}, inp <-chan anyThing) {
	defer close(out)
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// End of anyThingPipeDone
// ===========================================================================
