// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingChan producers

// anyThingChan returns a channel to receive
// all inputs
// before close.
func anyThingChan(inp ...anyThing) (out anymode) {
	cha := anymodeMakeChan()
	go chananyThing(cha, inp...)
	return cha
}

func chananyThing(out anymode, inp ...anyThing) {
	defer out.Close()
	for i := range inp {
		out.Provide(inp[i])
	}
}

// anyThingChanSlice returns a channel to receive
// all inputs
// before close.
func anyThingChanSlice(inp ...[]anyThing) (out anymode) {
	cha := anymodeMakeChan()
	go chananyThingSlice(cha, inp...)
	return cha
}

func chananyThingSlice(out anymode, inp ...[]anyThing) {
	defer out.Close()
	for i := range inp {
		for j := range inp[i] {
			out.Provide(inp[i][j])
		}
	}
}

// anyThingChanFuncNok returns a channel to receive
// all results of generator `gen`
// until `!ok`
// before close.
func anyThingChanFuncNok(gen func() (anyThing, bool)) (out anymode) {
	cha := anymodeMakeChan()
	go chananyThingFuncNok(cha, gen)
	return cha
}

func chananyThingFuncNok(out anymode, gen func() (anyThing, bool)) {
	defer out.Close()
	for {
		res, ok := gen() // generate
		if !ok {
			return
		}
		out.Provide(res)
	}
}

// anyThingChanFuncErr returns a channel to receive
// all results of generator `gen`
// until `err != nil`
// before close.
func anyThingChanFuncErr(gen func() (anyThing, error)) (out anymode) {
	cha := anymodeMakeChan()
	go chananyThingFuncErr(cha, gen)
	return cha
}

func chananyThingFuncErr(out anymode, gen func() (anyThing, error)) {
	defer out.Close()
	for {
		res, err := gen() // generate
		if err != nil {
			return
		}
		out.Provide(res)
	}
}

// End of anyThingChan producers
// ===========================================================================
