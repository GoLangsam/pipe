// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of PipeAny functions

// PipeAnyFunc returns a channel to receive
// every result of action `act` applied to `inp`
// before close.
//  Note: it 'could' be PipeAnyMap for functional people,
//  but 'map' has a very different meaning in go lang.
func PipeAnyFunc(inp chan Any, act func(a Any) Any) chan Any {
	out := make(chan Any)
	if act == nil { // Make `nil` value useful
		act = func(a Any) Any { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i) // apply action
		}
	}()
	return out
}

// PipeAnyBuffer returns a buffered channel with capacity `cap` to receive
// all `inp`
// before close.
func PipeAnyBuffer(inp chan Any, cap int) chan Any {
	out := make(chan Any, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// End of PipeAny functions
// ===========================================================================
