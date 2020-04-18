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
// Beg of anyDemand channel object

// anyDemand is a
// demand channel
type anyDemand struct {
	req chan struct{}
	ch  chan anyThing
}

// anyDemandMakeChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel.
func anyDemandMakeChan() *anyDemand {
	d := anyDemand{
		req: make(chan struct{}),
		ch:  make(chan anyThing),
	}
	return &d
}

// anyDemandMakeBuff returns
// a (pointer to a) fresh
// buffered
// demand channel
// (with capacity=`cap`).
func anyDemandMakeBuff(cap int) *anyDemand {
	d := anyDemand{
		req: make(chan struct{}),
		ch:  make(chan anyThing, cap),
	}
	return &d
}

// ---------------------------------------------------------------------------

// Get is the comma-ok multi-valued form to receive from the channel and
// reports whether a value was received from an open channel
// or not (as it has been closed).
//
// Get blocks until the request is accepted and value `val` has been received from `from`.
func (from *anyDemand) Get() (val anyThing, open bool) {
	from.req <- struct{}{}
	val, open = <-from.ch
	return
}

// ---------------------------------------------------------------------------
// Put or {`defer into.Close()` and Provide }

// Put is the send-upon-request method
// - aka "myAnyChan <- myAny".
//
// Put blocks until requested to send value `val` into `into` and
// reports whether the request channel was open.
//
// Put is a convenience for
//  if Next() { Send(v) } else { Close() }
//
// Put includes housekeeping:
// If `into` has been dropped, `into` is closed.
func (into *anyDemand) Put(val anyThing) (ok bool) {
	ok = into.Next()
	if ok {
		into.ch <- val
	} else {
		into.Close()
	}
	return
}

// Provide is the low-level send-upon-request method
// - aka "myAnyChan <- myAny".
//
// Note: Provide is low-level - its cousin `Put`
// includes housekeeping: `Put`
// closes the channel upon nok.
//
// Hint: Provide is useful in constructors
// together with `defer into.Close()`.
func (into *anyDemand) Provide(val anyThing) (ok bool) {
	ok = into.Next()
	if ok {
		into.ch <- val
	}
	return ok
}

// ---------------------------------------------------------------------------
// Next... => Send

// NextGetFrom `from` for `into` and report success.
//
// Follow it with `into.Send( f(val) )`, if ok.
//
// NextGetFrom includes housekeeping:
// If `into` has been dropped or `from` has been closed,
// `from` is dropped and `into` is closed.
func (into *anyDemand) NextGetFrom(from *anyDemand) (val anyThing, ok bool) {
	ok = into.Next()
	if ok {
		val, ok = from.Get()
	}
	if !ok {
		from.Drop()
		into.Close()
	}
	return
}

// Next is the request method.
// It blocks until a request is received and
// reports whether the request channel was open.
//
// A successful Next is to be followed by one Send(v).
func (into *anyDemand) Next() (ok bool) {
	_, ok = <-into.req
	return
}

// Send is to be used after a successful Next()
func (into *anyDemand) Send(val anyThing) {
	into.ch <- val
}

// ---------------------------------------------------------------------------
// Drop & Close signal requesting / sending as being finished

// Drop is to be called by a consumer when finished requesting.
// The request channel is closed in order to broadcast this.
//
// In order to avoid deadlock, pending sends are drained.
func (from *anyDemand) Drop() {
	close(from.req)
	go func(from *anyDemand) {
		for range from.ch {
		} // drain values - there could be some
	}(from)
}

// Close is to be called by a producer when finished sending.
// The value channel is closed in order to broadcast this.
//
// In order to avoid deadlock, pending requests are drained.
func (into *anyDemand) Close() {
	close(into.ch)
	go func(into *anyDemand) {
		for range into.req {
		} // drain requests - there could be some
	}(into)
}

// ---------------------------------------------------------------------------
// obtain directional handshaking channels
// (for use e.g. in `select` statements)

// From returns the handshaking channels
// (for use e.g. in `select` statements)
// to receive values:
//  `req` to send a request `req <- struct{}{}` and
//  `rcv` to reveive such requested value from.
func (from *anyDemand) From() (req chan<- struct{}, rcv <-chan anyThing) {
	return from.req, from.ch
}

// Into returns the handshaking channels
// (for use e.g. in `select` statements)
// to send values:
//  `req` to receive a request `<-req` and
//  `snd` to send such requested value into.
func (into *anyDemand) Into() (req <-chan struct{}, snd chan<- anyThing) {
	return into.req, into.ch
}

// ---------------------------------------------------------------------------

// New returns a new similar channel.
//
// Useful e.g. when embedded anonymously.
func (c *anyDemand) New() *anyDemand {
	return anyDemandMakeChan()
}

// Self returns itself.
//
// Useful e.g. when embedded anonymously
// and e.g. wrappers for multi-value methods are required.
func (c *anyDemand) Self() *anyDemand {
	return c
}

// Cap reports the capacity of the underlying value channel.
func (c *anyDemand) Cap() int {
	return cap(c.ch)
}

// Len reports the length of the underlying value channel.
func (c *anyDemand) Len() int {
	return len(c.ch)
}

// End of anyDemand channel object
// ===========================================================================

// ===========================================================================
// Beg of anyThingChan producers

// anyThingChan returns a channel to receive
// all inputs
// before close.
func anyThingChan(inp ...anyThing) (out *anyDemand) {
	cha := anyDemandMakeChan()
	go chananyThing(cha, inp...)
	return cha
}

func chananyThing(out *anyDemand, inp ...anyThing) {
	defer out.Close()
	for i := range inp {
		out.Put(inp[i])
	}
}

// anyThingChanSlice returns a channel to receive
// all inputs
// before close.
func anyThingChanSlice(inp ...[]anyThing) (out *anyDemand) {
	cha := anyDemandMakeChan()
	go chananyThingSlice(cha, inp...)
	return cha
}

func chananyThingSlice(out *anyDemand, inp ...[]anyThing) {
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
func anyThingChanFuncNok(gen func() (anyThing, bool)) (out *anyDemand) {
	cha := anyDemandMakeChan()
	go chananyThingFuncNok(cha, gen)
	return cha
}

func chananyThingFuncNok(out *anyDemand, gen func() (anyThing, bool)) {
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
func anyThingChanFuncErr(gen func() (anyThing, error)) (out *anyDemand) {
	cha := anyDemandMakeChan()
	go chananyThingFuncErr(cha, gen)
	return cha
}

func chananyThingFuncErr(out *anyDemand, gen func() (anyThing, error)) {
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
func anyThingPipe(inp *anyDemand, ops ...func(a anyThing)) (out *anyDemand) {
	cha := anyDemandMakeChan()
	go pipeanyThing(cha, inp, ops...)
	return cha
}

func pipeanyThing(out *anyDemand, inp *anyDemand, ops ...func(a anyThing)) {
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
func anyThingPipeFunc(inp *anyDemand, acts ...func(a anyThing) anyThing) (out *anyDemand) {
	cha := anyDemandMakeChan()
	go pipeanyThingFunc(cha, inp, acts...)
	return cha
}

func pipeanyThingFunc(out *anyDemand, inp *anyDemand, acts ...func(a anyThing) anyThing) {
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
func anyThingTube(ops ...func(a anyThing)) (tube func(inp *anyDemand) (out *anyDemand)) {

	return func(inp *anyDemand) (out *anyDemand) {
		return anyThingPipe(inp, ops...)
	}
}

// anyThingTubeFunc returns a closure around PipeanyThingFunc (_, acts...).
func anyThingTubeFunc(acts ...func(a anyThing) anyThing) (tube func(inp *anyDemand) (out *anyDemand)) {

	return func(inp *anyDemand) (out *anyDemand) {
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
func anyThingDone(inp *anyDemand, ops ...func(a anyThing)) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doneanyThing(sig, inp, ops...)
	return sig
}

func doneanyThing(done chan<- struct{}, inp *anyDemand, ops ...func(a anyThing)) {
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
func anyThingDoneFunc(inp *anyDemand, acts ...func(a anyThing) anyThing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doneanyThingFunc(sig, inp, acts...)
	return sig
}

func doneanyThingFunc(done chan<- struct{}, inp *anyDemand, acts ...func(a anyThing) anyThing) {
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
func anyThingDoneSlice(inp *anyDemand) (done <-chan []anyThing) {
	sig := make(chan []anyThing)
	go doneanyThingSlice(sig, inp)
	return sig
}

func doneanyThingSlice(done chan<- []anyThing, inp *anyDemand) {
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
func anyThingFini(ops ...func(a anyThing)) func(inp *anyDemand) (done <-chan struct{}) {

	return func(inp *anyDemand) (done <-chan struct{}) {
		return anyThingDone(inp, ops...)
	}
}

// anyThingFiniFunc returns a closure around `anyThingDoneFunc(_, acts...)`.
func anyThingFiniFunc(acts ...func(a anyThing) anyThing) func(inp *anyDemand) (done <-chan struct{}) {

	return func(inp *anyDemand) (done <-chan struct{}) {
		return anyThingDoneFunc(inp, acts...)
	}
}

// anyThingFiniSlice returns a closure around `anyThingDoneSlice(_)`.
func anyThingFiniSlice() func(inp *anyDemand) (done <-chan []anyThing) {

	return func(inp *anyDemand) (done <-chan []anyThing) {
		return anyThingDoneSlice(inp)
	}
}

// End of anyThingFini closures
// ===========================================================================

// ===========================================================================
// Beg of anyThingPair functions

// anyThingPair returns a pair of channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func anyThingPair(inp *anyDemand) (out1, out2 *anyDemand) {
	cha1 := anyDemandMakeChan()
	cha2 := anyDemandMakeChan()
	go pairanyThing(cha1, cha2, inp)
	return cha1, cha2
}

func pairanyThing(out1, out2 *anyDemand, inp *anyDemand) {
	defer out1.Close()
	defer out2.Close()
	for i, ok := inp.Get(); ok; i, ok = inp.Get() {
		out1.Put(i)
		out2.Put(i)
	}
}

// End of anyThingPair functions
// ===========================================================================
