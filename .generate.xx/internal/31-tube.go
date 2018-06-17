// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of TubeAny closures

// TubeAnyFunc returns a closure around PipeAnyFunc (_, act).
func TubeAnyFunc(act func(a Any) Any) (tube func(inp Anymode) (out Anymode)) {

	return func(inp Anymode) (out Anymode) {
		return PipeAnyFunc(inp, act)
	}
}

// TubeAnyBuffer returns a closure around PipeAnyBuffer (_, cap).
func TubeAnyBuffer(cap int) (tube func(inp Anymode) (out Anymode)) {

	return func(inp Anymode) (out Anymode) {
		return PipeAnyBuffer(inp, cap)
	}
}

// End of TubeAny closures
// ===========================================================================
