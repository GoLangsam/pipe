// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate genny -in $GOFILE	-out ../xxs/internal/$GOFILE	gen "anymode=*anySupply"
//go:generate genny -in $GOFILE	-out ../xxl/internal/$GOFILE	gen "anymode=*anyDemand"
//go:generate genny -in $GOFILE	-out ../xxsl/internal/$GOFILE	gen "anymode=anyThingChannel"

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// anymode is the generic channel type connecting the pipe network components.
type anymode generic.Type

// ===========================================================================
// Beg of anyThingChan producers

// anyThingChan returns a channel to receive
// all inputs
// before close.
func anyThingChan(inp ...anyThing) (out anymode) {
	cha := anymodeMakeChan()
	go chananyThing(cha, inp...)
	return cha
}

func chananyThing(out anymode, inp ...anyThing) {
	defer out.Close()
	for i := range inp {
		out.Put(inp[i])
	}
}

// anyThingChanSlice returns a channel to receive
// all inputs
// before close.
func anyThingChanSlice(inp ...[]anyThing) (out anymode) {
	cha := anymodeMakeChan()
	go chananyThingSlice(cha, inp...)
	return cha
}

func chananyThingSlice(out anymode, inp ...[]anyThing) {
	defer out.Close()
	for i := range inp {
		for j := range inp[i] {
			out.Put(inp[i][j])
		}
	}
}

// anyThingChanFuncNok returns a channel to receive
// all results of generator `gen`
// until `!ok`
// before close.
func anyThingChanFuncNok(gen func() (anyThing, bool)) (out anymode) {
	cha := anymodeMakeChan()
	go chananyThingFuncNok(cha, gen)
	return cha
}

func chananyThingFuncNok(out anymode, gen func() (anyThing, bool)) {
	defer out.Close()
	for {
		res, ok := gen() // generate
		if !ok {
			return
		}
		out.Put(res)
	}
}

// anyThingChanFuncErr returns a channel to receive
// all results of generator `gen`
// until `err != nil`
// before close.
func anyThingChanFuncErr(gen func() (anyThing, error)) (out anymode) {
	cha := anymodeMakeChan()
	go chananyThingFuncErr(cha, gen)
	return cha
}

func chananyThingFuncErr(out anymode, gen func() (anyThing, error)) {
	defer out.Close()
	for {
		res, err := gen() // generate
		if err != nil {
			return
		}
		out.Put(res)
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
func anyThingPipe(inp anymode, ops ...func(a anyThing)) (out anymode) {
	cha := anymodeMakeChan()
	go pipeanyThing(cha, inp, ops...)
	return cha
}

func pipeanyThing(out anymode, inp anymode, ops ...func(a anyThing)) {
	defer out.Close()
	for i, ok := inp.Get(); ok; i, ok = inp.Get() {
		for _, op := range ops {
			if op != nil {
				op(i) // chain action
			}
		}
		out.Put(i) // send it
	}
}

// anyThingPipeFunc
// will chain every `act` to every `inp` and
// returns a channel to receive
// each result
// before close.
func anyThingPipeFunc(inp anymode, acts ...func(a anyThing) anyThing) (out anymode) {
	cha := anymodeMakeChan()
	go pipeanyThingFunc(cha, inp, acts...)
	return cha
}

func pipeanyThingFunc(out anymode, inp anymode, acts ...func(a anyThing) anyThing) {
	defer out.Close()
	for i, ok := inp.Get(); ok; i, ok = inp.Get() {
		for _, act := range acts {
			if act != nil {
				i = act(i) // chain action
			}
		}
		out.Put(i) // send result
	}
}

// End of anyThingPipe functions
// ===========================================================================

// ===========================================================================
// Beg of anyThingTube closures around anyThingPipe

// anyThingTube returns a closure around PipeanyThing (_, ops...).
func anyThingTube(ops ...func(a anyThing)) (tube func(inp anymode) (out anymode)) {

	return func(inp anymode) (out anymode) {
		return anyThingPipe(inp, ops...)
	}
}

// anyThingTubeFunc returns a closure around PipeanyThingFunc (_, acts...).
func anyThingTubeFunc(acts ...func(a anyThing) anyThing) (tube func(inp anymode) (out anymode)) {

	return func(inp anymode) (out anymode) {
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
func anyThingDone(inp anymode, ops ...func(a anyThing)) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doneanyThing(sig, inp, ops...)
	return sig
}

func doneanyThing(done chan<- struct{}, inp anymode, ops ...func(a anyThing)) {
	defer close(done)
	for i, ok := inp.Get(); ok; i, ok = inp.Get() {
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
func anyThingDoneFunc(inp anymode, acts ...func(a anyThing) anyThing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doneanyThingFunc(sig, inp, acts...)
	return sig
}

func doneanyThingFunc(done chan<- struct{}, inp anymode, acts ...func(a anyThing) anyThing) {
	defer close(done)
	for i, ok := inp.Get(); ok; i, ok = inp.Get() {
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
func anyThingDoneSlice(inp anymode) (done <-chan []anyThing) {
	sig := make(chan []anyThing)
	go doneanyThingSlice(sig, inp)
	return sig
}

func doneanyThingSlice(done chan<- []anyThing, inp anymode) {
	defer close(done)
	slice := []anyThing{}
	for i, ok := inp.Get(); ok; i, ok = inp.Get() {
		slice = append(slice, i)
	}
	done <- slice
}

// End of anyThingDone terminators
// ===========================================================================

// ===========================================================================
// Beg of anyThingFini closures

// anyThingFini returns a closure around `anyThingDone(_, ops...)`.
func anyThingFini(ops ...func(a anyThing)) func(inp anymode) (done <-chan struct{}) {

	return func(inp anymode) (done <-chan struct{}) {
		return anyThingDone(inp, ops...)
	}
}

// anyThingFiniFunc returns a closure around `anyThingDoneFunc(_, acts...)`.
func anyThingFiniFunc(acts ...func(a anyThing) anyThing) func(inp anymode) (done <-chan struct{}) {

	return func(inp anymode) (done <-chan struct{}) {
		return anyThingDoneFunc(inp, acts...)
	}
}

// anyThingFiniSlice returns a closure around `anyThingDoneSlice(_)`.
func anyThingFiniSlice() func(inp anymode) (done <-chan []anyThing) {

	return func(inp anymode) (done <-chan []anyThing) {
		return anyThingDoneSlice(inp)
	}
}

// End of anyThingFini closures
// ===========================================================================

// ===========================================================================
// Beg of anyThingPair functions

// anyThingPair returns a pair of channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func anyThingPair(inp anymode) (out1, out2 anymode) {
	cha1 := anymodeMakeChan()
	cha2 := anymodeMakeChan()
	go pairanyThing(cha1, cha2, inp)
	return cha1, cha2
}

func pairanyThing(out1, out2 anymode, inp anymode) {
	defer out1.Close()
	defer out2.Close()
	for i, ok := inp.Get(); ok; i, ok = inp.Get() {
		out1.Put(i)
		out2.Put(i)
	}
}

// End of anyThingPair functions
// ===========================================================================
