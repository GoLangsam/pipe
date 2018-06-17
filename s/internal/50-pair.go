// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of PairAny functions

// PairAny returns a pair of channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PairAny(inp <-chan Any) (out1, out2 <-chan Any) {
	cha1 := make(chan Any)
	cha2 := make(chan Any)
	go pairAny(cha1, cha2, inp)
	return cha1, cha2
}

/* not used any more - kept for reference only.
func pairAny(out1, out2 chan<- Any, inp <-chan Any) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
} */

func pairAny(out1, out2 chan<- Any, inp <-chan Any) {
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

// End of PairAny functions
// ===========================================================================
