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
// everything from `inp`
// as well as
// all inputs
// before close.
func (inp anyThingFrom) anyThingFan2(inps ...anyThing) (out anyThingFrom) {
	return inp.anyThingFanIn2(anyThingChan(inps...))
}

// anyThingFan2Slice returns a channel to receive
// everything from `inp`
// as well as
// all inputs
// before close.
func (inp anyThingFrom) anyThingFan2Slice(inps ...[]anyThing) (out anyThingFrom) {
	return inp.anyThingFanIn2(anyThingChanSlice(inps...))
}

// anyThingFan2Chan returns a channel to receive
// everything from `inp`
// as well as
// everything from `inp2`
// before close.
//  Note: anyThingFan2Chan is nothing but anyThingFanIn2
func (inp anyThingFrom) anyThingFan2Chan(inp2 anyThingFrom) (out anyThingFrom) {
	return inp.anyThingFanIn2(inp2)
}

// anyThingFan2FuncNok returns a channel to receive
// everything from `inp`
// as well as
// all results of generator `gen`
// until `!ok`
// before close.
func (inp anyThingFrom) anyThingFan2FuncNok(gen func() (anyThing, bool)) (out anyThingFrom) {
	return inp.anyThingFanIn2(anyThingChanFuncNok(gen))
}

// anyThingFan2FuncErr returns a channel to receive
// everything from `inp`
// as well as
// all results of generator `gen`
// until `err != nil`
// before close.
func (inp anyThingFrom) anyThingFan2FuncErr(gen func() (anyThing, error)) (out anyThingFrom) {
	return inp.anyThingFanIn2(anyThingChanFuncErr(gen))
}

// End of anyThingFan2 easy fan-in's
// ===========================================================================
