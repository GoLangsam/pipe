// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// anyOwner is the generic who shall own the methods.
//  Note: Need to use `generic.Number` here as `generic.Type` is an interface and cannot have any method.
type anyOwner generic.Number

// ===========================================================================
// Beg of anyThingDaisyChain

// anyThingProc is the signature of the inner process of any linear pipe-network
//  Example: the identity proc:
// samesame := func(into chan<- anyThing, from <-chan anyThing) { into <- <-from }
//  Note: type anyThingProc is provided for documentation purpose only.
// The implementation uses the explicit function signature
// in order to avoid some genny-related issue.
//  Note: In https://talks.golang.org/2012/waza.slide#40
//  Rob Pike uses a anyThingProc named `worker`.
type anyThingProc func(into chan<- anyThing, from <-chan anyThing)

// Example: the identity proc - see `samesame` below
var _ anyThingProc = func(out chan<- anyThing, inp <-chan anyThing) {
	// `out <- <-inp` or `into <- <-from`
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// daisyanyThing returns a channel to receive all inp after having passed thru process `proc`.
func (my anyOwner) daisyanyThing(inp <-chan anyThing,
	proc func(into chan<- anyThing, from <-chan anyThing), // a anyThingProc process
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
func (my anyOwner) anyThingDaisyChain(inp chan anyThing,
	procs ...func(out chan<- anyThing, inp <-chan anyThing), // anyThingProc processes
) (
	out chan anyThing) { // to receive all results

	cha := inp

	if len(procs) < 1 {
		samesame := func(out chan<- anyThing, inp <-chan anyThing) {
			// `out <- <-inp` or `into <- <-from`
			defer close(out)
			for i := range inp {
				out <- i
			}
		}
		cha = my.daisyanyThing(cha, samesame)
	} else {
		for _, proc := range procs {
			cha = my.daisyanyThing(cha, proc)
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
func (my anyOwner) anyThingDaisyChaiN(inp chan anyThing, somany int,
	procs ...func(out chan<- anyThing, inp <-chan anyThing), // ProcanyThing processes
) (
	out chan anyThing) { // to receive all results

	cha := inp

	if somany < 1 {
		samesame := func(out chan<- anyThing, inp <-chan anyThing) {
			// `out <- <-inp` or `into <- <-from`
			defer close(out)
			for i := range inp {
				out <- i
			}
		}
		cha = my.daisyanyThing(cha, samesame)
	} else {
		for i := 0; i < somany; i++ {
			cha = my.anyThingDaisyChain(cha, procs...)
		}
	}
	return cha
}

// End of anyThingDaisyChain
// ===========================================================================