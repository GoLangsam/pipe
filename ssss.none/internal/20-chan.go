// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingChan producers

// anyThingChan returns a channel to receive
// all inputs
// before close.
func anyThingChan(inp ...anyThing) chan anyThing {
	out := make(chan anyThing)
	go func() {
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// anyThingChanSlice returns a channel to receive
// all inputs
// before close.
func anyThingChanSlice(inp ...[]anyThing) chan anyThing {
	out := make(chan anyThing)
	go func() {
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}()
	return out
}

// anyThingChanFuncNok returns a channel to receive
// all results of generator `gen`
// until `!ok`
// before close.
func anyThingChanFuncNok(gen func() (anyThing, bool)) chan anyThing {
	out := make(chan anyThing)
	go func() {
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

// anyThingChanFuncErr returns a channel to receive
// all results of generator `gen`
// until `err != nil`
// before close.
func anyThingChanFuncErr(gen func() (anyThing, error)) chan anyThing {
	out := make(chan anyThing)
	go func() {
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

// End of anyThingChan producers
// ===========================================================================
