// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of FiniAny closures

// FiniAny returns a closure around `DoneAny(_)`.
func FiniAny() func(inp Anymode) (done <-chan struct{}) {

	return func(inp Anymode) (done <-chan struct{}) {
		return DoneAny(inp)
	}
}

// FiniAnySlice returns a closure around `DoneAnySlice(_)`.
func FiniAnySlice() func(inp Anymode) (done <-chan []Any) {

	return func(inp Anymode) (done <-chan []Any) {
		return DoneAnySlice(inp)
	}
}

// FiniAnyFunc returns a closure around `DoneAnyFunc(_, act)`.
func FiniAnyFunc(act func(a Any)) func(inp Anymode) (done <-chan struct{}) {

	return func(inp Anymode) (done <-chan struct{}) {
		return DoneAnyFunc(inp, act)
	}
}

// End of FiniAny closures
// ===========================================================================
