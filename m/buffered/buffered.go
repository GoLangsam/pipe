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
// Beg of anyThingPipeBuffered - a buffered channel with capacity `cap` to receive

// anyThingPipeBuffered returns a buffered channel with capacity `cap` to receive
// all `inp`
// before close.
func (inp anyThingFrom) anyThingPipeBuffered(cap int) (out anyThingFrom) {
	cha := make(chan anyThing, cap)
	go inp.pipeanyThingBuffered(cha)
	return cha
}

func (inp anyThingFrom) pipeanyThingBuffered(out anyThingInto) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// anyThingTubeBuffered returns a closure around PipeanyThingBuffer (cap).
func (inp anyThingFrom) anyThingTubeBuffered(cap int) (tube func(inp anyThingFrom) (out anyThingFrom)) {

	return func(inp anyThingFrom) (out anyThingFrom) {
		return inp.anyThingPipeBuffered(cap)
	}
}

// End of anyThingPipeBuffered - a buffered channel with capacity `cap` to receive
// ===========================================================================
