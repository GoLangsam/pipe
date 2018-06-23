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
// Beg of anyThingPipeBuffered - a buffered channel with capacity `cap` to receive

// anyThingPipeBuffered returns a buffered channel with capacity `cap` to receive
// all `inp`
// before close.
func (my anyOwner) anyThingPipeBuffered(inp <-chan anyThing, cap int) (out <-chan anyThing) {
	cha := make(chan anyThing, cap)
	go my.pipeanyThingBuffered(cha, inp)
	return cha
}

func (my anyOwner) pipeanyThingBuffered(out chan<- anyThing, inp <-chan anyThing) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// anyThingTubeBuffered returns a closure around PipeanyThingBuffer (_, cap).
func (my anyOwner) anyThingTubeBuffered(cap int) (tube func(inp <-chan anyThing) (out <-chan anyThing)) {

	return func(inp <-chan anyThing) (out <-chan anyThing) {
		return my.anyThingPipeBuffered(inp, cap)
	}
}

// End of anyThingPipeBuffered - a buffered channel with capacity `cap` to receive
// ===========================================================================
