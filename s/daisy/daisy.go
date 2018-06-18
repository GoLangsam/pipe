// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// Any is the generic type flowing thru the pipe network.
type Any generic.Type

// ===========================================================================
// Beg of DaisyChainAny

// ProcAny is the signature of the inner process of any linear pipe-network
//  Example: the identity core:
// samesame := func(into chan<- Any, from <-chan Any) { into <- <-from }
//  Note: type ProcAny is provided for documentation purpose only.
// The implementation uses the explicit function signature
// in order to avoid some genny-related issue.
//  Note: In https://talks.golang.org/2012/waza.slide#40
//  Rob Pike uses a ProcAny named `worker`.
type ProcAny func(into chan<- Any, from <-chan Any)

// Example: the identity core - see `samesame` below
var _ ProcAny = func(into chan<- Any, from <-chan Any) { into <- <-from }

// daisyAny returns a channel to receive all inp after having passed thru process `proc`.
func daisyAny(inp <-chan Any,
	proc func(into chan<- Any, from <-chan Any), // a ProcAny process
) (
	out chan Any) { // a daisy to be chained

	cha := make(chan Any)
	go proc(cha, inp)
	return cha
}

// DaisyChainAny returns a channel to receive all inp
// after having passed
// thru the process(es) (`from` right `into` left)
// before close.
//
// Note: If no `tubes` are provided,
// `out` shall receive elements from `inp` unaltered (as a convenience),
// thus making a null value useful.
func DaisyChainAny(inp chan Any,
	procs ...func(into chan<- Any, from <-chan Any), // ProcAny processes
) (
	out chan Any) { // to receive all results

	cha := inp

	if len(procs) < 1 {
		samesame := func(into chan<- Any, from <-chan Any) { into <- <-from }
		cha = daisyAny(cha, samesame)
	} else {
		for _, proc := range procs {
			cha = daisyAny(cha, proc)
		}
	}
	return cha
}

// DaisyChaiNAny returns a channel to receive all inp
// after having passed
// `somany` times
// thru the process(es) (`from` right `into` left)
// before close.
//
// Note: If `somany` is less than 1 or no `tubes` are provided,
// `out` shall receive elements from `inp` unaltered (as a convenience),
// thus making null values useful.
//
//  Note: DaisyChaiNAny(inp, 1, procs) <==> DaisyChainAny(inp, procs)
func DaisyChaiNAny(inp chan Any, somany int,
	procs ...func(into chan<- Any, from <-chan Any), // ProcAny processes
) (
	out chan Any) { // to receive all results

	cha := inp

	if somany < 1 {
		samesame := func(into chan<- Any, from <-chan Any) { into <- <-from }
		cha = daisyAny(cha, samesame)
	} else {
		for i := 0; i < somany; i++ {
			cha = DaisyChainAny(cha, procs...)
		}
	}
	return cha
}

// End of DaisyChainAny
// ===========================================================================
