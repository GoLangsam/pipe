// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of ForkAny functions

// ForkAny returns two channels
// either of which is to receive
// every result of inp
// before close.
func ForkAny(inp <-chan Any) (out1, out2 <-chan Any) {
	cha1 := make(chan Any)
	cha2 := make(chan Any)
	go forkAny(cha1, cha2, inp)
	return cha1, cha2
}

/* not used any more - kept for reference only.
func forkAny(out1, out2 chan<- Any, inp <-chan Any) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
} */

func forkAny(out1, out2 chan<- Any, inp <-chan Any) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		select { // send first to whomever is ready to receive
		case out1 <- i:
			out2 <- i
		case out2 <- i:
			out1 <- i
		}
	}
}

// End of ForkAny functions
// ===========================================================================
