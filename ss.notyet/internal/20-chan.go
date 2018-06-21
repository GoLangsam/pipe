// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingChan producers

// anyThingChan returns a channel to receive
// all inputs
// before close.
func anyThingChan(inp ...anyThing) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go func(out chan<- anyThing, inp ...anyThing) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp)
	return cha
}

// anyThingChanSlice returns a channel to receive
// all inputs
// before close.
func anyThingChanSlice(inp ...[]anyThing) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go func(out chan<- anyThing, inp ...[]anyThing) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// anyThingChanFuncNok returns a channel to receive
// all results of generator `gen`
// until `!ok`
// before close.
func anyThingChanFuncNok(gen func() (anyThing, bool)) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go func(out chan<- anyThing, gen func() (anyThing, bool)) {
		defer close(out)
		for {
			res, ok := gen() // generate
			if !ok {
				return
			}
			out <- res
		}
	}(cha, gen)
	return cha
}

// anyThingChanFuncErr returns a channel to receive
// all results of generator `gen`
// until `err != nil`
// before close.
func anyThingChanFuncErr(gen func() (anyThing, error)) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go func(out chan<- anyThing, gen func() (anyThing, error)) {
		defer close(out)
		for {
			res, err := gen() // generate
			if err != nil {
				return
			}
			out <- res
		}
	}(cha, gen)
	return cha
}

// End of anyThingChan producers
// ===========================================================================
