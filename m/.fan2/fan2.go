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
// Beg of Fan2 easy fan-in's

// Fan2 returns a channel to receive
// everything from `inp`
// as well as
// all inputs
// before close.
func (inp anyThingFrom) Fan2(inps ...anyThing) (out anyThingFrom) {
	return inp.FanIn2(anyThingChan(inps...))
}

// Fan2Slice returns a channel to receive
// everything from `inp`
// as well as
// all inputs
// before close.
func (inp anyThingFrom) Fan2Slice(inps ...[]anyThing) (out anyThingFrom) {
	return inp.FanIn2(anyThingChanSlice(inps...))
}

// Fan2Chan returns a channel to receive
// everything from `inp`
// as well as
// everything from `inp2`
// before close.
//  Note: Fan2Chan is nothing but FanIn2
func (inp anyThingFrom) Fan2Chan(inp2 anyThingFrom) (out anyThingFrom) {
	return inp.FanIn2(inp2)
}

// Fan2FuncNok returns a channel to receive
// everything from `inp`
// as well as
// all results of generator `gen`
// until `!ok`
// before close.
func (inp anyThingFrom) Fan2FuncNok(gen func() (anyThing, bool)) (out anyThingFrom) {
	return inp.FanIn2(anyThingChanFuncNok(gen))
}

// Fan2FuncErr returns a channel to receive
// everything from `inp`
// as well as
// all results of generator `gen`
// until `err != nil`
// before close.
func (inp anyThingFrom) Fan2FuncErr(gen func() (anyThing, error)) (out anyThingFrom) {
	return inp.FanIn2(anyThingChanFuncErr(gen))
}

// End of Fan2 easy fan-in's
// ===========================================================================
