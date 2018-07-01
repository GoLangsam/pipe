// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingFini closures

// anyThingFini returns a closure around `anyThingDone(ops...)`.
func (inp anyThingFrom) anyThingFini(ops ...func(a anyThing)) func(inp anyThingFrom) (done <-chan struct{}) {

	return func(inp anyThingFrom) (done <-chan struct{}) {
		return inp.anyThingDone(ops...)
	}
}

// anyThingFiniFunc returns a closure around `anyThingDoneFunc(acts...)`.
func (inp anyThingFrom) anyThingFiniFunc(acts ...func(a anyThing)anyThing) func(inp anyThingFrom) (done <-chan struct{}) {

	return func(inp anyThingFrom) (done <-chan struct{}) {
		return inp.anyThingDoneFunc(acts...)
	}
}

// anyThingFiniSlice returns a closure around `anyThingDoneSlice()`.
func (inp anyThingFrom) anyThingFiniSlice() func(inp anyThingFrom) (done <-chan []anyThing) {

	return func(inp anyThingFrom) (done <-chan []anyThing) {
		return inp.anyThingDoneSlice()
	}
}

// End of anyThingFini closures
// ===========================================================================
