// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingChan producers

// anyThingChan returns a channel to receive
// all inputs
// before close.
func anyThingChan(inp ...anyThing) (out anyThingChannel) {
	cha := MakeAnyChannelChan()
	go chananyThing(cha, inp...)
	return cha
}

func chananyThing(out anyThingChannel, inp ...anyThing) {
	defer out.Close()
	for i := range inp {
		out.Provide(inp[i])
	}
}

// anyThingChanSlice returns a channel to receive
// all inputs
// before close.
func anyThingChanSlice(inp ...[]anyThing) (out anyThingChannel) {
	cha := MakeAnyChannelChan()
	go chananyThingSlice(cha, inp...)
	return cha
}

func chananyThingSlice(out anyThingChannel, inp ...[]anyThing) {
	defer out.Close()
	for i := range inp {
		for j := range inp[i] {
			out.Provide(inp[i][j])
		}
	}
}

// anyThingChanFuncNok returns a channel to receive
// all results of generator `gen`
// until `!ok`
// before close.
func anyThingChanFuncNok(gen func() (anyThing, bool)) (out anyThingChannel) {
	cha := MakeAnyChannelChan()
	go chananyThingFuncNok(cha, gen)
	return cha
}

func chananyThingFuncNok(out anyThingChannel, gen func() (anyThing, bool)) {
	defer out.Close()
	for {
		res, ok := gen() // generate
		if !ok {
			return
		}
		out.Provide(res)
	}
}

// anyThingChanFuncErr returns a channel to receive
// all results of generator `gen`
// until `err != nil`
// before close.
func anyThingChanFuncErr(gen func() (anyThing, error)) (out anyThingChannel) {
	cha := MakeAnyChannelChan()
	go chananyThingFuncErr(cha, gen)
	return cha
}

func chananyThingFuncErr(out anyThingChannel, gen func() (anyThing, error)) {
	defer out.Close()
	for {
		res, err := gen() // generate
		if err != nil {
			return
		}
		out.Provide(res)
	}
}

// End of anyThingChan producers
// ===========================================================================

// ===========================================================================
// Beg of anyThingPipe functions

// anyThingPipeFunc returns a channel to receive
// every result of action `act` applied to `inp`
// before close.
// Note: it 'could' be anyThingPipeMap for functional people,
// but 'map' has a very different meaning in go lang.
func anyThingPipeFunc(inp anyThingChannel, act func(a anyThing) anyThing) (out anyThingChannel) {
	cha := MakeAnyChannelChan()
	if act == nil {
		act = func(a anyThing) anyThing { return a }
	}
	go pipeanyThingFunc(cha, inp, act)
	return cha
}

func pipeanyThingFunc(out anyThingChannel, inp anyThingChannel, act func(a anyThing) anyThing) {
	defer out.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out.Provide(act(i))
	}
}

// anyThingPipeBuffer returns a buffered channel with capacity `cap` to receive
// all `inp`
// before close.
func anyThingPipeBuffer(inp anyThingChannel, cap int) (out anyThingChannel) {
	cha := MakeAnyChannelBuff(cap)
	go pipeanyThingBuffer(cha, inp)
	return cha
}

func pipeanyThingBuffer(out anyThingChannel, inp anyThingChannel) {
	defer out.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out.Provide(i)
	}
}

// End of anyThingPipe functions
// ===========================================================================

// ===========================================================================
// Beg of anyThingTube closures

// anyThingTubeFunc returns a closure around PipeanyThingFunc (_, act).
func anyThingTubeFunc(act func(a anyThing) anyThing) (tube func(inp anyThingChannel) (out anyThingChannel)) {

	return func(inp anyThingChannel) (out anyThingChannel) {
		return anyThingPipeFunc(inp, act)
	}
}

// anyThingTubeBuffer returns a closure around PipeanyThingBuffer (_, cap).
func anyThingTubeBuffer(cap int) (tube func(inp anyThingChannel) (out anyThingChannel)) {

	return func(inp anyThingChannel) (out anyThingChannel) {
		return anyThingPipeBuffer(inp, cap)
	}
}

// End of anyThingTube closures
// ===========================================================================

// ===========================================================================
// Beg of anyThingDone terminators

// anyThingDone returns a channel to receive
// one signal before close after `inp` has been drained.
func anyThingDone(inp anyThingChannel) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doitanyThing(sig, inp)
	return sig
}

func doitanyThing(done chan<- struct{}, inp anyThingChannel) {
	defer close(done)
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// anyThingDoneSlice returns a channel to receive
// a slice with every anyThing received on `inp`
// before close.
//
//  Note: Unlike anyThingDone, anyThingDoneSlice sends the fully accumulated slice, not just an event, once upon close of inp.
func anyThingDoneSlice(inp anyThingChannel) (done <-chan []anyThing) {
	sig := make(chan []anyThing)
	go doitanyThingSlice(sig, inp)
	return sig
}

func doitanyThingSlice(done chan<- []anyThing, inp anyThingChannel) {
	defer close(done)
	slice := []anyThing{}
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		slice = append(slice, i)
	}
	done <- slice
}

// anyThingDoneFunc returns a channel to receive
// one signal after `act` has been applied to every `inp`
// before close.
func anyThingDoneFunc(inp anyThingChannel, act func(a anyThing)) (done <-chan struct{}) {
	sig := make(chan struct{})
	if act == nil {
		act = func(a anyThing) { return }
	}
	go doitanyThingFunc(sig, inp, act)
	return sig
}

func doitanyThingFunc(done chan<- struct{}, inp anyThingChannel, act func(a anyThing)) {
	defer close(done)
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		act(i) // apply action
	}
	done <- struct{}{}
}

// End of anyThingDone terminators
// ===========================================================================

// ===========================================================================
// Beg of anyThingFini closures

// anyThingFini returns a closure around `DoneanyThing(_)`.
func anyThingFini() func(inp anyThingChannel) (done <-chan struct{}) {

	return func(inp anyThingChannel) (done <-chan struct{}) {
		return anyThingDone(inp)
	}
}

// anyThingFiniSlice returns a closure around `DoneanyThingSlice(_)`.
func anyThingFiniSlice() func(inp anyThingChannel) (done <-chan []anyThing) {

	return func(inp anyThingChannel) (done <-chan []anyThing) {
		return anyThingDoneSlice(inp)
	}
}

// anyThingFiniFunc returns a closure around `DoneanyThingFunc(_, act)`.
func anyThingFiniFunc(act func(a anyThing)) func(inp anyThingChannel) (done <-chan struct{}) {

	return func(inp anyThingChannel) (done <-chan struct{}) {
		return anyThingDoneFunc(inp, act)
	}
}

// End of anyThingFini closures
// ===========================================================================

// ===========================================================================
// Beg of anyThingPair functions

// anyThingPair returns a pair of channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func anyThingPair(inp anyThingChannel) (out1, out2 anyThingChannel) {
	cha1 := MakeAnyChannelChan()
	cha2 := MakeAnyChannelChan()
	go pairanyThing(cha1, cha2, inp)
	return cha1, cha2
}

func pairanyThing(out1, out2 anyThingChannel, inp anyThingChannel) {
	defer out1.Close()
	defer out2.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out1.Provide(i)
		out2.Provide(i)
	}
}

// End of anyThingPair functions
// ===========================================================================
