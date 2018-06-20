// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"time"

	"github.com/cheekybits/genny/generic"
)

// Any is the generic type flowing thru the pipe network.
type Any generic.Type

// ===========================================================================
// Beg of StrewAny

// StrewAny returns a slice (of size = size) of channels
// one of which shall receive any inp before close.
func StrewAny(inp <-chan Any, size int) (outS [](<-chan Any)) {
	chaS := make([]chan Any, size)
	for i := 0; i < size; i++ {
		chaS[i] = make(chan Any)
	}

	go strewAny(inp, chaS...)

	outS = make([]<-chan Any, size)
	for i := 0; i < size; i++ {
		outS[i] = chaS[i] // convert `chan` to `<-chan`
	}

	return outS
}

// c strewAny(inp <-chan Any, outS ...chan<- Any) {
// Note: go does not convert the passed slice `[]chan Any` to `[]chan<- Any` automatically.
// So, we do neither here, as we are lazy (we just call an internal helper function).
func strewAny(inp <-chan Any, outS ...chan Any) {

	for i := range inp {
		for !trySendAny(i, outS...) {
			time.Sleep(time.Millisecond * 10) // wait a little before retry
		} // !sent
	} // inp

	for o := range outS {
		close(outS[o])
	}
}

func trySendAny(inp Any, outS ...chan Any) bool {

	for o := range outS {

		select { // try to send
		case outS[o] <- inp:
			return true
		default:
			// keep trying
		}

	} // outS
	return false
}

// End of StrewAny
// ===========================================================================
