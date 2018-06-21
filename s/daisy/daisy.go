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
// Beg of anyThingDaisyChain

// ProcanyThing is the signature of the inner process of any linear pipe-network
//  Example: the identity core:
// samesame := func(into chan<- anyThing, from <-chan anyThing) { into <- <-from }
//  Note: type ProcanyThing is provided for documentation purpose only.
// The implementation uses the explicit function signature
// in order to avoid some genny-related issue.
//  Note: In https://talks.golang.org/2012/waza.slide#40
//  Rob Pike uses a ProcanyThing named `worker`.
type ProcanyThing func(into chan<- anyThing, from <-chan anyThing)

// Example: the identity core - see `samesame` below
var _ ProcanyThing = func(into chan<- anyThing, from <-chan anyThing) { into <- <-from }

// daisyanyThing returns a channel to receive all inp after having passed thru process `proc`.
func daisyanyThing(inp <-chan anyThing,
	proc func(into chan<- anyThing, from <-chan anyThing), // a ProcanyThing process
) (
	out chan anyThing) { // a daisy to be chained

	cha := make(chan anyThing)
	go proc(cha, inp)
	return cha
}

// anyThingDaisyChain returns a channel to receive all inp
// after having passed
// thru the process(es) (`from` right `into` left)
// before close.
//
// Note: If no `tubes` are provided,
// `out` shall receive elements from `inp` unaltered (as a convenience),
// thus making a null value useful.
func anyThingDaisyChain(inp chan anyThing,
	procs ...func(into chan<- anyThing, from <-chan anyThing), // ProcanyThing processes
) (
	out chan anyThing) { // to receive all results

	cha := inp

	if len(procs) < 1 {
		samesame := func(into chan<- anyThing, from <-chan anyThing) { into <- <-from }
		cha = daisyanyThing(cha, samesame)
	} else {
		for _, proc := range procs {
			cha = daisyanyThing(cha, proc)
		}
	}
	return cha
}

// anyThingDaisyChaiN returns a channel to receive all inp
// after having passed
// `somany` times
// thru the process(es) (`from` right `into` left)
// before close.
//
// Note: If `somany` is less than 1 or no `tubes` are provided,
// `out` shall receive elements from `inp` unaltered (as a convenience),
// thus making null values useful.
//
//  Note: anyThingDaisyChaiN(inp, 1, procs) <==> anyThingDaisyChain(inp, procs)
func anyThingDaisyChaiN(inp chan anyThing, somany int,
	procs ...func(into chan<- anyThing, from <-chan anyThing), // ProcanyThing processes
) (
	out chan anyThing) { // to receive all results

	cha := inp

	if somany < 1 {
		samesame := func(into chan<- anyThing, from <-chan anyThing) { into <- <-from }
		cha = daisyanyThing(cha, samesame)
	} else {
		for i := 0; i < somany; i++ {
			cha = anyThingDaisyChain(cha, procs...)
		}
	}
	return cha
}

// End of anyThingDaisyChain
// ===========================================================================
