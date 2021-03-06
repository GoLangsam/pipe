// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by golang.org/x/tools/cmd/bundle. DO NOT EDIT.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// anyThingFrom is a receive-only anyThing channel
type anyThingFrom <-chan anyThing

// anyThingInto is a send-only anyThing channel
type anyThingInto chan<- anyThing

// ===========================================================================
// Beg of anyThingMake creators

// anyThingMakeChan returns a new open channel
// (simply a 'chan anyThing' that is).
//  Note: No 'anyThing-producer' is launched here yet! (as is in all the other functions).
//  This is useful to easily create corresponding variables such as:
/*
   var myanyThingPipelineStartsHere := anyThingMakeChan()
   // ... lot's of code to design and build Your favourite "myanyThingWorkflowPipeline"
   // ...
   // ... *before* You start pouring data into it, e.g. simply via:
   for drop := range water {
       myanyThingPipelineStartsHere <- drop
   }
   close(myanyThingPipelineStartsHere)
*/
//  Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
//  (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
//  Note: as always (except for anyThingPipeBuffer) the channel is unbuffered.
//
func anyThingMakeChan() (out chan anyThing) {
	return make(chan anyThing)
}

// End of anyThingMake creators
// ===========================================================================

// ===========================================================================
// Beg of anyThingChan producers

// anyThingChan returns a channel to receive
// all inputs
// before close.
func anyThingChan(inp ...anyThing) (out anyThingFrom) {
	cha := make(chan anyThing)
	go chananyThing(cha, inp...)
	return cha
}

func chananyThing(out anyThingInto, inp ...anyThing) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// anyThingChanSlice returns a channel to receive
// all inputs
// before close.
func anyThingChanSlice(inp ...[]anyThing) (out anyThingFrom) {
	cha := make(chan anyThing)
	go chananyThingSlice(cha, inp...)
	return cha
}

func chananyThingSlice(out anyThingInto, inp ...[]anyThing) {
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
func anyThingChanFuncNok(gen func() (anyThing, bool)) (out anyThingFrom) {
	cha := make(chan anyThing)
	go chananyThingFuncNok(cha, gen)
	return cha
}

func chananyThingFuncNok(out anyThingInto, gen func() (anyThing, bool)) {
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
func anyThingChanFuncErr(gen func() (anyThing, error)) (out anyThingFrom) {
	cha := make(chan anyThing)
	go chananyThingFuncErr(cha, gen)
	return cha
}

func chananyThingFuncErr(out anyThingInto, gen func() (anyThing, error)) {
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

// ===========================================================================
// Beg of Pipe functions

// PipeFunc returns a channel to receive
// every result of action `act` applied to `inp`
// before close.
// Note: it 'could' be PipeMap for functional people,
// but 'map' has a very different meaning in go lang.
func (inp anyThingFrom) PipeFunc(act func(a anyThing) anyThing) (out anyThingFrom) {
	cha := make(chan anyThing)
	if act == nil { // Make `nil` value useful
		act = func(a anyThing) anyThing { return a }
	}
	go inp.pipeFunc(cha, act)
	return cha
}

func (inp anyThingFrom) pipeFunc(out anyThingInto, act func(a anyThing) anyThing) {
	defer close(out)
	for i := range inp {
		out <- act(i) // apply action
	}
}

// End of Pipe functions
// ===========================================================================

// ===========================================================================
// Beg of Tube closures around Pipe

// anyThingTubeFunc returns a closure around PipeFunc (_, act).
func anyThingTubeFunc(act func(a anyThing) anyThing) (tube func(inp anyThingFrom) (out anyThingFrom)) {

	return func(inp anyThingFrom) (out anyThingFrom) {
		return inp.PipeFunc(act)
	}
}

// End of Tube closures around anyThingPipe
// ===========================================================================

// ===========================================================================
// Beg of Done terminators

// Done
// will apply every `op` to every `inp` and
// returns a channel to receive
// one signal
// upon close.
func (inp anyThingFrom) Done(ops ...func(a anyThing)) (done <-chan struct{}) {
	sig := make(chan struct{})
	go inp.done(sig, ops...)
	return sig
}

func (inp anyThingFrom) done(done chan<- struct{}, ops ...func(a anyThing)) {
	defer close(done)
	for i := range inp {
		for _, op := range ops {
			if op != nil {
				op(i) // apply operation
			}
		}
	}
	done <- struct{}{}
}

// DoneFunc
// will chain every `act` to every `inp` and
// returns a channel to receive
// one signal
// upon close.
func (inp anyThingFrom) DoneFunc(acts ...func(a anyThing) anyThing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go inp.doneFunc(sig, acts...)
	return sig
}

func (inp anyThingFrom) doneFunc(done chan<- struct{}, acts ...func(a anyThing) anyThing) {
	defer close(done)
	for i := range inp {
		for _, act := range acts {
			if act != nil {
				i = act(i) // chain action
			}
		}
	}
	done <- struct{}{}
}

// DoneSlice returns a channel to receive
// a slice with every anyThing received on `inp`
// upon close.
//
//  Note: Unlike Done, DoneSlice sends the fully accumulated slice, not just an event, once upon close of inp.
func (inp anyThingFrom) DoneSlice() (done <-chan []anyThing) {
	sig := make(chan []anyThing)
	go inp.doneSlice(sig)
	return sig
}

func (inp anyThingFrom) doneSlice(done chan<- []anyThing) {
	defer close(done)
	slice := []anyThing{}
	for i := range inp {
		slice = append(slice, i)
	}
	done <- slice
}

// End of Done terminators
// ===========================================================================

// ===========================================================================
// Beg of Fini closures

// Fini returns a closure around `Done(ops...)`.
func (inp anyThingFrom) Fini(ops ...func(a anyThing)) func(inp anyThingFrom) (done <-chan struct{}) {

	return func(inp anyThingFrom) (done <-chan struct{}) {
		return inp.Done(ops...)
	}
}

// FiniFunc returns a closure around `DoneFunc(acts...)`.
func (inp anyThingFrom) FiniFunc(acts ...func(a anyThing) anyThing) func(inp anyThingFrom) (done <-chan struct{}) {

	return func(inp anyThingFrom) (done <-chan struct{}) {
		return inp.DoneFunc(acts...)
	}
}

// FiniSlice returns a closure around `DoneSlice()`.
func (inp anyThingFrom) FiniSlice() func(inp anyThingFrom) (done <-chan []anyThing) {

	return func(inp anyThingFrom) (done <-chan []anyThing) {
		return inp.DoneSlice()
	}
}

// End of Fini closures
// ===========================================================================

// ===========================================================================
// Beg of Pair functions

// Pair returns a pair of channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func (inp anyThingFrom) Pair() (out1, out2 anyThingFrom) {
	cha1 := make(chan anyThing)
	cha2 := make(chan anyThing)
	go inp.pair(cha1, cha2)
	return cha1, cha2
}

/* not used - kept for reference only.
func (inp anyThingFrom) pair(out1, out2 anyThingInto, inp anyThingFrom) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
} */

func (inp anyThingFrom) pair(out1, out2 anyThingInto) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		select { // send first to whomever is ready to receive
		case out1 <- i:
			out2 <- i
		case out2 <- i:
			out1 <- i
		}
	}
}

// End of Pair functions
// ===========================================================================

// ===========================================================================
// Beg of Fork functions

// Fork returns two channels
// either of which is to receive
// every result of inp
// before close.
func (inp anyThingFrom) Fork() (out1, out2 anyThingFrom) {
	cha1 := make(chan anyThing)
	cha2 := make(chan anyThing)
	go inp.fork(cha1, cha2)
	return cha1, cha2
}

/* not used - kept for reference only.
func (inp anyThingFrom) fork(out1, out2 anyThingInto) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
} */

func (inp anyThingFrom) fork(out1, out2 anyThingInto) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		select { // send first to whomever is ready to receive
		case out1 <- i:
			out2 <- i
		case out2 <- i:
			out1 <- i
		}
	}
}

// End of Fork functions
// ===========================================================================

// ===========================================================================
// Beg of FanIn2 simple binary Fan-In

// FanIn2 returns a channel to receive
// all from both `inp` and `inp2`
// before close.
func (inp anyThingFrom) FanIn2(inp2 anyThingFrom) (out anyThingFrom) {
	cha := make(chan anyThing)
	go inp.fanIn2(cha, inp2)
	return cha
}

/* not used - kept for reference only.
// (inp anyThingFrom) fanin2 as seen in Go Concurrency Patterns
func fanin2(out anyThingInto, inp, inp2 anyThingFrom) {
	for {
		select {
		case e := <-inp:
			out <- e
		case e := <-inp2:
			out <- e
		}
	}
} */

func (inp anyThingFrom) fanIn2(out anyThingInto, inp2 anyThingFrom) {
	defer close(out)

	var (
		closed bool     // we found a chan closed
		ok     bool     // did we read successfully?
		e      anyThing // what we've read
	)

	for !closed {
		select {
		case e, ok = <-inp:
			if ok {
				out <- e
			} else {
				inp = inp2    // swap inp2 into inp
				closed = true // break out of the loop
			}
		case e, ok = <-inp2:
			if ok {
				out <- e
			} else {
				closed = true // break out of the loop				}
			}
		}
	}

	// inp might not be closed yet. Drain it.
	for e = range inp {
		out <- e
	}
}

// End of FanIn2 simple binary Fan-In
// ===========================================================================
