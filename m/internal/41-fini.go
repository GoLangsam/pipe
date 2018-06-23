// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingFini closures

// anyThingFini returns a closure around `anyThingDone(_)`.
func (my anyOwner) anyThingFini() func(inp <-chan anyThing) (done <-chan struct{}) {

	return func(inp <-chan anyThing) (done <-chan struct{}) {
		return my.anyThingDone(inp)
	}
}

// anyThingFiniSlice returns a closure around `anyThingDoneSlice(_)`.
func (my anyOwner) anyThingFiniSlice() func(inp <-chan anyThing) (done <-chan []anyThing) {

	return func(inp <-chan anyThing) (done <-chan []anyThing) {
		return my.anyThingDoneSlice(inp)
	}
}

// anyThingFiniFunc returns a closure around `anyThingDoneFunc(_, act)`.
func (my anyOwner) anyThingFiniFunc(act func(a anyThing)) func(inp <-chan anyThing) (done <-chan struct{}) {

	return func(inp <-chan anyThing) (done <-chan struct{}) {
		return my.anyThingDoneFunc(inp, act)
	}
}

// End of anyThingFini closures
// ===========================================================================
