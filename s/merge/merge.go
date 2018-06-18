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
// Beg of MergeAny

// MergeAny returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be sorted ascending and free of duplicates.
// The passed binary boolean function `less` defines the applicable order.
//  Note: If no inputs are given, a closed channel is returned.
func MergeAny(less func(i, j Any) bool, inps ...<-chan Any) (out <-chan Any) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan Any)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeAny2(less, inps[0], MergeAny(less, inps[1:]...))
	}
}

// mergeAny2 takes two (eager) channels of comparable types,
// each of which needs to be sorted ascending and free of duplicates,
// and merges them into the returned channel, which will be sorted ascending and free of duplicates.
func mergeAny2(less func(i, j Any) bool, i1, i2 <-chan Any) (out <-chan Any) {
	cha := make(chan Any)
	go func(out chan<- Any, i1, i2 <-chan Any) {
		defer close(out)
		var (
			clos1, clos2 bool // we found the chan closed
			buff1, buff2 bool // we've read 'from', but not sent (yet)
			ok           bool // did we read successfully?
			from1, from2 Any  // what we've read
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

// Note: merge2 is not my own. Just: I forgot where found it - please accept my apologies.
// I'd love to learn about it's origin/author, so I can give credit.
// Thus: Your hint, dear reader, is highly appreciated!

// End of MergeAny
// ===========================================================================
