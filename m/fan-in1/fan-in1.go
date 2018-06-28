// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"time"

	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// ===========================================================================
// Beg of anyThingFanIn1 - fan-in using only one go routine

// anyThingFanIn1 returns a channel to receive all inputs arriving
// on variadic inps
// before close.
//
//  Note: Only one go routine is used for all receives,
//  which keeps trying open inputs in round-robin fashion
//  until all inputs are closed.
//
// See anyThingFanIn in `fan-in` for another implementation.
func (inp anyThingFrom) anyThingFanIn1(inpS ...anyThingFrom) (out anyThingFrom) {
	cha := make(chan anyThing)
	go fanin1anyThing(cha, append(inpS, inp)...)
	return cha
}

func fanin1anyThing(out anyThingInto, inpS ...anyThingFrom) {
	defer close(out)

	open := len(inpS)                 // assume: all are open
	closed := make([]bool, len(inpS)) // assume: each is not closed

	var item anyThing // item received
	var ok bool       // receive channel is open?
	var sent bool     // some v has been sent?

	for open > 0 {
		sent = false
		for i := range inpS {
			if !closed[i] {
				select { // try to receive
				case item, ok = <-inpS[i]:
					if ok {
						out <- item
						sent = true
					} else {
						closed[i] = true
						open--
					}
				default: // keep going
				} // try
			} // not closed
		} // inpS
		if !sent && open > 0 {
			time.Sleep(time.Millisecond * 10) // wait a little before retry
		}
	} // open
}

// End of anyThingFanIn1 - fan-in using only one go routine
// ===========================================================================
