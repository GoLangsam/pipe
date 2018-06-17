// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of ForkAny functions

// ForkAny returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func ForkAny(inp <-chan Any) (out1, out2 <-chan Any) {
	cha1 := make(chan Any)
	cha2 := make(chan Any)
	go func(out1, out2 chan<- Any) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2)
	return cha1, cha2
}

// End of ForkAny functions
// ===========================================================================
