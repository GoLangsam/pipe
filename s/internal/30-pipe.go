// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of PipeAny functions

// PipeAnyFunc returns a channel to receive
// every result of action `act` applied to `inp`
// before close.
// Note: it 'could' be PipeAnyMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeAnyFunc(inp <-chan Any, act func(a Any) Any) (out <-chan Any) {
	cha := make(chan Any)
	if act == nil { // Make `nil` value useful
		act = func(a Any) Any { return a }
	}
	go pipeAnyFunc(cha, inp, act)
	return cha
}

func pipeAnyFunc(out chan<- Any, inp <-chan Any, act func(a Any) Any) {
	defer close(out)
	for i := range inp {
			out <- act(i) // apply action
	}
}

// PipeAnyBuffer returns a buffered channel with capacity `cap` to receive
// all `inp`
// before close.
func PipeAnyBuffer(inp <-chan Any, cap int) (out <-chan Any) {
	cha := make(chan Any, cap)
	go pipeAnyBuffer(cha, inp)
	return cha
}

func pipeAnyBuffer(out chan<- Any, inp <-chan Any) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// End of PipeAny functions
// ===========================================================================
