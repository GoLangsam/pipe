// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of ThingMerge

// ThingMerge returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be sorted ascending and free of duplicates.
// The passed binary boolean function `less` defines the applicable order.
//  Note: If no inputs are given, a closed channel is returned.
func (inp ThingFrom) ThingMerge(less func(i, j Thing) bool, inps ...ThingFrom) (out ThingFrom) {
	var inpS []ThingFrom
	if inp == nil {
		inpS = inps
	} else {
		inpS = append(inps, inp)
	}

	if len(inpS) < 1 { // none: return a closed channel
		cha := make(chan Thing)
		defer close(cha)
		return cha
	} else if len(inpS) < 2 { // just one: return it
		return inpS[0]
	} else { // tail recurse
		return inpS[0].mergeThing(less, inpS[1].ThingMerge(less, inpS[2:]...))
	}
}

// mergeThing takes two (eager) channels of comparable types,
// each of which needs to be sorted ascending and free of duplicates,
// and merges them into the returned channel, which will be sorted ascending and free of duplicates.
func (inp ThingFrom) mergeThing(less func(i, j Thing) bool, inp2 ThingFrom) (out ThingFrom) {
	cha := make(chan Thing)
	go func(out ThingInto, inp, inp2 ThingFrom) {
		defer close(out)
		var (
			clos1, clos2 bool  // we found the chan closed
			buff1, buff2 bool  // we've read 'from', but not sent (yet)
			ok           bool  // did we read successfully?
			from1, from2 Thing // what we've read
		)

		for !clos1 || !clos2 {

			if !clos1 && !buff1 {
				if from1, ok = <-inp; ok {
					buff1 = true
				} else {
					clos1 = true
				}
			}

			if !clos2 && !buff2 {
				if from2, ok = <-inp2; ok {
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
	}(cha, inp, inp2)
	return cha
}

// Note: mergeThing is not my own.
// Just: I forgot where found the original merge2 - please accept my apologies.
// I'd love to learn about it's origin/author, so I can give credit.
// Thus: Your hint, dear reader, is highly appreciated!

// End of ThingMerge
// ===========================================================================
