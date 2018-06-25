// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of ThingFan2 easy fan-in's

// ThingFan2 returns a channel to receive
// everything from the given original channel `ori`
// as well as
// all inputs
// before close.
func ThingFan2(ori <-chan Thing, inp ...Thing) (out <-chan Thing) {
	return ThingFanIn2(ori, ThingChan(inp...))
}

// ThingFan2Slice returns a channel to receive
// everything from the given original channel `ori`
// as well as
// all inputs
// before close.
func ThingFan2Slice(ori <-chan Thing, inp ...[]Thing) (out <-chan Thing) {
	return ThingFanIn2(ori, ThingChanSlice(inp...))
}

// ThingFan2Chan returns a channel to receive
// everything from the given original channel `ori`
// as well as
// from the the input channel `inp`
// before close.
// Note: ThingFan2Chan is nothing but ThingFanIn2
func ThingFan2Chan(ori <-chan Thing, inp <-chan Thing) (out <-chan Thing) {
	return ThingFanIn2(ori, inp)
}

// ThingFan2FuncNok returns a channel to receive
// everything from the given original channel `ori`
// as well as
// all results of generator `gen`
// until `!ok`
// before close.
func ThingFan2FuncNok(ori <-chan Thing, gen func() (Thing, bool)) (out <-chan Thing) {
	return ThingFanIn2(ori, ThingChanFuncNok(gen))
}

// ThingFan2FuncErr returns a channel to receive
// everything from the given original channel `ori`
// as well as
// all results of generator `gen`
// until `err != nil`
// before close.
func ThingFan2FuncErr(ori <-chan Thing, gen func() (Thing, error)) (out <-chan Thing) {
	return ThingFanIn2(ori, ThingChanFuncErr(gen))
}

// End of ThingFan2 easy fan-in's
// ===========================================================================
