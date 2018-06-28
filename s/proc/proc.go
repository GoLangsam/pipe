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
// Beg of anyThingProc functions

// Problem: We can not know whether or not proc
// - terminates if inp becomes closed
// - closes out
// Solutions:
// - we request proc to be well behaved
// - we could use a send proxy, and launch proc in a separate goroutine.
//   But then we would not know when to stop waiting for answers

// anyThingPipeProc returns a channel to receive
// every result of processing function `proc` applied to `inp`
// before close.
func anyThingPipeProc(inp <-chan anyThing, proc func(into chan<- anyThing, from <-chan anyThing)) (out <-chan anyThing) {
	cha := make(chan anyThing)
	if proc == nil { // Make `nil` value useful
		proc = func(into chan<- anyThing, from <-chan anyThing) {
			// `out <- <-inp` or `into <- <-from`
			defer close(into)
			for i := range from {
				into <- i
			}
		}
	}
	go pipeanyThingProc(cha, inp, proc)
	return cha
}

func pipeanyThingProc(out chan<- anyThing, inp <-chan anyThing, proc func(into chan<- anyThing, from <-chan anyThing)) {
	defer close(out) // TODO: We do not know whether or not proc closed out
	for {
		proc(out, inp) // apply processing function
	}
}

// End of anyThingProc functions
// ===========================================================================
