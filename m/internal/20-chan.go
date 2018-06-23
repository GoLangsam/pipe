// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingChan producers

// anyThingChan returns a channel to receive
// all inputs
// before close.
func (my anyOwner) anyThingChan(inp ...anyThing) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go my.chananyThing(cha, inp...)
	return cha
}

func (my anyOwner) chananyThing(out chan<- anyThing, inp ...anyThing) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// anyThingChanSlice returns a channel to receive
// all inputs
// before close.
func (my anyOwner) anyThingChanSlice(inp ...[]anyThing) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go my.chananyThingSlice(cha, inp...)
	return cha
}

func (my anyOwner) chananyThingSlice(out chan<- anyThing, inp ...[]anyThing) {
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
func (my anyOwner) anyThingChanFuncNok(gen func() (anyThing, bool)) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go my.chananyThingFuncNok(cha, gen)
	return cha
}

func (my anyOwner) chananyThingFuncNok(out chan<- anyThing, gen func() (anyThing, bool)) {
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
func (my anyOwner) anyThingChanFuncErr(gen func() (anyThing, error)) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go my.chananyThingFuncErr(cha, gen)
	return cha
}

func (my anyOwner) chananyThingFuncErr(out chan<- anyThing, gen func() (anyThing, error)) {
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
