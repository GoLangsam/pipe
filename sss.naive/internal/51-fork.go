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
func ForkAny(inp chan Any) (chan Any, chan Any) {
	out1 := make(chan Any)
	out2 := make(chan Any)
	go func() {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			select { // send to whomever is ready to receive
			case out1 <- i:
			case out2 <- i:
			}
		}
	}()
	return out1, out2
}

// End of ForkAny functions
// ===========================================================================
