// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingFini closures

// anyThingFini returns a closure around `anyThingDone(_)`.
func anyThingFini() func(inp chan anyThing) (done chan struct{}) {

	return func(inp chan anyThing) (done chan struct{}) {
		return anyThingDone(inp)
	}
}

// anyThingFiniSlice returns a closure around `anyThingDoneSlice(_)`.
func anyThingFiniSlice() func(inp chan anyThing) (done chan []anyThing) {

	return func(inp chan anyThing) (done chan []anyThing) {
		return anyThingDoneSlice(inp)
	}
}

// anyThingFiniFunc returns a closure around `anyThingDoneFunc(_, act)`.
func anyThingFiniFunc(act func(a anyThing)) func(inp chan anyThing) (done chan struct{}) {

	return func(inp chan anyThing) (done chan struct{}) {
		return anyThingDoneFunc(inp, act)
	}
}

// End of anyThingFini closures
// ===========================================================================
