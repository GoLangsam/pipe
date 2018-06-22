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
// Beg of anyThingPipeBuffered - a buffered channel with capacity `cap` to receive

// anyThingPipeBuffered returns a buffered channel with capacity `cap` to receive
// all `inp`
// before close.
func anyThingPipeBuffered(inp <-chan anyThing, cap int) (out <-chan anyThing) {
	cha := make(chan anyThing, cap)
	go func(out chan<- anyThing, inp <-chan anyThing) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// anyThingTubeBuffered returns a closure around PipeanyThingBuffer (_, cap).
func anyThingTubeBuffered(cap int) (tube func(inp <-chan anyThing) (out <-chan anyThing)) {

	return func(inp <-chan anyThing) (out <-chan anyThing) {
		return anyThingPipeBuffer(inp, cap)
	}
}

// End of anyThingPipeBuffered - a buffered channel with capacity `cap` to receive
// ===========================================================================
