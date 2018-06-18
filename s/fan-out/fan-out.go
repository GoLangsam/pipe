// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// Any is the generic type flowing thru the pipe network.
type Any generic.Type

// ===========================================================================
// Beg of FanAnyOut

// FanAnyOut returns a slice (of size = size) of channels
// each of which shall receive any inp before close.
func FanAnyOut(inp <-chan Any, size int) (outS [](<-chan Any)) {
	chaS := make([]chan Any, size)
	for i := 0; i < size; i++ {
		chaS[i] = make(chan Any)
	}

	go fanAnyOut(inp, chaS...)

	outS = make([]<-chan Any, size)
	for i := 0; i < size; i++ {
		outS[i] = chaS[i] // convert `chan` to `<-chan`
	}

	return outS
}

// c fanAnyOut(inp <-chan Any, outs ...chan<- Any) {
func fanAnyOut(inp <-chan Any, outs ...chan Any) {

	for i := range inp {
		for o := range outs {
			outs[o] <- i
		}
	}

	for o := range outs {
		close(outs[o])
	}

}

// End of FanAnyOut
// ===========================================================================
