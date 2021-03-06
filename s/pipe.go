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

// ===========================================================================
// Beg of anyThingMake creators

// anyThingMakeChan returns a new open channel
// (simply a 'chan anyThing' that is).
//
// Note: No 'anyThing-producer' is launched here yet! (as is in all the other functions).
//  This is useful to easily create corresponding variables such as:
//
// 	var myanyThingPipelineStartsHere := anyThingMakeChan()
// 	// ... lot's of code to design and build Your favourite "myanyThingWorkflowPipeline"
// 	// ...
// 	// ... *before* You start pouring data into it, e.g. simply via:
// 	for drop := range water {
// 	    myanyThingPipelineStartsHere <- drop
// 	}
// 	close(myanyThingPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for anyThingPipeBuffer) the channel is unbuffered.
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
func anyThingChan(inp ...anyThing) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go chananyThing(cha, inp...)
	return cha
}

func chananyThing(out chan<- anyThing, inp ...anyThing) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// anyThingChanSlice returns a channel to receive
// all inputs
// before close.
func anyThingChanSlice(inp ...[]anyThing) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go chananyThingSlice(cha, inp...)
	return cha
}

func chananyThingSlice(out chan<- anyThing, inp ...[]anyThing) {
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
func anyThingChanFuncNok(gen func() (anyThing, bool)) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go chananyThingFuncNok(cha, gen)
	return cha
}

func chananyThingFuncNok(out chan<- anyThing, gen func() (anyThing, bool)) {
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
func anyThingChanFuncErr(gen func() (anyThing, error)) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go chananyThingFuncErr(cha, gen)
	return cha
}

func chananyThingFuncErr(out chan<- anyThing, gen func() (anyThing, error)) {
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
// Beg of anyThingPipe functions

// anyThingPipe
// will apply every `op` to every `inp` and
// returns a channel to receive
// each `inp`
// before close.
//
// Note: For functional people,
// this 'could' be named `anyThingMap`.
// Just: 'map' has a very different meaning in go lang.
func anyThingPipe(inp <-chan anyThing, ops ...func(a anyThing)) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go pipeanyThing(cha, inp, ops...)
	return cha
}

func pipeanyThing(out chan<- anyThing, inp <-chan anyThing, ops ...func(a anyThing)) {
	defer close(out)
	for i := range inp {
		for _, op := range ops {
			if op != nil {
				op(i) // chain action
			}
		}
		out <- i // send it
	}
}

// anyThingPipeFunc
// will chain every `act` to every `inp` and
// returns a channel to receive
// each result
// before close.
func anyThingPipeFunc(inp <-chan anyThing, acts ...func(a anyThing) anyThing) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go pipeanyThingFunc(cha, inp, acts...)
	return cha
}

func pipeanyThingFunc(out chan<- anyThing, inp <-chan anyThing, acts ...func(a anyThing) anyThing) {
	defer close(out)
	for i := range inp {
		for _, act := range acts {
			if act != nil {
				i = act(i) // chain action
			}
		}
		out <- i // send result
	}
}

// End of anyThingPipe functions
// ===========================================================================

// ===========================================================================
// Beg of anyThingTube closures around anyThingPipe

// anyThingTube returns a closure around PipeanyThing (_, ops...).
func anyThingTube(ops ...func(a anyThing)) (tube func(inp <-chan anyThing) (out <-chan anyThing)) {

	return func(inp <-chan anyThing) (out <-chan anyThing) {
		return anyThingPipe(inp, ops...)
	}
}

// anyThingTubeFunc returns a closure around PipeanyThingFunc (_, acts...).
func anyThingTubeFunc(acts ...func(a anyThing) anyThing) (tube func(inp <-chan anyThing) (out <-chan anyThing)) {

	return func(inp <-chan anyThing) (out <-chan anyThing) {
		return anyThingPipeFunc(inp, acts...)
	}
}

// End of anyThingTube closures around anyThingPipe
// ===========================================================================

// ===========================================================================
// Beg of anyThingDone terminators

// anyThingDone
// will apply every `op` to every `inp` and
// returns a channel to receive
// one signal
// upon close.
func anyThingDone(inp <-chan anyThing, ops ...func(a anyThing)) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doneanyThing(sig, inp, ops...)
	return sig
}

func doneanyThing(done chan<- struct{}, inp <-chan anyThing, ops ...func(a anyThing)) {
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

// anyThingDoneFunc
// will chain every `act` to every `inp` and
// returns a channel to receive
// one signal
// upon close.
func anyThingDoneFunc(inp <-chan anyThing, acts ...func(a anyThing) anyThing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doneanyThingFunc(sig, inp, acts...)
	return sig
}

func doneanyThingFunc(done chan<- struct{}, inp <-chan anyThing, acts ...func(a anyThing) anyThing) {
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

// anyThingDoneSlice returns a channel to receive
// a slice with every anyThing received on `inp`
// upon close.
//
//  Note: Unlike anyThingDone, anyThingDoneSlice sends the fully accumulated slice, not just an event, once upon close of inp.
func anyThingDoneSlice(inp <-chan anyThing) (done <-chan []anyThing) {
	sig := make(chan []anyThing)
	go doneanyThingSlice(sig, inp)
	return sig
}

func doneanyThingSlice(done chan<- []anyThing, inp <-chan anyThing) {
	defer close(done)
	slice := []anyThing{}
	for i := range inp {
		slice = append(slice, i)
	}
	done <- slice
}

// End of anyThingDone terminators
// ===========================================================================

// ===========================================================================
// Beg of anyThingFini closures

// anyThingFini returns a closure around `anyThingDone(_, ops...)`.
func anyThingFini(ops ...func(a anyThing)) func(inp <-chan anyThing) (done <-chan struct{}) {

	return func(inp <-chan anyThing) (done <-chan struct{}) {
		return anyThingDone(inp, ops...)
	}
}

// anyThingFiniFunc returns a closure around `anyThingDoneFunc(_, acts...)`.
func anyThingFiniFunc(acts ...func(a anyThing) anyThing) func(inp <-chan anyThing) (done <-chan struct{}) {

	return func(inp <-chan anyThing) (done <-chan struct{}) {
		return anyThingDoneFunc(inp, acts...)
	}
}

// anyThingFiniSlice returns a closure around `anyThingDoneSlice(_)`.
func anyThingFiniSlice() func(inp <-chan anyThing) (done <-chan []anyThing) {

	return func(inp <-chan anyThing) (done <-chan []anyThing) {
		return anyThingDoneSlice(inp)
	}
}

// End of anyThingFini closures
// ===========================================================================

// ===========================================================================
// Beg of anyThingPair functions

// anyThingPair returns a pair of channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func anyThingPair(inp <-chan anyThing) (out1, out2 <-chan anyThing) {
	cha1 := make(chan anyThing)
	cha2 := make(chan anyThing)
	go pairanyThing(cha1, cha2, inp)
	return cha1, cha2
}

/* not used - kept for reference only.
func pairanyThing(out1, out2 chan<- anyThing, inp <-chan anyThing) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
} */

func pairanyThing(out1, out2 chan<- anyThing, inp <-chan anyThing) {
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

// End of anyThingPair functions
// ===========================================================================

// ===========================================================================
// Beg of anyThingFork functions

// anyThingFork returns two channels
// either of which is to receive
// every result of inp
// before close.
func anyThingFork(inp <-chan anyThing) (out1, out2 <-chan anyThing) {
	cha1 := make(chan anyThing)
	cha2 := make(chan anyThing)
	go forkanyThing(cha1, cha2, inp)
	return cha1, cha2
}

/* not used - kept for reference only.
func forkanyThing(out1, out2 chan<- anyThing, inp <-chan anyThing) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
} */

func forkanyThing(out1, out2 chan<- anyThing, inp <-chan anyThing) {
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

// End of anyThingFork functions
// ===========================================================================

// ===========================================================================
// Beg of anyThingFanIn2 simple binary Fan-In

// anyThingFanIn2 returns a channel to receive
// all from both `inp` and `inp2`
// before close.
func anyThingFanIn2(inp, inp2 <-chan anyThing) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go fanIn2anyThing(cha, inp, inp2)
	return cha
}

/* not used - kept for reference only.
// fanin2anyThing as seen in Go Concurrency Patterns
func fanin2anyThing(out chan<- anyThing, inp, inp2 <-chan anyThing) {
	for {
		select {
		case e := <-inp:
			out <- e
		case e := <-inp2:
			out <- e
		}
	}
} */

func fanIn2anyThing(out chan<- anyThing, inp, inp2 <-chan anyThing) {
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

// End of anyThingFanIn2 simple binary Fan-In
// ===========================================================================
