// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of ChanAny producers

// ChanAny returns a channel to receive
// all inputs
// before close.
func ChanAny(inp ...Any) chan Any {
	out := make(chan Any)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanAnySlice returns a channel to receive
// all inputs
// before close.
func ChanAnySlice(inp ...[]Any) chan Any {
	out := make(chan Any)
	go func() {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}()
	return out
}

// ChanAnyFuncNok returns a channel to receive
// all results of generator `gen`
// until `!ok`
// before close.
func ChanAnyFuncNok(gen func() (Any, bool)) chan Any {
	out := make(chan Any)
	go func() {
		defer close(out)
		for {
			res, ok := gen() // generate
			if !ok {
				return
			}
			out <- res
		}
	}()
	return out
}

// ChanAnyFuncErr returns a channel to receive
// all results of generator `gen`
// until `err != nil`
// before close.
func ChanAnyFuncErr(gen func() (Any, error)) chan Any {
	out := make(chan Any)
	go func() {
		defer close(out)
		for {
			res, err := gen() // generate
			if err != nil {
				return
			}
			out <- res
		}
	}()
	return out
}

// End of ChanAny producers
// ===========================================================================
