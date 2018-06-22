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
// Beg of AnySupply channel object

// AnySupply is a
// supply channel
type AnySupply struct {
	dat chan anyThing
	//  chan struct{}
}

// MakeAnySupplyChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeAnySupplyChan() *AnySupply {
	d := AnySupply{
		dat: make(chan anyThing),
		// : make(chan struct{}),
	}
	return &d
}

// MakeAnySupplyBuff returns
// a (pointer to a) fresh
// buffered (with capacity=`cap`)
// supply channel
func MakeAnySupplyBuff(cap int) *AnySupply {
	d := AnySupply{
		dat: make(chan anyThing, cap),
		// : make(chan struct{}),
	}
	return &d
}

// Provide is the send method
// - aka "myAnyChan <- myAny"
func (c *AnySupply) Provide(dat anyThing) {
	// .req
	c.dat <- dat
}

// Receive is the receive operator as method
// - aka "myAny := <-myAnyChan"
func (c *AnySupply) Receive() (dat anyThing) {
	// eq <- struct{}{}
	return <-c.dat
}

// Request is the comma-ok multi-valued form of Receive and
// reports whether a received value was sent before the anyThing channel was closed
func (c *AnySupply) Request() (dat anyThing, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// Close closes the underlying anyThing channel
func (c *AnySupply) Close() {
	close(c.dat)
}

// Cap reports the capacity of the underlying anyThing channel
func (c *AnySupply) Cap() int {
	return cap(c.dat)
}

// Len reports the length of the underlying anyThing channel
func (c *AnySupply) Len() int {
	return len(c.dat)
}

// End of AnySupply channel object
// ===========================================================================

// ===========================================================================
// Beg of anyThingChan producers

// anyThingChan returns a channel to receive
// all inputs
// before close.
func anyThingChan(inp ...anyThing) (out *AnySupply) {
	cha := MakeAnySupplyChan()
	go chananyThing(cha, inp...)
	return cha
}

func chananyThing(out *AnySupply, inp ...anyThing) {
	defer out.Close()
	for i := range inp {
		out.Provide(inp[i])
	}
}

// anyThingChanSlice returns a channel to receive
// all inputs
// before close.
func anyThingChanSlice(inp ...[]anyThing) (out *AnySupply) {
	cha := MakeAnySupplyChan()
	go chananyThingSlice(cha, inp...)
	return cha
}

func chananyThingSlice(out *AnySupply, inp ...[]anyThing) {
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
func anyThingChanFuncNok(gen func() (anyThing, bool)) (out *AnySupply) {
	cha := MakeAnySupplyChan()
	go chananyThingFuncNok(cha, gen)
	return cha
}

func chananyThingFuncNok(out *AnySupply, gen func() (anyThing, bool)) {
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
func anyThingChanFuncErr(gen func() (anyThing, error)) (out *AnySupply) {
	cha := MakeAnySupplyChan()
	go chananyThingFuncErr(cha, gen)
	return cha
}

func chananyThingFuncErr(out *AnySupply, gen func() (anyThing, error)) {
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
func anyThingPipeFunc(inp *AnySupply, act func(a anyThing) anyThing) (out *AnySupply) {
	cha := MakeAnySupplyChan()
	if act == nil {
		act = func(a anyThing) anyThing { return a }
	}
	go pipeanyThingFunc(cha, inp, act)
	return cha
}

func pipeanyThingFunc(out *AnySupply, inp *AnySupply, act func(a anyThing) anyThing) {
	defer out.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out.Provide(act(i))
	}
}

// anyThingPipeBuffer returns a buffered channel with capacity `cap` to receive
// all `inp`
// before close.
func anyThingPipeBuffer(inp *AnySupply, cap int) (out *AnySupply) {
	cha := MakeAnySupplyBuff(cap)
	go pipeanyThingBuffer(cha, inp)
	return cha
}

func pipeanyThingBuffer(out *AnySupply, inp *AnySupply) {
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
func anyThingTubeFunc(act func(a anyThing) anyThing) (tube func(inp *AnySupply) (out *AnySupply)) {

	return func(inp *AnySupply) (out *AnySupply) {
		return anyThingPipeFunc(inp, act)
	}
}

// anyThingTubeBuffer returns a closure around PipeanyThingBuffer (_, cap).
func anyThingTubeBuffer(cap int) (tube func(inp *AnySupply) (out *AnySupply)) {

	return func(inp *AnySupply) (out *AnySupply) {
		return anyThingPipeBuffer(inp, cap)
	}
}

// End of anyThingTube closures
// ===========================================================================

// ===========================================================================
// Beg of anyThingDone terminators

// anyThingDone returns a channel to receive
// one signal before close after `inp` has been drained.
func anyThingDone(inp *AnySupply) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doitanyThing(sig, inp)
	return sig
}

func doitanyThing(done chan<- struct{}, inp *AnySupply) {
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
func anyThingDoneSlice(inp *AnySupply) (done <-chan []anyThing) {
	sig := make(chan []anyThing)
	go doitanyThingSlice(sig, inp)
	return sig
}

func doitanyThingSlice(done chan<- []anyThing, inp *AnySupply) {
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
func anyThingDoneFunc(inp *AnySupply, act func(a anyThing)) (done <-chan struct{}) {
	sig := make(chan struct{})
	if act == nil {
		act = func(a anyThing) { return }
	}
	go doitanyThingFunc(sig, inp, act)
	return sig
}

func doitanyThingFunc(done chan<- struct{}, inp *AnySupply, act func(a anyThing)) {
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
func anyThingFini() func(inp *AnySupply) (done <-chan struct{}) {

	return func(inp *AnySupply) (done <-chan struct{}) {
		return anyThingDone(inp)
	}
}

// anyThingFiniSlice returns a closure around `DoneanyThingSlice(_)`.
func anyThingFiniSlice() func(inp *AnySupply) (done <-chan []anyThing) {

	return func(inp *AnySupply) (done <-chan []anyThing) {
		return anyThingDoneSlice(inp)
	}
}

// anyThingFiniFunc returns a closure around `DoneanyThingFunc(_, act)`.
func anyThingFiniFunc(act func(a anyThing)) func(inp *AnySupply) (done <-chan struct{}) {

	return func(inp *AnySupply) (done <-chan struct{}) {
		return anyThingDoneFunc(inp, act)
	}
}

// End of anyThingFini closures
// ===========================================================================

// ===========================================================================
// Beg of anyThingPair functions

// anyThingPair returns a pair of channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func anyThingPair(inp *AnySupply) (out1, out2 *AnySupply) {
	cha1 := MakeAnySupplyChan()
	cha2 := MakeAnySupplyChan()
	go pairanyThing(cha1, cha2, inp)
	return cha1, cha2
}

func pairanyThing(out1, out2 *AnySupply, inp *AnySupply) {
	defer out1.Close()
	defer out2.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out1.Provide(i)
		out2.Provide(i)
	}
}

// End of anyThingPair functions
// ===========================================================================
