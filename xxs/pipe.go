// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by golang.org/x/tools/cmd/bundle. DO NOT EDIT.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

type Any generic.Type

// ===========================================================================
// Beg of AnySupply channel object

// AnySupply is a
// supply channel
type AnySupply struct {
	dat chan Any
	//  chan struct{}
}

// MakeAnySupplyChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeAnySupplyChan() *AnySupply {
	d := AnySupply{
		dat: make(chan Any),
		//	req: make(chan struct{}),
	}
	return &d
}

// MakeAnySupplyBuff returns
// a (pointer to a) fresh
// buffered (with capacity=`cap`)
// supply channel
func MakeAnySupplyBuff(cap int) *AnySupply {
	d := AnySupply{
		dat: make(chan Any, cap),
		//	req: make(chan struct{}),
	}
	return &d
}

// Provide is the send method
// - aka "myAnyChan <- myAny"
func (c *AnySupply) Provide(dat Any) {
	// .req
	c.dat <- dat
}

// Receive is the receive operator as method
// - aka "myAny := <-myAnyChan"
func (c *AnySupply) Receive() (dat Any) {
	// eq <- struct{}{}
	return <-c.dat
}

// Request is the comma-ok multi-valued form of Receive and
// reports whether a received value was sent before the Any channel was closed
func (c *AnySupply) Request() (dat Any, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// Close closes the underlying Any channel
func (c *AnySupply) Close() {
	close(c.dat)
}

// Cap reports the capacity of the underlying Any channel
func (c *AnySupply) Cap() int {
	return cap(c.dat)
}

// Len reports the length of the underlying Any channel
func (c *AnySupply) Len() int {
	return len(c.dat)
}

// End of AnySupply channel object
// ===========================================================================

// ===========================================================================
// Beg of ChanAny producers

// ChanAny returns a channel to receive
// all inputs
// before close.
func ChanAny(inp ...Any) (out *AnySupply) {
	cha := MakeAnySupplyChan()
	go chanAny(cha, inp...)
	return cha
}

func chanAny(out *AnySupply, inp ...Any) {
	defer out.Close()
	for i := range inp {
		out.Provide(inp[i])
	}
}

// ChanAnySlice returns a channel to receive
// all inputs
// before close.
func ChanAnySlice(inp ...[]Any) (out *AnySupply) {
	cha := MakeAnySupplyChan()
	go chanAnySlice(cha, inp...)
	return cha
}

func chanAnySlice(out *AnySupply, inp ...[]Any) {
	defer out.Close()
	for i := range inp {
		for j := range inp[i] {
			out.Provide(inp[i][j])
		}
	}
}

// ChanAnyFuncNok returns a channel to receive
// all results of generator `gen`
// until `!ok`
// before close.
func ChanAnyFuncNok(gen func() (Any, bool)) (out *AnySupply) {
	cha := MakeAnySupplyChan()
	go chanAnyFuncNok(cha, gen)
	return cha
}

func chanAnyFuncNok(out *AnySupply, gen func() (Any, bool)) {
	defer out.Close()
	for {
		res, ok := gen() // generate
		if !ok {
			return
		}
		out.Provide(res)
	}
}

// ChanAnyFuncErr returns a channel to receive
// all results of generator `gen`
// until `err != nil`
// before close.
func ChanAnyFuncErr(gen func() (Any, error)) (out *AnySupply) {
	cha := MakeAnySupplyChan()
	go chanAnyFuncErr(cha, gen)
	return cha
}

func chanAnyFuncErr(out *AnySupply, gen func() (Any, error)) {
	defer out.Close()
	for {
		res, err := gen() // generate
		if err != nil {
			return
		}
		out.Provide(res)
	}
}

// End of ChanAny producers
// ===========================================================================

// ===========================================================================
// Beg of PipeAny functions

// PipeAnyFunc returns a channel to receive
// every result of action `act` applied to `inp`
// before close.
// Note: it 'could' be PipeAnyMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeAnyFunc(inp *AnySupply, act func(a Any) Any) (out *AnySupply) {
	cha := MakeAnySupplyChan()
	if act == nil {
		act = func(a Any) Any { return a }
	}
	go pipeAnyFunc(cha, inp, act)
	return cha
}

func pipeAnyFunc(out *AnySupply, inp *AnySupply, act func(a Any) Any) {
	defer out.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out.Provide(act(i))
	}
}

// PipeAnyBuffer returns a buffered channel with capacity `cap` to receive
// all `inp`
// before close.
func PipeAnyBuffer(inp *AnySupply, cap int) (out *AnySupply) {
	cha := MakeAnySupplyBuff(cap)
	go pipeAnyBuffer(cha, inp)
	return cha
}

func pipeAnyBuffer(out *AnySupply, inp *AnySupply) {
	defer out.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out.Provide(i)
	}
}

// End of PipeAny functions
// ===========================================================================

// ===========================================================================
// Beg of TubeAny closures

// TubeAnyFunc returns a closure around PipeAnyFunc (_, act).
func TubeAnyFunc(act func(a Any) Any) (tube func(inp *AnySupply) (out *AnySupply)) {

	return func(inp *AnySupply) (out *AnySupply) {
		return PipeAnyFunc(inp, act)
	}
}

// TubeAnyBuffer returns a closure around PipeAnyBuffer (_, cap).
func TubeAnyBuffer(cap int) (tube func(inp *AnySupply) (out *AnySupply)) {

	return func(inp *AnySupply) (out *AnySupply) {
		return PipeAnyBuffer(inp, cap)
	}
}

// End of TubeAny closures
// ===========================================================================

// ===========================================================================
// Beg of DoneAny terminators

// DoneAny returns a channel to receive
// one signal before close after `inp` has been drained.
func DoneAny(inp *AnySupply) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doitAny(sig, inp)
	return sig
}

func doitAny(done chan<- struct{}, inp *AnySupply) {
	defer close(done)
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneAnySlice returns a channel to receive
// a slice with every Any received on `inp`
// before close.
//
//  Note: Unlike DoneAny, DoneAnySlice sends the fully accumulated slice, not just an event, once upon close of inp.
func DoneAnySlice(inp *AnySupply) (done <-chan []Any) {
	sig := make(chan []Any)
	go doitAnySlice(sig, inp)
	return sig
}

func doitAnySlice(done chan<- []Any, inp *AnySupply) {
	defer close(done)
	slice := []Any{}
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		slice = append(slice, i)
	}
	done <- slice
}

// DoneAnyFunc returns a channel to receive
// one signal after `act` has been applied to every `inp`
// before close.
func DoneAnyFunc(inp *AnySupply, act func(a Any)) (done <-chan struct{}) {
	sig := make(chan struct{})
	if act == nil {
		act = func(a Any) { return }
	}
	go doitAnyFunc(sig, inp, act)
	return sig
}

func doitAnyFunc(done chan<- struct{}, inp *AnySupply, act func(a Any)) {
	defer close(done)
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		act(i) // apply action
	}
	done <- struct{}{}
}

// End of DoneAny terminators
// ===========================================================================

// ===========================================================================
// Beg of FiniAny closures

// FiniAny returns a closure around `DoneAny(_)`.
func FiniAny() func(inp *AnySupply) (done <-chan struct{}) {

	return func(inp *AnySupply) (done <-chan struct{}) {
		return DoneAny(inp)
	}
}

// FiniAnySlice returns a closure around `DoneAnySlice(_)`.
func FiniAnySlice() func(inp *AnySupply) (done <-chan []Any) {

	return func(inp *AnySupply) (done <-chan []Any) {
		return DoneAnySlice(inp)
	}
}

// FiniAnyFunc returns a closure around `DoneAnyFunc(_, act)`.
func FiniAnyFunc(act func(a Any)) func(inp *AnySupply) (done <-chan struct{}) {

	return func(inp *AnySupply) (done <-chan struct{}) {
		return DoneAnyFunc(inp, act)
	}
}

// End of FiniAny closures
// ===========================================================================

// ===========================================================================
// Beg of PairAny functions

// PairAny returns a pair of channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PairAny(inp *AnySupply) (out1, out2 *AnySupply) {
	cha1 := MakeAnySupplyChan()
	cha2 := MakeAnySupplyChan()
	go pairAny(cha1, cha2, inp)
	return cha1, cha2
}

func pairAny(out1, out2 *AnySupply, inp *AnySupply) {
	defer out1.Close()
	defer out2.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out1.Provide(i)
		out2.Provide(i)
	}
}

// End of PairAny functions
// ===========================================================================
