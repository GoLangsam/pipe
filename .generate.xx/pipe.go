// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate genny -in $GOFILE	-out ../xxs/internal/$GOFILE	gen "Anymode=*AnySupply"
//go:generate genny -in $GOFILE	-out ../xxl/internal/$GOFILE	gen "Anymode=*AnyDemand"
//go:generate genny -in $GOFILE	-out ../xxsl/internal/$GOFILE	gen "Anymode=AnyChannel"

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// Anymode is the generic channel type connecting the pipe network components.
type Anymode generic.Type

// ===========================================================================
// Beg of ChanAny producers

// ChanAny returns a channel to receive
// all inputs
// before close.
func ChanAny(inp ...Any) (out Anymode) {
	cha := MakeAnymodeChan()
	go chanAny(cha, inp...)
	return cha
}

func chanAny(out Anymode, inp ...Any) {
	defer out.Close()
	for i := range inp {
		out.Provide(inp[i])
	}
}

// ChanAnySlice returns a channel to receive
// all inputs
// before close.
func ChanAnySlice(inp ...[]Any) (out Anymode) {
	cha := MakeAnymodeChan()
	go chanAnySlice(cha, inp...)
	return cha
}

func chanAnySlice(out Anymode, inp ...[]Any) {
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
func ChanAnyFuncNok(gen func() (Any, bool)) (out Anymode) {
	cha := MakeAnymodeChan()
	go chanAnyFuncNok(cha, gen)
	return cha
}

func chanAnyFuncNok(out Anymode, gen func() (Any, bool)) {
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
func ChanAnyFuncErr(gen func() (Any, error)) (out Anymode) {
	cha := MakeAnymodeChan()
	go chanAnyFuncErr(cha, gen)
	return cha
}

func chanAnyFuncErr(out Anymode, gen func() (Any, error)) {
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
func PipeAnyFunc(inp Anymode, act func(a Any) Any) (out Anymode) {
	cha := MakeAnymodeChan()
	if act == nil {
		act = func(a Any) Any { return a }
	}
	go pipeAnyFunc(cha, inp, act)
	return cha
}

func pipeAnyFunc(out Anymode, inp Anymode, act func(a Any) Any) {
	defer out.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out.Provide(act(i))
	}
}

// PipeAnyBuffer returns a buffered channel with capacity `cap` to receive
// all `inp`
// before close.
func PipeAnyBuffer(inp Anymode, cap int) (out Anymode) {
	cha := MakeAnymodeBuff(cap)
	go pipeAnyBuffer(cha, inp)
	return cha
}

func pipeAnyBuffer(out Anymode, inp Anymode) {
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
func TubeAnyFunc(act func(a Any) Any) (tube func(inp Anymode) (out Anymode)) {

	return func(inp Anymode) (out Anymode) {
		return PipeAnyFunc(inp, act)
	}
}

// TubeAnyBuffer returns a closure around PipeAnyBuffer (_, cap).
func TubeAnyBuffer(cap int) (tube func(inp Anymode) (out Anymode)) {

	return func(inp Anymode) (out Anymode) {
		return PipeAnyBuffer(inp, cap)
	}
}

// End of TubeAny closures
// ===========================================================================

// ===========================================================================
// Beg of DoneAny terminators

// DoneAny returns a channel to receive
// one signal before close after `inp` has been drained.
func DoneAny(inp Anymode) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doitAny(sig, inp)
	return sig
}

func doitAny(done chan<- struct{}, inp Anymode) {
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
func DoneAnySlice(inp Anymode) (done <-chan []Any) {
	sig := make(chan []Any)
	go doitAnySlice(sig, inp)
	return sig
}

func doitAnySlice(done chan<- []Any, inp Anymode) {
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
func DoneAnyFunc(inp Anymode, act func(a Any)) (done <-chan struct{}) {
	sig := make(chan struct{})
	if act == nil {
		act = func(a Any) { return }
	}
	go doitAnyFunc(sig, inp, act)
	return sig
}

func doitAnyFunc(done chan<- struct{}, inp Anymode, act func(a Any)) {
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
func FiniAny() func(inp Anymode) (done <-chan struct{}) {

	return func(inp Anymode) (done <-chan struct{}) {
		return DoneAny(inp)
	}
}

// FiniAnySlice returns a closure around `DoneAnySlice(_)`.
func FiniAnySlice() func(inp Anymode) (done <-chan []Any) {

	return func(inp Anymode) (done <-chan []Any) {
		return DoneAnySlice(inp)
	}
}

// FiniAnyFunc returns a closure around `DoneAnyFunc(_, act)`.
func FiniAnyFunc(act func(a Any)) func(inp Anymode) (done <-chan struct{}) {

	return func(inp Anymode) (done <-chan struct{}) {
		return DoneAnyFunc(inp, act)
	}
}

// End of FiniAny closures
// ===========================================================================

// ===========================================================================
// Beg of PairAny functions

// PairAny returns a pair of channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PairAny(inp Anymode) (out1, out2 Anymode) {
	cha1 := MakeAnymodeChan()
	cha2 := MakeAnymodeChan()
	go pairAny(cha1, cha2, inp)
	return cha1, cha2
}

func pairAny(out1, out2 Anymode, inp Anymode) {
	defer out1.Close()
	defer out2.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out1.Provide(i)
		out2.Provide(i)
	}
}

// End of PairAny functions
// ===========================================================================
//
