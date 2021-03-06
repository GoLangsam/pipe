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
// Beg of anyThingMerge

// anyThingMerge returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be sorted ascending and free of duplicates.
// The passed binary boolean function `less` defines the applicable order.
//  Note: If no inputs are given, a closed channel is returned.
func anyThingMerge(less func(i, j anyThing) bool, inps ...<-chan anyThing) (out <-chan anyThing) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan anyThing)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeanyThing(less, inps[0], anyThingMerge(less, inps[1:]...))
	}
}

// mergeanyThing takes two (eager) channels of comparable types,
// each of which needs to be sorted ascending and free of duplicates,
// and merges them into the returned channel, which will be sorted ascending and free of duplicates.
func mergeanyThing(less func(i, j anyThing) bool, i1, i2 <-chan anyThing) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go func(out chan<- anyThing, i1, i2 <-chan anyThing) {
		defer close(out)
		var (
			clos1, clos2 bool     // we found the chan closed
			buff1, buff2 bool     // we've read 'from', but not sent (yet)
			ok           bool     // did we read successfully?
			from1, from2 anyThing // what we've read
		)

		for !clos1 || !clos2 {

			if !clos1 && !buff1 {
				if from1, ok = <-i1; ok {
					buff1 = true
				} else {
					clos1 = true
				}
			}

			if !clos2 && !buff2 {
				if from2, ok = <-i2; ok {
					buff2 = true
				} else {
					clos2 = true
				}
			}

			if clos1 && !buff1 {
				from1 = from2
			}
			if clos2 && !buff2 {
				from2 = from1
			}

			if less(from1, from2) {
				out <- from1
				buff1 = false
			} else if less(from2, from1) {
				out <- from2
				buff2 = false
			} else {
				out <- from1 // == from2
				buff1 = false
				buff2 = false
			}
		}
	}(cha, i1, i2)
	return cha
}

// Note: mergeanyThing is not my own.
// Just: I forgot where found the original merge2 - please accept my apologies.
// I'd love to learn about it's origin/author, so I can give credit.
// Thus: Your hint, dear reader, is highly appreciated!

// End of anyThingMerge
// ===========================================================================
