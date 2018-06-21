// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// ===========================================================================
// Beg of anyThingFan2 easy fan-in's

// anyThingFan2 returns a channel to receive
// everything from the given original channel `ori`
// as well as
// all inputs
// before close.
func anyThingFan2(ori <-chan anyThing, inp ...anyThing) (out <-chan anyThing) {
	return anyThingFanIn2(ori, anyThingChan(inp...))
}

// anyThingFan2Slice returns a channel to receive
// everything from the given original channel `ori`
// as well as
// all inputs
// before close.
func anyThingFan2Slice(ori <-chan anyThing, inp ...[]anyThing) (out <-chan anyThing) {
	return anyThingFanIn2(ori, anyThingChanSlice(inp...))
}

// anyThingFan2Chan returns a channel to receive
// everything from the given original channel `ori`
// as well as
// from the the input channel `inp`
// before close.
//  Note: anyThingFan2Chan is nothing but anyThingFanIn2
func anyThingFan2Chan(ori <-chan anyThing, inp <-chan anyThing) (out <-chan anyThing) {
	return anyThingFanIn2(ori, inp)
}

// anyThingFan2FuncNok returns a channel to receive
// everything from the given original channel `ori`
// as well as
// all results of generator `gen`
// until `!ok`
// before close.
func anyThingFan2FuncNok(ori <-chan anyThing, gen func() (anyThing, bool)) (out <-chan anyThing) {
	return anyThingFanIn2(ori, anyThingChanFuncNok(gen))
}

// anyThingFan2FuncErr returns a channel to receive
// everything from the given original channel `ori`
// as well as
// all results of generator `gen`
// until `err != nil`
// before close.
func anyThingFan2FuncErr(ori <-chan anyThing, gen func() (anyThing, error)) (out <-chan anyThing) {
	return anyThingFanIn2(ori, anyThingChanFuncErr(gen))
}

// End of anyThingFan2 easy fan-in's
// ===========================================================================
