// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingTube closures around anyThingPipe

// anyThingTube returns a closure around PipeanyThing (_, ops...).
func anyThingTube(ops ...func(a anyThing)) (tube func(inp <-chan anyThing) (out <-chan anyThing)) {

	return func(inp <-chan anyThing) (out <-chan anyThing) {
		return anyThingPipe(inp, ops...)
	}
}

// anyThingTubeFunc returns a closure around PipeanyThingFunc (_, acts...).
func anyThingTubeFunc(acts ...func(a anyThing) anyThing) (tube func(inp <-chan anyThing) (out <-chan anyThing)) {

	return func(inp <-chan anyThing) (out <-chan anyThing) {
		return anyThingPipeFunc(inp, acts...)
	}
}

// End of anyThingTube closures around anyThingPipe
// ===========================================================================
