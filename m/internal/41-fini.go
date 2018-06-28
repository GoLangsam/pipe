// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingFini closures

// anyThingFini returns a closure around `anyThingDone()`.
func (inp anyThingRoC) anyThingFini() func(inp anyThingRoC) (done <-chan struct{}) {

	return func(inp anyThingRoC) (done <-chan struct{}) {
		return inp.anyThingDone()
	}
}

// anyThingFiniSlice returns a closure around `anyThingDoneSlice()`.
func (inp anyThingRoC) anyThingFiniSlice() func(inp anyThingRoC) (done <-chan []anyThing) {

	return func(inp anyThingRoC) (done <-chan []anyThing) {
		return inp.anyThingDoneSlice()
	}
}

// anyThingFiniFunc returns a closure around `anyThingDoneFunc(act)`.
func (inp anyThingRoC) anyThingFiniFunc(act func(a anyThing)) func(inp anyThingRoC) (done <-chan struct{}) {

	return func(inp anyThingRoC) (done <-chan struct{}) {
		return inp.anyThingDoneFunc(act)
	}
}

// End of anyThingFini closures
// ===========================================================================
