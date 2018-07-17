// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// ===========================================================================
// Beg of PipeEnter/Leave - Flapdoors observed by a Waiter

// anyThingWaiter - as implemented by `*sync.WaitGroup` -
// attends Flapdoors and keeps counting
// who enters and who leaves.
//
// Use DoneWait to learn about
// when the facilities are closed.
//
// Note: You may also use Your provided `*sync.WaitGroup.Wait()`
// to know when to close the facilities.
// Just: DoneWait is more convenient
// as it also closes the primary channel for You.
//
// Just make sure to have _all_ entrances and exits attended,
// and `Wait()` only *after* You've started flooding the facilities.
type anyThingWaiter interface {
	Add(delta int)
	Done()
	Wait()
}

// Note: The name is intentionally generic in order to avoid eventual multiple-declaration clashes.

// PipeEnter returns a channel to receive
// all `inp`
// and registers throughput
// as arrival
// on the given `sync.WaitGroup`
// until close.
func (inp anyThingFrom) PipeEnter(wg anyThingWaiter) (out anyThingFrom) {
	cha := make(chan anyThing)
	go inp.pipeEnter(cha, wg)
	return cha
}

// PipeLeave returns a channel to receive
// all `inp`
// and registers throughput
// as departure
// on the given `sync.WaitGroup`
// until close.
func (inp anyThingFrom) PipeLeave(wg anyThingWaiter) (out anyThingFrom) {
	cha := make(chan anyThing)
	go inp.pipeLeave(cha, wg)
	return cha
}

// DoneLeave returns a channel to receive
// one signal after
// all throughput on `inp`
// has been registered
// as departure
// on the given `sync.WaitGroup`
// before close.
func (inp anyThingFrom) DoneLeave(wg anyThingWaiter) (done <-chan struct{}) {
	sig := make(chan struct{})
	go inp.doneLeave(sig, wg)
	return sig
}

func (inp anyThingFrom) pipeEnter(out anyThingInto, wg anyThingWaiter) {
	defer close(out)
	for i := range inp {
		wg.Add(1)
		out <- i
	}
}

func (inp anyThingFrom) pipeLeave(out anyThingInto, wg anyThingWaiter) {
	defer close(out)
	for i := range inp {
		out <- i
		wg.Done()
	}
}

func (inp anyThingFrom) doneLeave(done chan<- struct{}, wg anyThingWaiter) {
	defer close(done)
	for i := range inp {
		_ = i // discard
		wg.Done()
	}
	done <- struct{}{}
}

// TubeEnter returns a closure around PipeEnter (wg)
// registering throughput
// as arrival
// on the given `sync.WaitGroup`.
func (inp anyThingFrom) TubeEnter(wg anyThingWaiter) (tube func(inp anyThingFrom) (out anyThingFrom)) {

	return func(inp anyThingFrom) (out anyThingFrom) {
		return inp.PipeEnter(wg)
	}
}

// TubeLeave returns a closure around PipeLeave (wg)
// registering throughput
// as departure
// on the given `sync.WaitGroup`.
func (inp anyThingFrom) TubeLeave(wg anyThingWaiter) (tube func(inp anyThingFrom) (out anyThingFrom)) {

	return func(inp anyThingFrom) (out anyThingFrom) {
		return inp.PipeLeave(wg)
	}
}

// FiniLeave returns a closure around `DoneLeave(wg)`
// registering throughput
// as departure
// on the given `sync.WaitGroup`.
func (inp anyThingFrom) FiniLeave(wg anyThingWaiter) func(inp anyThingFrom) (done <-chan struct{}) {

	return func(inp anyThingFrom) (done <-chan struct{}) {
		return inp.DoneLeave(wg)
	}
}

// DoneWait returns a channel to receive
// one signal
// after wg.Wait() has returned and out has been closed
// before close.
//
// Note: Use only *after* You've started flooding the facilities.
func (out anyThingInto) DoneWait(wg anyThingWaiter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go out.doneWait(cha, wg)
	return cha
}

func (out anyThingInto) doneWait(done chan<- struct{}, wg anyThingWaiter) {
	defer close(done)
	wg.Wait()
	close(out)
	done <- struct{}{} // not really needed - but looks better
}

// FiniWait returns a closure around `DoneWait(wg)`.
func (out anyThingInto) FiniWait(wg anyThingWaiter) func(out anyThingInto) (done <-chan struct{}) {

	return func(out anyThingInto) (done <-chan struct{}) {
		return out.DoneWait(wg)
	}
}

// End of PipeEnter/Leave - Flapdoors observed by a Waiter
// ===========================================================================
