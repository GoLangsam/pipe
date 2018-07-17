// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of Fini closures

// Fini returns a closure around `Done(ops...)`.
func (inp anyThingFrom) Fini(ops ...func(a anyThing)) func(inp anyThingFrom) (done <-chan struct{}) {

	return func(inp anyThingFrom) (done <-chan struct{}) {
		return inp.Done(ops...)
	}
}

// FiniFunc returns a closure around `DoneFunc(acts...)`.
func (inp anyThingFrom) FiniFunc(acts ...func(a anyThing) anyThing) func(inp anyThingFrom) (done <-chan struct{}) {

	return func(inp anyThingFrom) (done <-chan struct{}) {
		return inp.DoneFunc(acts...)
	}
}

// FiniSlice returns a closure around `DoneSlice()`.
func (inp anyThingFrom) FiniSlice() func(inp anyThingFrom) (done <-chan []anyThing) {

	return func(inp anyThingFrom) (done <-chan []anyThing) {
		return inp.DoneSlice()
	}
}

// End of Fini closures
// ===========================================================================
