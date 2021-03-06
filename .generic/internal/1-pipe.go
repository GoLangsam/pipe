// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by golang.org/x/tools/cmd/bundle. DO NOT EDIT.

package pipe

// ===========================================================================
// Beg of ThingMake creators

// ThingMakeChan returns a new open channel
// (simply a 'chan Thing' that is).
//
// Note: No 'Thing-producer' is launched here yet! (as is in all the other functions).
//  This is useful to easily create corresponding variables such as:
//
// var myThingPipelineStartsHere := ThingMakeChan()
// // ... lot's of code to design and build Your favourite "myThingWorkflowPipeline"
// 	// ...
// 	// ... *before* You start pouring data into it, e.g. simply via:
// 	for drop := range water {
// myThingPipelineStartsHere <- drop
// 	}
// close(myThingPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for ThingPipeBuffer) the channel is unbuffered.
//
func ThingMakeChan() (out chan Thing) {
	return make(chan Thing)
}

// End of ThingMake creators
// ===========================================================================

// ===========================================================================
// Beg of ThingChan producers

// ThingChan returns a channel to receive
// all inputs
// before close.
func ThingChan(inp ...Thing) (out <-chan Thing) {
	cha := make(chan Thing)
	go chanThing(cha, inp...)
	return cha
}

func chanThing(out chan<- Thing, inp ...Thing) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ThingChanSlice returns a channel to receive
// all inputs
// before close.
func ThingChanSlice(inp ...[]Thing) (out <-chan Thing) {
	cha := make(chan Thing)
	go chanThingSlice(cha, inp...)
	return cha
}

func chanThingSlice(out chan<- Thing, inp ...[]Thing) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ThingChanFuncNok returns a channel to receive
// all results of generator `gen`
// until `!ok`
// before close.
func ThingChanFuncNok(gen func() (Thing, bool)) (out <-chan Thing) {
	cha := make(chan Thing)
	go chanThingFuncNok(cha, gen)
	return cha
}

func chanThingFuncNok(out chan<- Thing, gen func() (Thing, bool)) {
	defer close(out)
	for {
		res, ok := gen() // generate
		if !ok {
			return
		}
		out <- res
	}
}

// ThingChanFuncErr returns a channel to receive
// all results of generator `gen`
// until `err != nil`
// before close.
func ThingChanFuncErr(gen func() (Thing, error)) (out <-chan Thing) {
	cha := make(chan Thing)
	go chanThingFuncErr(cha, gen)
	return cha
}

func chanThingFuncErr(out chan<- Thing, gen func() (Thing, error)) {
	defer close(out)
	for {
		res, err := gen() // generate
		if err != nil {
			return
		}
		out <- res
	}
}

// End of ThingChan producers
// ===========================================================================

// ===========================================================================
// Beg of ThingPipe functions

// ThingPipe
// will apply every `op` to every `inp` and
// returns a channel to receive
// each `inp`
// before close.
//
// Note: For functional people,
// this 'could' be named `ThingMap`.
// Just: 'map' has a very different meaning in go lang.
func ThingPipe(inp <-chan Thing, ops ...func(a Thing)) (out <-chan Thing) {
	cha := make(chan Thing)
	go pipeThing(cha, inp, ops...)
	return cha
}

func pipeThing(out chan<- Thing, inp <-chan Thing, ops ...func(a Thing)) {
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

// ThingPipeFunc
// will chain every `act` to every `inp` and
// returns a channel to receive
// each result
// before close.
func ThingPipeFunc(inp <-chan Thing, acts ...func(a Thing) Thing) (out <-chan Thing) {
	cha := make(chan Thing)
	go pipeThingFunc(cha, inp, acts...)
	return cha
}

func pipeThingFunc(out chan<- Thing, inp <-chan Thing, acts ...func(a Thing) Thing) {
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

// End of ThingPipe functions
// ===========================================================================

// ===========================================================================
// Beg of ThingTube closures around ThingPipe

// ThingTube returns a closure around PipeThing (_, ops...).
func ThingTube(ops ...func(a Thing)) (tube func(inp <-chan Thing) (out <-chan Thing)) {

	return func(inp <-chan Thing) (out <-chan Thing) {
		return ThingPipe(inp, ops...)
	}
}

// ThingTubeFunc returns a closure around PipeThingFunc (_, acts...).
func ThingTubeFunc(acts ...func(a Thing) Thing) (tube func(inp <-chan Thing) (out <-chan Thing)) {

	return func(inp <-chan Thing) (out <-chan Thing) {
		return ThingPipeFunc(inp, acts...)
	}
}

// End of ThingTube closures around ThingPipe
// ===========================================================================

// ===========================================================================
// Beg of ThingDone terminators

// ThingDone
// will apply every `op` to every `inp` and
// returns a channel to receive
// one signal
// upon close.
func ThingDone(inp <-chan Thing, ops ...func(a Thing)) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doneThing(sig, inp, ops...)
	return sig
}

func doneThing(done chan<- struct{}, inp <-chan Thing, ops ...func(a Thing)) {
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

// ThingDoneFunc
// will chain every `act` to every `inp` and
// returns a channel to receive
// one signal
// upon close.
func ThingDoneFunc(inp <-chan Thing, acts ...func(a Thing) Thing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doneThingFunc(sig, inp, acts...)
	return sig
}

func doneThingFunc(done chan<- struct{}, inp <-chan Thing, acts ...func(a Thing) Thing) {
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

// ThingDoneSlice returns a channel to receive
// a slice with every Thing received on `inp`
// upon close.
//
// Note: Unlike ThingDone, ThingDoneSlice sends the fully accumulated slice, not just an event, once upon close of inp.
func ThingDoneSlice(inp <-chan Thing) (done <-chan []Thing) {
	sig := make(chan []Thing)
	go doneThingSlice(sig, inp)
	return sig
}

func doneThingSlice(done chan<- []Thing, inp <-chan Thing) {
	defer close(done)
	slice := []Thing{}
	for i := range inp {
		slice = append(slice, i)
	}
	done <- slice
}

// End of ThingDone terminators
// ===========================================================================

// ===========================================================================
// Beg of ThingFini closures

// ThingFini returns a closure around `ThingDone(_, ops...)`.
func ThingFini(ops ...func(a Thing)) func(inp <-chan Thing) (done <-chan struct{}) {

	return func(inp <-chan Thing) (done <-chan struct{}) {
		return ThingDone(inp, ops...)
	}
}

// ThingFiniFunc returns a closure around `ThingDoneFunc(_, acts...)`.
func ThingFiniFunc(acts ...func(a Thing) Thing) func(inp <-chan Thing) (done <-chan struct{}) {

	return func(inp <-chan Thing) (done <-chan struct{}) {
		return ThingDoneFunc(inp, acts...)
	}
}

// ThingFiniSlice returns a closure around `ThingDoneSlice(_)`.
func ThingFiniSlice() func(inp <-chan Thing) (done <-chan []Thing) {

	return func(inp <-chan Thing) (done <-chan []Thing) {
		return ThingDoneSlice(inp)
	}
}

// End of ThingFini closures
// ===========================================================================

// ===========================================================================
// Beg of ThingPair functions

// ThingPair returns a pair of channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func ThingPair(inp <-chan Thing) (out1, out2 <-chan Thing) {
	cha1 := make(chan Thing)
	cha2 := make(chan Thing)
	go pairThing(cha1, cha2, inp)
	return cha1, cha2
}

/* not used - kept for reference only.
func pairThing ( out1 , out2 chan <- Thing , inp <- chan Thing ) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
} */

func pairThing(out1, out2 chan<- Thing, inp <-chan Thing) {
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

// End of ThingPair functions
// ===========================================================================

// ===========================================================================
// Beg of ThingFork functions

// ThingFork returns two channels
// either of which is to receive
// every result of inp
// before close.
func ThingFork(inp <-chan Thing) (out1, out2 <-chan Thing) {
	cha1 := make(chan Thing)
	cha2 := make(chan Thing)
	go forkThing(cha1, cha2, inp)
	return cha1, cha2
}

/* not used - kept for reference only.
func forkThing ( out1 , out2 chan <- Thing , inp <- chan Thing ) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
} */

func forkThing(out1, out2 chan<- Thing, inp <-chan Thing) {
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

// End of ThingFork functions
// ===========================================================================

// ===========================================================================
// Beg of ThingFanIn2 simple binary Fan-In

// ThingFanIn2 returns a channel to receive
// all from both `inp` and `inp2`
// before close.
func ThingFanIn2(inp, inp2 <-chan Thing) (out <-chan Thing) {
	cha := make(chan Thing)
	go fanIn2Thing(cha, inp, inp2)
	return cha
}

/* not used - kept for reference only.
// fanin2Thing as seen in Go Concurrency Patterns
func fanin2Thing ( out chan <- Thing , inp , inp2 <- chan Thing ) {
	for {
		select {
		case e := <-inp:
			out <- e
		case e := <-inp2:
			out <- e
		}
	}
} */

func fanIn2Thing(out chan<- Thing, inp, inp2 <-chan Thing) {
	defer close(out)

	var (
		closed bool  // we found a chan closed
		ok     bool  // did we read successfully?
		e      Thing // what we've read
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

// End of ThingFanIn2 simple binary Fan-In
// ===========================================================================
