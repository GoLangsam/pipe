// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of ThingPipeBuffered - a buffered channel with capacity `cap` to receive

// ThingPipeBuffered returns a buffered channel with capacity `cap` to receive
// all `inp`
// before close.
func ThingPipeBuffered(inp <-chan Thing, cap int) (out <-chan Thing) {
	cha := make(chan Thing, cap)
	go pipeThingBuffered(cha, inp)
	return cha
}

func pipeThingBuffered(out chan<- Thing, inp <-chan Thing) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// ThingTubeBuffered returns a closure around PipeThingBuffer (_, cap).
func ThingTubeBuffered(cap int) (tube func(inp <-chan Thing) (out <-chan Thing)) {

	return func(inp <-chan Thing) (out <-chan Thing) {
		return ThingPipeBuffered(inp, cap)
	}
}

// End of ThingPipeBuffered - a buffered channel with capacity `cap` to receive
// ===========================================================================
