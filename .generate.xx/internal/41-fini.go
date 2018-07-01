// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingFini closures

// anyThingFini returns a closure around `anyThingDone(_, ops...)`.
func anyThingFini(ops ...func(a anyThing)) func(inp anymode) (done <-chan struct{}) {

	return func(inp anymode) (done <-chan struct{}) {
		return anyThingDone(inp, ops...)
	}
}

// anyThingFiniFunc returns a closure around `anyThingDoneFunc(_, acts...)`.
func anyThingFiniFunc(acts ...func(a anyThing)anyThing) func(inp anymode) (done <-chan struct{}) {

	return func(inp anymode) (done <-chan struct{}) {
		return anyThingDoneFunc(inp, acts...)
	}
}

// anyThingFiniSlice returns a closure around `anyThingDoneSlice(_)`.
func anyThingFiniSlice() func(inp anymode) (done <-chan []anyThing) {

	return func(inp anymode) (done <-chan []anyThing) {
		return anyThingDoneSlice(inp)
	}
}

// End of anyThingFini closures
// ===========================================================================
