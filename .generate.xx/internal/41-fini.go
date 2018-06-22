// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingFini closures

// anyThingFini returns a closure around `DoneanyThing(_)`.
func anyThingFini() func(inp Anymode) (done <-chan struct{}) {

	return func(inp Anymode) (done <-chan struct{}) {
		return DoneanyThing(inp)
	}
}

// anyThingFiniSlice returns a closure around `DoneanyThingSlice(_)`.
func anyThingFiniSlice() func(inp Anymode) (done <-chan []anyThing) {

	return func(inp Anymode) (done <-chan []anyThing) {
		return DoneanyThingSlice(inp)
	}
}

// anyThingFiniFunc returns a closure around `DoneanyThingFunc(_, act)`.
func anyThingFiniFunc(act func(a anyThing)) func(inp Anymode) (done <-chan struct{}) {

	return func(inp Anymode) (done <-chan struct{}) {
		return DoneanyThingFunc(inp, act)
	}
}

// End of anyThingFini closures
// ===========================================================================
