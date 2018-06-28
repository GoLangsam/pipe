// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingChan producers

// anyThingChan returns a channel to receive
// all inputs
// before close.
func anyThingChan(inp ...anyThing) (out anyThingFrom) {
	cha := make(chan anyThing)
	go chananyThing(cha, inp...)
	return cha
}

func chananyThing(out anyThingInto, inp ...anyThing) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// anyThingChanSlice returns a channel to receive
// all inputs
// before close.
func anyThingChanSlice(inp ...[]anyThing) (out anyThingFrom) {
	cha := make(chan anyThing)
	go chananyThingSlice(cha, inp...)
	return cha
}

func chananyThingSlice(out anyThingInto, inp ...[]anyThing) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// anyThingChanFuncNok returns a channel to receive
// all results of generator `gen`
// until `!ok`
// before close.
func anyThingChanFuncNok(gen func() (anyThing, bool)) (out anyThingFrom) {
	cha := make(chan anyThing)
	go chananyThingFuncNok(cha, gen)
	return cha
}

func chananyThingFuncNok(out anyThingInto, gen func() (anyThing, bool)) {
	defer close(out)
	for {
		res, ok := gen() // generate
		if !ok {
			return
		}
		out <- res
	}
}

// anyThingChanFuncErr returns a channel to receive
// all results of generator `gen`
// until `err != nil`
// before close.
func anyThingChanFuncErr(gen func() (anyThing, error)) (out anyThingFrom) {
	cha := make(chan anyThing)
	go chananyThingFuncErr(cha, gen)
	return cha
}

func chananyThingFuncErr(out anyThingInto, gen func() (anyThing, error)) {
	defer close(out)
	for {
		res, err := gen() // generate
		if err != nil {
			return
		}
		out <- res
	}
}

// End of anyThingChan producers
// ===========================================================================
