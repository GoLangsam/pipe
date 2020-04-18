// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of Tube closures around Pipe

// anyThingTubeFunc returns a closure around PipeFunc (_, act).
func anyThingTubeFunc(act func(a anyThing) anyThing) (tube func(inp anyThingFrom) (out anyThingFrom)) {

	return func(inp anyThingFrom) (out anyThingFrom) {
		return inp.PipeFunc(act)
	}
}

// End of Tube closures around anyThingPipe
// ===========================================================================
