// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

type Any generic.Type

// ===========================================================================
// Beg of Fan2Any fan-in's

// Fan2Any returns a channel to receive
// everything from the given original channel `ori`
// as well as
// all inputs
// before close.
func Fan2Any(ori <-chan Any, inp ...Any) (out <-chan Any) {
	return FanIn2Any(ori, ChanAny(inp...))
}

// Fan2AnySlice returns a channel to receive
// everything from the given original channel `ori`
// as well as
// all inputs
// before close.
func Fan2AnySlice(ori <-chan Any, inp ...[]Any) (out <-chan Any) {
	return FanIn2Any(ori, ChanAnySlice(inp...))
}

// Fan2AnyChan returns a channel to receive
// everything from the given original channel `ori`
// as well as
// from the the input channel `inp`
// before close.
//  Note: Fan2AnyChan is nothing but FanIn2Any
func Fan2AnyChan(ori <-chan Any, inp <-chan Any) (out <-chan Any) {
	return FanIn2Any(ori, inp)
}

// Fan2AnyFuncNok returns a channel to receive
// everything from the given original channel `ori`
// as well as
// all results of generator `gen`
// until `!ok`
// before close.
func Fan2AnyFuncNok(ori <-chan Any, gen func() (Any, bool)) (out <-chan Any) {
	return FanIn2Any(ori, ChanAnyFuncNok(gen))
}

// Fan2AnyFuncErr returns a channel to receive
// everything from the given original channel `ori`
// as well as
// all results of generator `gen`
// until `err != nil`
// before close.
func Fan2AnyFuncErr(ori <-chan Any, gen func() (Any, error)) (out <-chan Any) {
	return FanIn2Any(ori, ChanAnyFuncErr(gen))
}

// End of Fan2Any fan-in's
// ===========================================================================
